package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type Busui struct {
	PendudukID        int64     `json:"penduduk_id"`
	PosyanduID        int64     `json:"posyandu_id"`
	Nama              string    `json:"nama,omitempty"` // hasil JOIN dari tabel penduduk
	TanggalPersalinan time.Time `json:"tanggal_persalinan"`
	AnakKe            int8      `json:"anak_ke"`
	AsiEksklusif      bool      `json:"asi_eksklusif"`
}

type BusuiResponse struct {
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
	Busui struct {
		TanggalPersalinan time.Time `json:"tanggal_persalinan"`
		AnakKe            int8      `json:"anak_ke"`
		AsiEksklusif      bool      `json:"asi_eksklusif"`
		PosyanduID        int64     `json:"posyandu_id"`
		PosyanduNama      string    `json:"posyandu_nama"`
	} `json:"busui"`
}

func ValidateBusui(v *validator.Validator, b *Busui) {
	v.Check(b.PendudukID > 0, "penduduk_id", "harus diisi")
	v.Check(b.PosyanduID > 0, "posyandu_id", "harus diisi")

	v.Check(!b.TanggalPersalinan.IsZero(), "tanggal_persalinan", "harus diisi")

	v.Check(b.AnakKe > 0, "anak_ke", "harus lebih dari 0")
	v.Check(b.AnakKe <= 20, "anak_ke", "tidak valid")
}

type BusuiModel struct {
	DB *sql.DB
}

func (m BusuiModel) InsertTx(ctx context.Context, tx *sql.Tx, busui *Busui) error {
	query := `
INSERT INTO busui (
    penduduk_id,
    posyandu_id,
    tanggal_persalinan,
    anak_ke,
    asi_eksklusif
)
VALUES ($1, $2, $3, $4, $5)
`

	_, err := tx.ExecContext(ctx, query,
		busui.PendudukID,
		busui.PosyanduID,
		busui.TanggalPersalinan,
		busui.AnakKe,
		busui.AsiEksklusif,
	)

	if err != nil {
		return err
	}
	return nil
}

func (m BusuiModel) GetAll(posyandu_id int64, nama string, filters Filters) ([]*BusuiResponse, Metadata, error) {

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
		
			b.tanggal_persalinan,
			b.anak_ke,
			b.asi_eksklusif,
			b.posyandu_id,
			pos.nama
	FROM busui b
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
	busui_all := []*BusuiResponse{}

	for rows.Next() {

		var bumil BusuiResponse

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
			&bumil.Busui.TanggalPersalinan,
			&bumil.Busui.AnakKe,
			&bumil.Busui.AsiEksklusif,
			&bumil.Busui.PosyanduID,
			&bumil.Busui.PosyanduNama,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		busui_all = append(busui_all, &bumil)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return busui_all, metadata, nil
}

func (m BusuiModel) ValidateBusuiInclude(posyanduID, pendudukID int64) error {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM busui
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
		return errors.New("busui tidak terdaftar di posyandu ini")
	}

	return nil
}
