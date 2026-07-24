package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type Bumil struct {
	PendudukID int64  `json:"penduduk_id"`
	PosyanduID int64  `json:"posyandu_id"`
	Nama       string `json:"nama,omitempty"` // hasil JOIN dari tabel penduduk

	HPHT time.Time `json:"hpht"`
	HPL  time.Time `json:"hpl"`

	Gravida int `json:"gravida"`
	Para    int `json:"para"`
	Abortus int `json:"abortus"`
}

type BumilResponse struct {
	Penduduk struct {
		ID            int64  `json:"id"`
		NIK           int64  `json:"nik"`
		Nama          string `json:"nama"`
		JenisKelamin  string `json:"jenis_kelamin"`
		TanggalLahir  string `json:"tanggal_lahir"`
		KelurahanID   string `json:"kelurahan_id"`
		KelurahanNama string `json:"kelurahan_nama"`
		Alamat        string `json:"alamat"`
		NoHP          string `json:"no_hp"`
	} `json:"penduduk"`
	Bumil struct {
		HPHT         time.Time `json:"hpht"`
		HPL          time.Time `json:"hpl"`
		Gravida      int       `json:"gravida"`
		Para         int       `json:"para"`
		Abortus      int       `json:"abortus"`
		PosyanduID   int64     `json:"posyandu_id"`
		PosyanduNama string    `json:"posyandu_nama"`
	} `json:"bumil"`
}

func ValidateBumil(v *validator.Validator, b *Bumil) {
	v.Check(b.PendudukID > 0, "penduduk_id", "harus diisi")
	v.Check(b.PosyanduID > 0, "posyandu_id", "harus diisi")

	v.Check(!b.HPHT.IsZero(), "hpht", "harus diisi")
	v.Check(!b.HPL.IsZero(), "hpl", "harus diisi")

	v.Check(b.HPL.After(b.HPHT), "hpl", "harus setelah HPHT")

	v.Check(b.Gravida >= 1, "gravida", "minimal 1")
	v.Check(b.Gravida <= 20, "gravida", "tidak valid")

	v.Check(b.Para >= 0, "para", "tidak valid")
	v.Check(b.Para <= b.Gravida, "para", "tidak boleh lebih besar dari gravida")

	v.Check(b.Abortus >= 0, "abortus", "tidak valid")
	v.Check(b.Abortus <= b.Gravida, "abortus", "tidak boleh lebih besar dari gravida")
	v.Check(
		b.Para+b.Abortus <= b.Gravida,
		"gravida",
		"gravida harus lebih besar atau sama dengan para + abortus",
	)
}

type BumilModel struct {
	DB *sql.DB
}

func (m BumilModel) InsertTx(ctx context.Context, tx *sql.Tx, bumil *Bumil) error {
	query := `
INSERT INTO bumil (
    penduduk_id,
    posyandu_id,
    hpht,
    hpl,
    gravida,
		para,
		abortus
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

	_, err := tx.ExecContext(ctx, query,
		bumil.PendudukID,
		bumil.PosyanduID,
		bumil.HPHT,
		bumil.HPL,
		bumil.Gravida,
		bumil.Para,
		bumil.Abortus,
	)

	if err != nil {
		return err
	}
	return nil
}

func (m BumilModel) GetAll(posyandu_id int64, nama string, filters Filters) ([]*BumilResponse, Metadata, error) {

	query := fmt.Sprintf(`
		SELECT count(*) OVER(),
			p.id,
			p.nik,
			p.nama,
			p.jenis_kelamin,
			p.tanggal_lahir,
			p.kelurahan_id,
			k.name,
			p.alamat,
			p.no_hp,
		
			b.hpht,
			b.hpl,
			b.gravida,
			b.para,
			b.abortus,
			b.posyandu_id,
			pos.nama
	FROM bumil b
	JOIN penduduk p ON p.id = b.penduduk_id
	JOIN posyandu pos ON pos.id = b.posyandu_id
	JOIN kelurahan k ON k.id = p.kelurahan_id
	WHERE (LOWER(p.nama) LIKE LOWER('%%' || $1 || '%%') OR $1 = '')
  AND b.posyandu_id = $2
  AND p.deleted_at IS NULL
	ORDER BY %s %s, p.id ASC
	LIMIT $3 OFFSET $4`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{nama, posyandu_id, filters.limit(), filters.offset()}
	// args := []any{nama, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	bumil_all := []*BumilResponse{}

	for rows.Next() {

		var bumil BumilResponse

		err := rows.Scan(
			&totalRecords,
			&bumil.Penduduk.ID,
			&bumil.Penduduk.NIK,
			&bumil.Penduduk.Nama,
			&bumil.Penduduk.JenisKelamin,
			&bumil.Penduduk.TanggalLahir,
			&bumil.Penduduk.KelurahanID,
			&bumil.Penduduk.KelurahanNama,
			&bumil.Penduduk.Alamat,
			&bumil.Penduduk.NoHP,
			&bumil.Bumil.HPHT,
			&bumil.Bumil.HPL,
			&bumil.Bumil.Gravida,
			&bumil.Bumil.Para,
			&bumil.Bumil.Abortus,
			&bumil.Bumil.PosyanduID,
			&bumil.Bumil.PosyanduNama,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		bumil_all = append(bumil_all, &bumil)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return bumil_all, metadata, nil
}

func (m BumilModel) ValidateBumilInclude(posyanduID, pendudukID int64) error {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM bumil
			WHERE posyandu_id = $1
			  AND penduduk_id = $2
		)
	`

	var exists bool

	err := m.DB.QueryRow(query, posyanduID, pendudukID).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("bumil tidak terdaftar di posyandu ini")
	}

	return nil
}
