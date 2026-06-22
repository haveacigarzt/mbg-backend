package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type PengeluaranHarian struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	AlokasiHarianID int64     `json:"alokasi_harian_id"`

	Produk      string `json:"produk"`
	Jumlah      int64  `json:"jumlah"`
	Satuan      string `json:"satuan"`
	HargaSatuan int64  `json:"harga_satuan"`
	Subtotal    int64  `json:"subtotal"`
}

func ValidatePengeluaranHarian(v *validator.Validator, pengeluaran *PengeluaranHarian) {
	v.Check(pengeluaran.AlokasiHarianID > 0, "alokasi_harian_id", "must be provided")

	v.Check(pengeluaran.Produk != "", "produk", "must be provided")
	v.Check(len(pengeluaran.Produk) <= 255, "produk", "must not be more than 255 bytes long")

	v.Check(pengeluaran.Jumlah > 0, "jumlah", "must be greater than 0")

	v.Check(pengeluaran.Satuan != "", "satuan", "must be provided")
	v.Check(len(pengeluaran.Satuan) <= 50, "satuan", "must not be more than 50 bytes long")

	v.Check(pengeluaran.HargaSatuan > 0, "harga_satuan", "must be greater than 0")
}

type PengeluaranHarianModel struct {
	DB *sql.DB
}

func (m PengeluaranHarianModel) Insert(pengeluaran *PengeluaranHarian) (string, error) {
	query := `
		INSERT INTO pengeluaran_harian (
			alokasi_harian_id,
			produk,
			jumlah,
			satuan,
			harga_satuan
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING
			id,
			created_at,
			updated_at,
			subtotal,
			(
				SELECT TO_CHAR(ah.tanggal, 'YYYY-MM-DD')
				FROM alokasi_harian ah
				WHERE ah.id = alokasi_harian_id
			)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tanggal string

	err := m.DB.QueryRowContext(
		ctx,
		query,
		pengeluaran.AlokasiHarianID,
		pengeluaran.Produk,
		pengeluaran.Jumlah,
		pengeluaran.Satuan,
		pengeluaran.HargaSatuan,
	).Scan(
		&pengeluaran.ID,
		&pengeluaran.CreatedAt,
		&pengeluaran.UpdatedAt,
		&pengeluaran.Subtotal,
		&tanggal,
	)

	if err != nil {
		return "", err
	}

	return tanggal, nil
}

func (m PengeluaranHarianModel) Delete(id int64) (string, error) {
	if id < 1 {
		return "", ErrRecordNotFound
	}

	query := `
		DELETE FROM pengeluaran_harian ph
		USING alokasi_harian ah
		WHERE ph.alokasi_harian_id = ah.id
		  AND ph.id = $1
		RETURNING TO_CHAR(ah.tanggal, 'YYYY-MM-DD')
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tanggal string

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&tanggal)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return "", ErrRecordNotFound
		default:
			return "", err
		}
	}

	return tanggal, nil
}

func (m PengeluaranHarianModel) GetAll(alokasiHarianID int64, filters Filters) ([]*PengeluaranHarian, Metadata, error) {

	query := fmt.Sprintf(`
	SELECT count(*) OVER(),
		id,
		alokasi_harian_id,
		produk,
		jumlah,
		satuan,
		harga_satuan,
		subtotal,
		created_at
	FROM pengeluaran_harian
	WHERE alokasi_harian_id = $1
	ORDER BY %s %s, id ASC
	LIMIT $2 OFFSET $3`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{alokasiHarianID, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	pengeluaranHarian := []*PengeluaranHarian{}

	for rows.Next() {

		var pengeluaran PengeluaranHarian

		err := rows.Scan(
			&totalRecords,
			&pengeluaran.ID,
			&pengeluaran.AlokasiHarianID,
			&pengeluaran.Produk,
			&pengeluaran.Jumlah,
			&pengeluaran.Satuan,
			&pengeluaran.HargaSatuan,
			&pengeluaran.Subtotal,
			&pengeluaran.CreatedAt,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pengeluaranHarian = append(pengeluaranHarian, &pengeluaran)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pengeluaranHarian, metadata, nil
}

func (m PengeluaranHarianModel) GetAllByTanggal(tanggal string, filters Filters) ([]*PengeluaranHarian, Metadata, error) {

	query := `
	SELECT
		count(*) OVER(),
		ph.id,
		ph.alokasi_harian_id,
		ph.produk,
		ph.jumlah,
		ph.satuan,
		ph.harga_satuan,
		ph.subtotal,
		ph.created_at
	FROM pengeluaran_harian ph
	INNER JOIN alokasi_harian ah
		ON ah.id = ph.alokasi_harian_id
	WHERE ah.tanggal = $1::date
	`

	if filters.PageSize > 0 {
		query += fmt.Sprintf(`
		ORDER BY %s %s
		LIMIT $2 OFFSET $3
	`, filters.sortColumn(), filters.sortDirection())
	}

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
	pengeluaranHarian := []*PengeluaranHarian{}

	for rows.Next() {

		var pengeluaran PengeluaranHarian

		err := rows.Scan(
			&totalRecords,
			&pengeluaran.ID,
			&pengeluaran.AlokasiHarianID,
			&pengeluaran.Produk,
			&pengeluaran.Jumlah,
			&pengeluaran.Satuan,
			&pengeluaran.HargaSatuan,
			&pengeluaran.Subtotal,
			&pengeluaran.CreatedAt,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pengeluaranHarian = append(pengeluaranHarian, &pengeluaran)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	if filters.PageSize == 0 {
		return pengeluaranHarian, Metadata{
			TotalRecords: totalRecords,
		}, nil
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pengeluaranHarian, metadata, nil
}
