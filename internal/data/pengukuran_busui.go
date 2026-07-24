package data

import (
	"context"
	"database/sql"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type PengukuranBusui struct {
	ID          int64     `json:"id"`
	PendudukID  int64     `json:"penduduk_id"`
	Tanggal     time.Time `json:"tanggal"`
	BeratBadan  float64   `json:"berat_badan"`
	TinggiBadan float64   `json:"tinggi_badan"`
	Hemoglobin  float64   `json:"hemoglobin"`
	Lila        float64   `json:"lila"`
	Catatan     string    `json:"catatan"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     int32     `json:"version"`
}

type PengukuranBusuiResponse struct {
	ID          int64     `json:"id"`
	PendudukID  int64     `json:"penduduk_id"`
	Nama        string    `json:"nama"`
	Tanggal     time.Time `json:"tanggal"`
	BeratBadan  float64   `json:"berat_badan"`
	TinggiBadan float64   `json:"tinggi_badan"`
	Hemoglobin  float64   `json:"hemoglobin"`
	Lila        float64   `json:"lila"`
	Catatan     string    `json:"catatan"`
	CreatedAt   time.Time `json:"created_at"`
}

func ValidatePengukuranBusui(v *validator.Validator, p *PengukuranBusui) {
	v.Check(p.PendudukID > 0, "penduduk_id", "harus diisi")

	v.Check(!p.Tanggal.IsZero(), "tanggal", "harus diisi")

	if p.BeratBadan > 0 {
		v.Check(p.BeratBadan >= 20, "berat_badan", "tidak valid")
		v.Check(p.BeratBadan <= 300, "berat_badan", "tidak valid")
	}

	if p.TinggiBadan > 0 {
		v.Check(p.TinggiBadan >= 100, "tinggi_badan", "tidak valid")
		v.Check(p.TinggiBadan <= 250, "tinggi_badan", "tidak valid")
	}

	if p.Hemoglobin > 0 {
		v.Check(p.Hemoglobin >= 3, "hemoglobin", "tidak valid")
		v.Check(p.Hemoglobin <= 25, "hemoglobin", "tidak valid")
	}

	if p.Lila > 0 {
		v.Check(p.Lila >= 10, "lila", "tidak valid")
		v.Check(p.Lila <= 60, "lila", "tidak valid")
	}

	v.Check(len(p.Catatan) <= 1000, "catatan", "maksimal 1000 karakter")
}

type PengukuranBusuiModel struct {
	DB *sql.DB
}

func (m PengukuranBusuiModel) Insert(pengukuran_busui *PengukuranBusui) error {
	query := `
	INSERT INTO pengukuran_busui (
			penduduk_id,
			tanggal,
			berat_badan,
			tinggi_badan,
			hemoglobin,
			lila,
			catatan
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, created_at, version`

	args := []any{
		pengukuran_busui.PendudukID,
		pengukuran_busui.Tanggal,
		pengukuran_busui.BeratBadan,
		pengukuran_busui.TinggiBadan,
		pengukuran_busui.Hemoglobin,
		pengukuran_busui.Lila,
		pengukuran_busui.Catatan,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&pengukuran_busui.ID,
		&pengukuran_busui.CreatedAt,
		&pengukuran_busui.Version,
	)
}

func (m PengukuranBusuiModel) GetAll(posyandu_id int64,
	bulan string, // format: "2026-06", kosong jika tidak difilter
	nama string, filters Filters) ([]*PengukuranBusuiResponse, Metadata, error) {

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
        pb.tanggal,
        pb.berat_badan,
        pb.tinggi_badan,
        pb.hemoglobin,
        pb.lila,
        pb.catatan
    FROM pengukuran_busui pb
    JOIN penduduk p
        ON p.id = pb.penduduk_id
    JOIN busui b
        ON b.penduduk_id = p.id
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
	pengukuran_busui_all := []*PengukuranBusuiResponse{}

	for rows.Next() {

		var pengukuran_busui PengukuranBusuiResponse

		err := rows.Scan(
			&totalRecords,
			&pengukuran_busui.ID,
			&pengukuran_busui.PendudukID,
			&pengukuran_busui.Nama,
			&pengukuran_busui.Tanggal,
			&pengukuran_busui.BeratBadan,
			&pengukuran_busui.TinggiBadan,
			&pengukuran_busui.Hemoglobin,
			&pengukuran_busui.Lila,
			&pengukuran_busui.Catatan,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pengukuran_busui_all = append(pengukuran_busui_all, &pengukuran_busui)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pengukuran_busui_all, metadata, nil
}
