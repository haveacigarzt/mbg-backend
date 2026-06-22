package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type Posyandu struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Nama           string    `json:"nama"`
	Alamat         string    `json:"alamat"`
	Kecamatan      string    `json:"kecamatan,omitempty"`
	Kelurahan      string    `json:"kelurahan,omitempty"`
	Kecamatan_ID   int64     `json:"kecamatan_id"`
	Kelurahan_ID   int64     `json:"kelurahan_id"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	JumlahBalita   int       `json:"jumlah_balita"`
	JumlahIbuHamil int       `json:"jumlah_ibu_hamil"`
	SPPGID         int64     `json:"sppg_id"`
	Version        int32     `json:"version"`
}

func ValidatePosyandu(v *validator.Validator, posyandu *Posyandu) {
	v.Check(posyandu.Nama != "", "nama", "must be provided")
	v.Check(len(posyandu.Nama) <= 200, "nama", "must not be more than 200 bytes long")

	v.Check(posyandu.Alamat != "", "alamat", "must be provided")
	v.Check(len(posyandu.Alamat) <= 500, "alamat", "must not be more than 500 bytes long")

	v.Check(posyandu.Kecamatan_ID != 0, "kecamatan_id", "must be provided")

	v.Check(posyandu.Kelurahan_ID != 0, "kelurahan_id", "must be provided")

	v.Check(posyandu.JumlahBalita >= 0,
		"jumlah_balita",
		"must be greater than or equal to 0",
	)

	v.Check(posyandu.JumlahIbuHamil >= 0,
		"jumlah_ibu_hamil",
		"must be greater than or equal to 0",
	)

	v.Check(
		posyandu.Latitude >= -90 && posyandu.Latitude <= 90,
		"latitude",
		"must be between -90 and 90",
	)

	v.Check(
		posyandu.Longitude >= -180 && posyandu.Longitude <= 180,
		"longitude",
		"must be between -180 and 180",
	)

	v.Check(
		posyandu.SPPGID > 0,
		"sppg_id",
		"must be provided",
	)
}

type PosyanduModel struct {
	DB *sql.DB
}

func (m PosyanduModel) Insert(posyandu *Posyandu) error {
	query := `
INSERT INTO posyandu (
    nama,
    alamat,
    kecamatan_id,
    kelurahan_id,
    latitude,
    longitude,
    jumlah_balita,
    jumlah_ibu_hamil,
    sppg_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, created_at, version`

	args := []any{
		posyandu.Nama,
		posyandu.Alamat,
		posyandu.Kecamatan_ID,
		posyandu.Kelurahan_ID,
		posyandu.Latitude,
		posyandu.Longitude,
		posyandu.JumlahBalita,
		posyandu.JumlahIbuHamil,
		posyandu.SPPGID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&posyandu.ID,
		&posyandu.CreatedAt,
		&posyandu.Version,
	)
}

func (m PosyanduModel) GetAll(nama string, kecamatan_id int64, kelurahan_id int64, sppg_id int64, jumlahBalita int, jumlahIbuHamil int, filters Filters) ([]*Posyandu, Metadata, error) {

	query := fmt.Sprintf(`
	SELECT count(*) OVER(),
		s.id,
		s.created_at,
		s.nama,
		s.alamat,
		s.jumlah_balita,
		s.jumlah_ibu_hamil,
		s.kecamatan_id,
		k.name AS kecamatan,
		s.kelurahan_id,
		kel.name AS kelurahan,
		s.latitude,
		s.longitude,
		s.sppg_id,
		s.version
	FROM posyandu s
	LEFT JOIN kecamatan k ON k.id = s.kecamatan_id
	LEFT JOIN kelurahan kel ON kel.id = s.kelurahan_id
	WHERE (LOWER(s.nama) LIKE LOWER('%%' || $1 || '%%') OR $1 = '')
	AND (s.kecamatan_id = $2 OR $2 = 0)
	AND (s.kelurahan_id = $3 OR $3 = 0)
	AND (s.sppg_id = $4 OR $4 = 0)
	AND (s.jumlah_balita >= $5 OR $5 = 0)
	AND (s.jumlah_ibu_hamil >= $6 OR $6 = 0)
	ORDER BY %s %s, id ASC
	LIMIT $7 OFFSET $8`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{nama, kecamatan_id, kelurahan_id, sppg_id, jumlahBalita, jumlahIbuHamil, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	posyandu_all := []*Posyandu{}

	for rows.Next() {

		var posyandu Posyandu

		err := rows.Scan(
			&totalRecords,
			&posyandu.ID,
			&posyandu.CreatedAt,
			&posyandu.Nama,
			&posyandu.Alamat,
			&posyandu.JumlahBalita,
			&posyandu.JumlahIbuHamil,
			&posyandu.Kecamatan_ID,
			&posyandu.Kecamatan,
			&posyandu.Kelurahan_ID,
			&posyandu.Kelurahan,
			&posyandu.Latitude,
			&posyandu.Longitude,
			&posyandu.SPPGID,
			&posyandu.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		posyandu_all = append(posyandu_all, &posyandu)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return posyandu_all, metadata, nil
}

func (m PosyanduModel) Get(id int64) (*Posyandu, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT
			s.id,
			s.created_at,
			s.nama,
			s.alamat,
			s.jumlah_balita,
			s.jumlah_ibu_hamil,
			s.kecamatan_id,
			k.name AS kecamatan,
			s.kelurahan_id,
			kel.name AS kelurahan,
			s.latitude,
			s.longitude,
			s.sppg_id,
			s.version
		FROM posyandu s
		LEFT JOIN kecamatan k ON k.id = s.kecamatan_id
		LEFT JOIN kelurahan kel ON kel.id = s.kelurahan_id
		WHERE s.id = $1
	`

	var posyandu Posyandu

	// Use the context.WithTimeout() function to create a context.Context which carries a
	// 3-second timeout deadline. Note that we're using the empty context.Background()
	// as the 'parent' context.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&posyandu.ID,
		&posyandu.CreatedAt,
		&posyandu.Nama,
		&posyandu.Alamat,
		&posyandu.JumlahBalita,
		&posyandu.JumlahIbuHamil,
		&posyandu.Kecamatan_ID,
		&posyandu.Kecamatan,
		&posyandu.Kelurahan_ID,
		&posyandu.Kelurahan,
		&posyandu.Latitude,
		&posyandu.Longitude,
		&posyandu.SPPGID,
		&posyandu.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &posyandu, nil
}

func (m PosyanduModel) Update(posyandu *Posyandu) error {
	query := `
UPDATE posyandu
SET
	nama = $1,
	alamat = $2,
	kecamatan_id = $3,
	kelurahan_id = $4,
	latitude = $5,
	longitude = $6,
	jumlah_balita = $7,
	jumlah_ibu_hamil = $8,
	version = version + 1
WHERE id = $9 AND version = $10
RETURNING version`

	args := []any{
		posyandu.Nama,
		posyandu.Alamat,
		posyandu.Kecamatan_ID,
		posyandu.Kelurahan_ID,
		posyandu.Latitude,
		posyandu.Longitude,
		posyandu.JumlahBalita,
		posyandu.JumlahIbuHamil,
		posyandu.ID,
		posyandu.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&posyandu.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m PosyanduModel) Delete(id int64) error {
	// Return an ErrRecordNotFound error if the movie ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM posyandu
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
