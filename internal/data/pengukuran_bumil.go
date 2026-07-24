package data

import (
	"context"
	"database/sql"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type PengukuranBumil struct {
	ID                    int64     `json:"id"`
	PendudukID            int64     `json:"penduduk_id"`
	Tanggal               time.Time `json:"tanggal"`
	BeratBadan            float64   `json:"berat_badan"`
	TinggiBadan           float64   `json:"tinggi_badan"`
	Hemoglobin            float64   `json:"hemoglobin"`
	Lila                  float64   `json:"lila"`
	UsiaKehamilanMinggu   int       `json:"usia_kehamilan_minggu"`
	TekananDarahSistolik  int       `json:"tekanan_darah_sistolik"`
	TekananDarahDiastolik int       `json:"tekanan_darah_diastolik"`
	Catatan               string    `json:"catatan"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	Version               int32     `json:"version"`
}

type PengukuranBumilResponse struct {
	ID                    int64     `json:"id"`
	PendudukID            int64     `json:"penduduk_id"`
	Nama                  string    `json:"nama"`
	Tanggal               time.Time `json:"tanggal"`
	BeratBadan            float64   `json:"berat_badan"`
	TinggiBadan           float64   `json:"tinggi_badan"`
	Hemoglobin            float64   `json:"hemoglobin"`
	Lila                  float64   `json:"lila"`
	UsiaKehamilanMinggu   int       `json:"usia_kehamilan_minggu"`
	TekananDarahSistolik  int       `json:"tekanan_darah_sistolik"`
	TekananDarahDiastolik int       `json:"tekanan_darah_diastolik"`
	Catatan               string    `json:"catatan"`
	CreatedAt             time.Time `json:"created_at"`
}

func ValidatePengukuranBumil(v *validator.Validator, p *PengukuranBumil) {
	v.Check(p.PendudukID > 0, "penduduk_id", "harus diisi")

	v.Check(!p.Tanggal.IsZero(), "tanggal", "harus diisi")

	if p.UsiaKehamilanMinggu > 0 {
		v.Check(
			p.UsiaKehamilanMinggu >= 1 && p.UsiaKehamilanMinggu <= 45,
			"usia_kehamilan_minggu",
			"tidak valid",
		)
	}

	if p.BeratBadan > 0 {
		v.Check(
			p.BeratBadan >= 20 && p.BeratBadan <= 300,
			"berat_badan",
			"tidak valid",
		)
	}

	if p.TinggiBadan > 0 {
		v.Check(
			p.TinggiBadan >= 100 && p.TinggiBadan <= 250,
			"tinggi_badan",
			"tidak valid",
		)
	}

	if p.Lila > 0 {
		v.Check(
			p.Lila >= 10 && p.Lila <= 60,
			"lila",
			"tidak valid",
		)
	}

	if p.Hemoglobin > 0 {
		v.Check(
			p.Hemoglobin >= 3 && p.Hemoglobin <= 25,
			"hemoglobin",
			"tidak valid",
		)
	}

	if p.TekananDarahSistolik > 0 {
		v.Check(
			p.TekananDarahSistolik >= 50 && p.TekananDarahSistolik <= 250,
			"tekanan_darah_sistolik",
			"tidak valid",
		)
	}

	if p.TekananDarahDiastolik > 0 {
		v.Check(
			p.TekananDarahDiastolik >= 30 && p.TekananDarahDiastolik <= 150,
			"tekanan_darah_diastolik",
			"tidak valid",
		)
	}

	v.Check(
		len(p.Catatan) <= 1000,
		"catatan",
		"maksimal 1000 karakter",
	)
}

type PengukuranBumilModel struct {
	DB *sql.DB
}

func (m PengukuranBumilModel) Insert(pengukuran_bumil *PengukuranBumil) error {
	query := `
	INSERT INTO pengukuran_bumil (
			penduduk_id,
			tanggal,
			berat_badan,
			tinggi_badan,
			hemoglobin,
			usia_kehamilan_minggu,
			tekanan_darah_sistolik,
			tekanan_darah_diastolik,
			lila,
			catatan
	)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id, created_at, version`

	args := []any{
		pengukuran_bumil.PendudukID,
		pengukuran_bumil.Tanggal,
		pengukuran_bumil.BeratBadan,
		pengukuran_bumil.TinggiBadan,
		pengukuran_bumil.Hemoglobin,
		pengukuran_bumil.UsiaKehamilanMinggu,
		pengukuran_bumil.TekananDarahSistolik,
		pengukuran_bumil.TekananDarahDiastolik,
		pengukuran_bumil.Lila,
		pengukuran_bumil.Catatan,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&pengukuran_bumil.ID,
		&pengukuran_bumil.CreatedAt,
		&pengukuran_bumil.Version,
	)
}

func (m PengukuranBumilModel) GetAll(posyandu_id int64, bulan string, nama string, filters Filters) ([]*PengukuranBumilResponse, Metadata, error) {

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
			pb.usia_kehamilan_minggu,
			pb.tekanan_darah_sistolik,
			pb.tekanan_darah_diastolik,
			pb.catatan
		FROM pengukuran_bumil pb
		JOIN penduduk p
			ON p.id = pb.penduduk_id
		JOIN bumil b
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
	pengukuran_bumil_all := []*PengukuranBumilResponse{}

	for rows.Next() {

		var pengukuran_bumil PengukuranBumilResponse

		err := rows.Scan(
			&totalRecords,
			&pengukuran_bumil.ID,
			&pengukuran_bumil.PendudukID,
			&pengukuran_bumil.Nama,
			&pengukuran_bumil.Tanggal,
			&pengukuran_bumil.BeratBadan,
			&pengukuran_bumil.TinggiBadan,
			&pengukuran_bumil.Hemoglobin,
			&pengukuran_bumil.Lila,
			&pengukuran_bumil.UsiaKehamilanMinggu,
			&pengukuran_bumil.TekananDarahSistolik,
			&pengukuran_bumil.TekananDarahDiastolik,
			&pengukuran_bumil.Catatan,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pengukuran_bumil_all = append(pengukuran_bumil_all, &pengukuran_bumil)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pengukuran_bumil_all, metadata, nil
}
