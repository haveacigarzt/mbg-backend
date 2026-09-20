package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type Balita struct {
	PendudukID   int64  `json:"penduduk_id"`
	PosyanduID   int64  `json:"posyandu_id"`
	Nama         string `json:"nama,omitempty"` // hasil JOIN dari tabel penduduk
	IbuID        int64  `json:"ibu_id"`
	AnakKe       int8   `json:"anak_ke"`
	BeratLahir   int    `json:"berat_lahir"`
	PanjangLahir int    `json:"panjang_lahir"`
}

type BAResponse struct {
	IbuID        int64  `json:"ibu_id"`
	IbuNama      string `json:"ibu_nama"`
	AnakKe       int8   `json:"anak_ke"`
	BeratLahir   int    `json:"berat_lahir"`
	PanjangLahir int    `json:"panjang_lahir"`
	PosyanduID   int64  `json:"posyandu_id"`
	PosyanduNama string `json:"posyandu_nama"`
	StatusAktif  bool   `json:"status_aktif"`
}

type BalitaResponse struct {
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
	Balita BAResponse `json:"balita"`
}

func ValidateBalita(v *validator.Validator, b *Balita) {
	v.Check(b.PendudukID > 0, "penduduk_id", "harus diisi")
	v.Check(b.PosyanduID > 0, "posyandu_id", "harus diisi")

	if b.IbuID != 0 {
		v.Check(b.IbuID > 0, "ibu_id", "tidak valid")
	}

	v.Check(b.AnakKe > 0, "anak_ke", "harus lebih dari 0")
	v.Check(b.AnakKe <= 20, "anak_ke", "tidak valid")

	v.Check(b.BeratLahir > 0, "berat_lahir", "harus diisi")
	v.Check(b.BeratLahir >= 500, "berat_lahir", "tidak valid")
	v.Check(b.BeratLahir <= 7000, "berat_lahir", "tidak valid")

	v.Check(b.PanjangLahir > 0, "panjang_lahir", "harus diisi")
	v.Check(b.PanjangLahir >= 20, "panjang_lahir", "tidak valid")
	v.Check(b.PanjangLahir <= 70, "panjang_lahir", "tidak valid")
}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type BalitaModel struct {
	DB *sql.DB
}

func (m BalitaModel) InsertTx(ctx context.Context, tx *sql.Tx, balita *Balita) error {
	query := `
INSERT INTO balita (
    penduduk_id,
    posyandu_id,
    ibu_id,
    anak_ke,
    berat_lahir,
    panjang_lahir
)
VALUES ($1, $2, $3, $4, $5, $6)
`

	_, err := tx.ExecContext(ctx, query,
		balita.PendudukID,
		balita.PosyanduID,
		balita.IbuID,
		balita.AnakKe,
		balita.BeratLahir,
		balita.PanjangLahir,
	)

	if err != nil {
		return err
	}
	return nil
}

func (m BalitaModel) GetAll(posyandu_id int64, nama string, filters Filters) ([]*BalitaResponse, Metadata, error) {

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
		
			b.ibu_id,
			i.nama,
			b.anak_ke,
			b.berat_lahir,
			b.panjang_lahir,
			b.posyandu_id,
			b.status_aktif,
			pos.nama
	FROM balita b
	JOIN penduduk p ON p.id = b.penduduk_id
	JOIN posyandu pos ON pos.id = b.posyandu_id
	JOIN kelurahan k ON k.id = p.kelurahan_id
	JOIN penduduk i ON i.id = b.ibu_id
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
	balita_all := []*BalitaResponse{}

	for rows.Next() {

		var balita BalitaResponse

		err := rows.Scan(
			&totalRecords,
			&balita.Penduduk.ID,
			&balita.Penduduk.NIK,
			&balita.Penduduk.Nama,
			&balita.Penduduk.JenisKelamin,
			&balita.Penduduk.TanggalLahir,
			&balita.Penduduk.KelurahanID,
			&balita.Penduduk.KelurahanNama,
			&balita.Penduduk.Alamat,
			&balita.Penduduk.NoHP,
			&balita.Balita.IbuID,
			&balita.Balita.IbuNama,
			&balita.Balita.AnakKe,
			&balita.Balita.BeratLahir,
			&balita.Balita.PanjangLahir,
			&balita.Balita.PosyanduID,
			&balita.Balita.StatusAktif,
			&balita.Balita.PosyanduNama,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		balita_all = append(balita_all, &balita)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return balita_all, metadata, nil
}

func (m BalitaModel) ValidateBalitaInclude(posyanduID, pendudukID int64) error {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM balita
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
		return errors.New("balita tidak terdaftar di posyandu ini")
	}

	return nil
}
