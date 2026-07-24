package data

import (
	"context"
	"database/sql"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type PengukuranBalita struct {
	ID            int64     `json:"id"`
	PendudukID    int64     `json:"penduduk_id"`
	Tanggal       time.Time `json:"tanggal"`
	UmurBulan     int16     `json:"umur_bulan"`
	BeratBadan    float64   `json:"berat_badan"`
	TinggiBadan   float64   `json:"tinggi_badan"`
	LingkarKepala float64   `json:"lingkar_kepala"`
	Lila          float64   `json:"lila"`
	Catatan       string    `json:"catatan"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Version       int32     `json:"version"`
}

type PengukuranBalitaResponse struct {
	ID            int64     `json:"id"`
	PendudukID    int64     `json:"penduduk_id"`
	Nama          string    `json:"nama"`
	IbuID         int64     `json:"ibu_id"`
	IbuNama       string    `json:"ibu_nama"`
	Tanggal       time.Time `json:"tanggal"`
	UmurBulan     int16     `json:"umur_bulan"`
	BeratBadan    float64   `json:"berat_badan"`
	TinggiBadan   float64   `json:"tinggi_badan"`
	LingkarKepala float64   `json:"lingkar_kepala"`
	Lila          float64   `json:"lila"`
	Catatan       string    `json:"catatan"`
	CreatedAt     time.Time `json:"created_at"`
}

func ValidatePengukuranBalita(v *validator.Validator, p *PengukuranBalita) {
	v.Check(p.PendudukID > 0, "penduduk_id", "harus diisi")

	v.Check(!p.Tanggal.IsZero(), "tanggal", "harus diisi")

	v.Check(p.UmurBulan >= 0, "umur_bulan", "tidak valid")
	v.Check(p.UmurBulan <= 60, "umur_bulan", "umur balita maksimal 60 bulan")

	v.Check(p.BeratBadan > 0, "berat_badan", "harus lebih dari 0")
	v.Check(p.BeratBadan <= 50, "berat_badan", "tidak valid")

	v.Check(p.TinggiBadan > 0, "tinggi_badan", "harus lebih dari 0")
	v.Check(p.TinggiBadan <= 150, "tinggi_badan", "tidak valid")

	if p.LingkarKepala > 0 {
		v.Check(p.LingkarKepala >= 20, "lingkar_kepala", "tidak valid")
		v.Check(p.LingkarKepala <= 70, "lingkar_kepala", "tidak valid")
	}

	if p.Lila > 0 {
		v.Check(p.Lila >= 5, "lila", "tidak valid")
		v.Check(p.Lila <= 40, "lila", "tidak valid")
	}

	v.Check(len(p.Catatan) <= 1000, "catatan", "maksimal 1000 karakter")
}

type PengukuranBalitaModel struct {
	DB *sql.DB
}

func (m PengukuranBalitaModel) Insert(pengukuran_balita *PengukuranBalita) error {
	query := `
	INSERT INTO pengukuran_balita (
			penduduk_id,
			tanggal,
			umur_bulan,
			berat_badan,
			tinggi_badan,
			lingkar_kepala,
			lila,
			catatan
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING id, created_at, version`

	args := []any{
		pengukuran_balita.PendudukID,
		pengukuran_balita.Tanggal,
		pengukuran_balita.UmurBulan,
		pengukuran_balita.BeratBadan,
		pengukuran_balita.TinggiBadan,
		pengukuran_balita.LingkarKepala,
		pengukuran_balita.Lila,
		pengukuran_balita.Catatan,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&pengukuran_balita.ID,
		&pengukuran_balita.CreatedAt,
		&pengukuran_balita.Version,
	)
}

func (m PengukuranBalitaModel) GetAll(posyandu_id int64, bulan string, nama string, filters Filters) ([]*PengukuranBalitaResponse, Metadata, error) {

	where := `
		WHERE
			(LOWER(p.nama) LIKE LOWER('%' || $1 || '%') OR $1 = '')
			AND b.posyandu_id = $2
			AND p.deleted_at IS NULL
`

	args := []any{nama, posyandu_id}
	argPos := 3

	if bulan != "" {
		where += fmt.Sprintf(`
        AND pb.tanggal >= $%d::date
        AND pb.tanggal < ($%d::date + INTERVAL '1 month')
    `, argPos, argPos)

		args = append(args, bulan+"-01")
		argPos++
	}

	args = append(args, filters.limit(), filters.offset())

	query := fmt.Sprintf(`
		SELECT count(*) OVER(),
			pb.id,
			pb.penduduk_id,
			p.nama,
			b.ibu_id,
			pi.nama AS ibu_nama,
			pb.tanggal,
			pb.umur_bulan,
			pb.berat_badan,
			pb.tinggi_badan,
			pb.lingkar_kepala,
			pb.lila,
			pb.catatan
		FROM pengukuran_balita pb
		JOIN penduduk p
			ON p.id = pb.penduduk_id
		JOIN balita b
			ON b.penduduk_id = p.id
		JOIN penduduk pi
			ON pi.id = b.ibu_id
		%s
		ORDER BY %s %s, p.id ASC
		LIMIT $%d OFFSET $%d`,
		where,
		filters.sortColumn(),
		filters.sortDirection(),
		argPos,
		argPos+1,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	pengukuran_balita_all := []*PengukuranBalitaResponse{}

	for rows.Next() {

		var pengukuran_balita PengukuranBalitaResponse

		err := rows.Scan(
			&totalRecords,
			&pengukuran_balita.ID,
			&pengukuran_balita.PendudukID,
			&pengukuran_balita.Nama,
			&pengukuran_balita.IbuID,
			&pengukuran_balita.IbuNama,
			&pengukuran_balita.Tanggal,
			&pengukuran_balita.UmurBulan,
			&pengukuran_balita.BeratBadan,
			&pengukuran_balita.TinggiBadan,
			&pengukuran_balita.LingkarKepala,
			&pengukuran_balita.Lila,
			&pengukuran_balita.Catatan,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pengukuran_balita_all = append(pengukuran_balita_all, &pengukuran_balita)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pengukuran_balita_all, metadata, nil
}
