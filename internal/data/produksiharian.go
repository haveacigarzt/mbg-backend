package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type ProduksiHarian struct {
	ID                   int64     `json:"id"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	SPPGID               int64     `json:"sppg_id,omitempty"`
	Tanggal              string    `json:"tanggal"`
	WaktuMulai           time.Time `json:"waktu_mulai"`
	EstimasiWaktuSelesai time.Time `json:"estimasi_waktu_selesai"`
}

type ProduksiHarianOutput struct {
	ID                   *int64     `json:"id"`
	CreatedAt            *time.Time `json:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at"`
	SPPGID               int64      `json:"sppg_id"`
	SPPGNama             string     `json:"sppg_nama"`
	Tanggal              *string    `json:"tanggal"`
	WaktuMulai           *time.Time `json:"waktu_mulai"`
	EstimasiWaktuSelesai *time.Time `json:"estimasi_waktu_selesai"`
	RowID                int        `json:"row_id"`
}

func ValidateProduksiHarian(v *validator.Validator, produksi *ProduksiHarian) {
	v.Check(produksi.SPPGID > 0, "sppg_id", "must be provided")

	v.Check(produksi.Tanggal != "", "tanggal", "must be provided")

	_, err := time.Parse("2006-01-02", produksi.Tanggal)
	v.Check(err == nil, "tanggal", "must be a valid date (YYYY-MM-DD)")

	v.Check(!produksi.WaktuMulai.IsZero(),
		"waktu_mulai",
		"must be provided")

	v.Check(!produksi.EstimasiWaktuSelesai.IsZero(),
		"estimasi_waktu_selesai",
		"must be provided")

	v.Check(
		produksi.EstimasiWaktuSelesai.After(produksi.WaktuMulai),
		"estimasi_waktu_selesai",
		"must be after waktu_mulai",
	)
}

func ValidateGetProduksiHarian(v *validator.Validator, produksi *ProduksiHarian) {

	v.Check(produksi.Tanggal != "", "tanggal", "must be provided")

	_, err := time.Parse("2006-01-02", produksi.Tanggal)
	v.Check(err == nil, "tanggal", "must be a valid date (YYYY-MM-DD)")
}

type ProduksiHarianModel struct {
	DB *sql.DB
}

func (m ProduksiHarianModel) Get(tanggal string, sppgID int64) (*ProduksiHarian, error) {
	query := `
		SELECT
			id,
			created_at,
			updated_at,
			sppg_id,
			waktu_mulai,
			estimasi_waktu_selesai,
			TO_CHAR(tanggal, 'YYYY-MM-DD') AS tanggal
		FROM produksi_harian
		WHERE sppg_id = $1 AND tanggal = $2
	`

	var produksiHarian ProduksiHarian

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(
		ctx,
		query,
		sppgID,
		tanggal,
	).Scan(
		&produksiHarian.ID,
		&produksiHarian.CreatedAt,
		&produksiHarian.UpdatedAt,
		&produksiHarian.SPPGID,
		&produksiHarian.WaktuMulai,
		&produksiHarian.EstimasiWaktuSelesai,
		&produksiHarian.Tanggal,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &produksiHarian, nil
}

func (m ProduksiHarianModel) GetAll(tanggal string, filters Filters) ([]*ProduksiHarianOutput, Metadata, error) {
	query := `
	SELECT
		count(*) OVER(),
		s.id AS sppg_id,
		s.nama AS sppg_nama,
		ph.id,
		ph.created_at,
		ph.updated_at,
		ph.waktu_mulai,
		ph.estimasi_waktu_selesai,
		TO_CHAR(ph.tanggal, 'YYYY-MM-DD') AS tanggal
	FROM sppg s
	LEFT JOIN produksi_harian ph
		ON ph.sppg_id = s.id
		AND ph.tanggal = $1::date
`
	if filters.PageSize > 0 {
		query += fmt.Sprintf(`
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, filters.sortColumn(), filters.sortDirection())
	}
	//
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{tanggal}

	if filters.PageSize > 0 {
		args = append(args, filters.limit(), filters.offset())
	}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	produksiHarian := []*ProduksiHarianOutput{}
	id := 0

	for rows.Next() {

		var produksi ProduksiHarianOutput

		err := rows.Scan(
			&totalRecords,
			&produksi.SPPGID,
			&produksi.SPPGNama,
			&produksi.ID,
			&produksi.CreatedAt,
			&produksi.UpdatedAt,
			&produksi.WaktuMulai,
			&produksi.EstimasiWaktuSelesai,
			&produksi.Tanggal,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		if produksi.Tanggal == nil {
			produksi.Tanggal = &tanggal
		}
		produksi.RowID = id
		id++

		produksiHarian = append(produksiHarian, &produksi)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	if filters.PageSize == 0 {
		return produksiHarian, Metadata{
			TotalRecords: totalRecords,
		}, nil
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return produksiHarian, metadata, nil
}

func (m ProduksiHarianModel) Insert(produksi *ProduksiHarian) error {
	query := `
		INSERT INTO produksi_harian (
			sppg_id,
			waktu_mulai,
			estimasi_waktu_selesai,
			tanggal
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			id,
			created_at,
			updated_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(
		ctx,
		query,
		produksi.SPPGID,
		produksi.WaktuMulai,
		produksi.EstimasiWaktuSelesai,
		produksi.Tanggal,
	).Scan(
		&produksi.ID,
		&produksi.CreatedAt,
		&produksi.UpdatedAt,
	)
}

func (m ProduksiHarianModel) Update(produksi *ProduksiHarian) error {
	query := `
		UPDATE produksi_harian
		SET
			waktu_mulai = $1,
			estimasi_waktu_selesai = $2,
			updated_at = NOW()
		WHERE id = $3
		RETURNING updated_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(
		ctx,
		query,
		produksi.WaktuMulai,
		produksi.EstimasiWaktuSelesai,
		produksi.ID,
	).Scan(&produksi.UpdatedAt)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrRecordNotFound
		default:
			return err
		}
	}

	return nil
}
