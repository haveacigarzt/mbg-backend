package data

import (
	"context"
	"database/sql"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type PengukuranAnak struct {
	ID          int64     `json:"id"`
	PendudukID  int64     `json:"penduduk_id"`
	Tanggal     time.Time `json:"tanggal"`
	UmurBulan   int16     `json:"umur_bulan"`
	BeratBadan  float64   `json:"berat_badan"`
	TinggiBadan float64   `json:"tinggi_badan"`
	Catatan     string    `json:"catatan"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     int32     `json:"version"`
}

type PengukuranPesertaDidikResponse struct {
	ID          int64     `json:"id"`
	PendudukID  int64     `json:"penduduk_id"`
	Nama        string    `json:"nama"`
	Kelas       string    `json:"kelas"`
	Rombel      string    `json:"rombel"`
	Tanggal     time.Time `json:"tanggal"`
	UmurBulan   int16     `json:"umur_bulan"`
	BeratBadan  float64   `json:"berat_badan"`
	TinggiBadan float64   `json:"tinggi_badan"`
	Catatan     string    `json:"catatan"`
	CreatedAt   time.Time `json:"created_at"`
}

func ValidatePengukuranAnak(v *validator.Validator, p *PengukuranAnak) {
	v.Check(p.PendudukID > 0, "penduduk_id", "harus diisi")

	v.Check(!p.Tanggal.IsZero(), "tanggal", "harus diisi")

	v.Check(p.UmurBulan >= 0, "umur_bulan", "tidak valid")
	v.Check(p.UmurBulan <= 60, "umur_bulan", "tidak valid")

	v.Check(p.BeratBadan > 0, "berat_badan", "harus lebih dari 0")
	v.Check(p.BeratBadan <= 50, "berat_badan", "tidak valid")

	v.Check(p.TinggiBadan > 0, "tinggi_badan", "harus lebih dari 0")
	v.Check(p.TinggiBadan <= 150, "tinggi_badan", "tidak valid")

	v.Check(len(p.Catatan) <= 1000, "catatan", "maksimal 1000 karakter")
}

type PengukuranAnakModel struct {
	DB *sql.DB
}

func (m PengukuranAnakModel) Insert(pengukuran_anak *PengukuranAnak) error {
	query := `
	INSERT INTO pengukuran_anak (
			penduduk_id,
			tanggal,
			umur_bulan,
			berat_badan,
			tinggi_badan,
			catatan
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id, created_at, version`

	args := []any{
		pengukuran_anak.PendudukID,
		pengukuran_anak.Tanggal,
		pengukuran_anak.UmurBulan,
		pengukuran_anak.BeratBadan,
		pengukuran_anak.TinggiBadan,
		pengukuran_anak.Catatan,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&pengukuran_anak.ID,
		&pengukuran_anak.CreatedAt,
		&pengukuran_anak.Version,
	)
}

func (m PengukuranAnakModel) GetAllPSD(sekolah_id int64, bulan string, nama string, filters Filters) ([]*PengukuranPesertaDidikResponse, Metadata, error) {

	where := `
		WHERE
			(LOWER(p.nama) LIKE LOWER('%' || $1 || '%') OR $1 = '')
			AND pd.sekolah_id = $2
			AND p.deleted_at IS NULL
`

	args := []any{nama, sekolah_id}
	argPos := 3

	if bulan != "" {
		where += fmt.Sprintf(`
        AND pa.tanggal >= $%d::date
        AND pa.tanggal < ($%d::date + INTERVAL '1 month')
    `, argPos, argPos)

		args = append(args, bulan+"-01")
		argPos++
	}

	args = append(args, filters.limit(), filters.offset())

	query := fmt.Sprintf(`
		SELECT count(*) OVER(),
			pa.id,
			pa.penduduk_id,
			p.nama,
			pd.kelas,
			pd.rombel,
			pa.tanggal,
			pa.umur_bulan,
			pa.berat_badan,
			pa.tinggi_badan,
			pa.catatan
		FROM pengukuran_anak pa
		JOIN penduduk p
			ON p.id = pa.penduduk_id
		JOIN peserta_didik pd
			ON pd.penduduk_id = p.id
		JOIN sekolah s
			ON s.id = pd.sekolah_id
		JOIN kelurahan k
			ON k.id = p.kelurahan_id
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
	pengukuran_peserta_didik_all := []*PengukuranPesertaDidikResponse{}

	for rows.Next() {

		var pengukuran_peserta_didik PengukuranPesertaDidikResponse

		err := rows.Scan(
			&totalRecords,
			&pengukuran_peserta_didik.ID,
			&pengukuran_peserta_didik.PendudukID,
			&pengukuran_peserta_didik.Nama,
			&pengukuran_peserta_didik.Kelas,
			&pengukuran_peserta_didik.Rombel,
			&pengukuran_peserta_didik.Tanggal,
			&pengukuran_peserta_didik.UmurBulan,
			&pengukuran_peserta_didik.BeratBadan,
			&pengukuran_peserta_didik.TinggiBadan,
			&pengukuran_peserta_didik.Catatan,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pengukuran_peserta_didik_all = append(pengukuran_peserta_didik_all, &pengukuran_peserta_didik)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pengukuran_peserta_didik_all, metadata, nil
}
