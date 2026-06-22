package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type AlokasiHarian struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	SPPGID    int64     `json:"sppg_id,omitempty"`
	SPPGNama  string    `json:"sppg_nama,omitempty"`
	Jumlah    int64     `json:"jumlah"`
	Tanggal   string    `json:"tanggal"`
}

type AlokasiHarianOutput struct {
	SPPGID   int64  `json:"sppg_id"`
	SPPGNama string `json:"sppg_nama"`

	ID        *int64     `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	Jumlah    *int64     `json:"jumlah"`
	Tanggal   *string    `json:"tanggal"`
	RowID     int        `json:"row_id"`
}

func ValidateAlokasiHarian(v *validator.Validator, alokasi *AlokasiHarian) {
	v.Check(alokasi.SPPGID > 0, "sppg_id", "must be provided")

	v.Check(alokasi.Jumlah > 0, "jumlah", "must be greater than 0")

	v.Check(alokasi.Tanggal != "", "tanggal", "must be provided")

	_, err := time.Parse("2006-01-02", alokasi.Tanggal)
	v.Check(err == nil, "tanggal", "must be a valid date (YYYY-MM-DD)")
}

func ValidateGetAlokasiHarian(v *validator.Validator, alokasi *AlokasiHarian) {

	v.Check(alokasi.Tanggal != "", "tanggal", "must be provided")

	_, err := time.Parse("2006-01-02", alokasi.Tanggal)
	v.Check(err == nil, "tanggal", "must be a valid date (YYYY-MM-DD)")
}

type AlokasiHarianModel struct {
	DB *sql.DB
}

func (m AlokasiHarianModel) Get(tanggal string, sppgID int64) (*AlokasiHarian, error) {
	query := `
		SELECT
			ah.id,
		ah.created_at,
		ah.updated_at,
		ah.sppg_id,
		s.nama,
		ah.jumlah,
		TO_CHAR(ah.tanggal, 'YYYY-MM-DD') AS tanggal
	FROM alokasi_harian ah
	INNER JOIN sppg s
		ON s.id = ah.sppg_id
		WHERE ah.sppg_id = $1 AND ah.tanggal = $2
	`

	var alokasiHarian AlokasiHarian

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(
		ctx,
		query,
		sppgID,
		tanggal,
	).Scan(
		&alokasiHarian.ID,
		&alokasiHarian.CreatedAt,
		&alokasiHarian.UpdatedAt,
		&alokasiHarian.SPPGID,
		&alokasiHarian.SPPGNama,
		&alokasiHarian.Jumlah,
		&alokasiHarian.Tanggal,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &alokasiHarian, nil
}
func (m AlokasiHarianModel) GetByID(alokasiID int64) (*AlokasiHarian, error) {
	query := `
		SELECT
			ah.id,
		ah.created_at,
		ah.updated_at,
		ah.sppg_id,
		s.nama,
		ah.jumlah,
		TO_CHAR(ah.tanggal, 'YYYY-MM-DD') AS tanggal
	FROM alokasi_harian ah
	INNER JOIN sppg s
		ON s.id = ah.sppg_id
		WHERE ah.id = $1
	`

	var alokasiHarian AlokasiHarian

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(
		ctx,
		query,
		alokasiID,
	).Scan(
		&alokasiHarian.ID,
		&alokasiHarian.CreatedAt,
		&alokasiHarian.UpdatedAt,
		&alokasiHarian.SPPGID,
		&alokasiHarian.SPPGNama,
		&alokasiHarian.Jumlah,
		&alokasiHarian.Tanggal,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &alokasiHarian, nil
}

func (m AlokasiHarianModel) GetAll(tanggal string, filters Filters) ([]*AlokasiHarianOutput, Metadata, error) {
	query := `
	SELECT
		count(*) OVER(),
		s.id AS sppg_id,
		s.nama,
		ah.id,
		ah.created_at,
		ah.updated_at,
		ah.jumlah,
		TO_CHAR(ah.tanggal, 'YYYY-MM-DD') AS tanggal
	FROM sppg s
	LEFT JOIN alokasi_harian ah
		ON ah.sppg_id = s.id
		AND ah.tanggal = $1::date
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
	alokasiHarian := []*AlokasiHarianOutput{}
	id := 0

	for rows.Next() {

		var alokasi AlokasiHarianOutput

		err := rows.Scan(
			&totalRecords,
			&alokasi.SPPGID,
			&alokasi.SPPGNama,
			&alokasi.ID,
			&alokasi.CreatedAt,
			&alokasi.UpdatedAt,
			&alokasi.Jumlah,
			&alokasi.Tanggal,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		if alokasi.Tanggal == nil {
			alokasi.Tanggal = &tanggal
		}
		alokasi.RowID = id
		id++

		alokasiHarian = append(alokasiHarian, &alokasi)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	if filters.PageSize == 0 {
		return alokasiHarian, Metadata{
			TotalRecords: totalRecords,
		}, nil
	}
	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return alokasiHarian, metadata, nil
}

func (m AlokasiHarianModel) Insert(alokasi *AlokasiHarian) error {
	query := `
		INSERT INTO alokasi_harian (
			sppg_id,
			jumlah,
			tanggal
		)
		VALUES ($1, $2, $3)
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
		alokasi.SPPGID,
		alokasi.Jumlah,
		alokasi.Tanggal,
	).Scan(
		&alokasi.ID,
		&alokasi.CreatedAt,
		&alokasi.UpdatedAt,
	)
}

func (m AlokasiHarianModel) Update(alokasi *AlokasiHarian) error {
	query := `
		UPDATE alokasi_harian
		SET
			jumlah = $1,
			updated_at = NOW()
		WHERE id = $2
		RETURNING updated_at
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(
		ctx,
		query,
		alokasi.Jumlah,
		alokasi.ID,
	).Scan(&alokasi.UpdatedAt)

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
