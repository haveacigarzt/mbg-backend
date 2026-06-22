package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type Sekolah struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"-"`
	Nama         string    `json:"nama"`
	Alamat       string    `json:"alamat"`
	Tingkat      string    `json:"tingkat"`
	JumlahSiswa  int       `json:"jumlah_siswa"`
	Kecamatan    string    `json:"kecamatan,omitempty"`
	Kelurahan    string    `json:"kelurahan,omitempty"`
	Kecamatan_ID int64     `json:"kecamatan_id"`
	Kelurahan_ID int64     `json:"kelurahan_id"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Version      int32     `json:"version"`
	SPPGID       int64     `json:"sppg_id"`
}

func ValidateSekolah(v *validator.Validator, sekolah *Sekolah) {
	v.Check(sekolah.Nama != "", "nama", "must be provided")
	v.Check(len(sekolah.Nama) <= 200, "nama", "must not be more than 200 bytes long")

	v.Check(sekolah.Alamat != "", "alamat", "must be provided")
	v.Check(len(sekolah.Alamat) <= 500, "alamat", "must not be more than 500 bytes long")

	v.Check(sekolah.Tingkat != "", "tingkat", "must be provided")
	v.Check(
		validator.PermittedValue(sekolah.Tingkat, "SD", "SMP", "SMA"),
		"tingkat",
		"must be one of SD, SMP, or SMA",
	)

	v.Check(sekolah.JumlahSiswa != 0, "jumlah_siswa", "must be provided")
	v.Check(sekolah.JumlahSiswa > 0,
		"jumlah_siswa",
		"must be greater than 0",
	)

	v.Check(sekolah.Kecamatan_ID != 0, "kecamatan", "must be provided")

	v.Check(sekolah.Kelurahan_ID != 0, "kelurahan", "must be provided")

	v.Check(
		sekolah.Latitude >= -90 && sekolah.Latitude <= 90,
		"latitude",
		"must be between -90 and 90",
	)

	v.Check(
		sekolah.Longitude >= -180 && sekolah.Longitude <= 180,
		"longitude",
		"must be between -180 and 180",
	)
}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type SekolahModel struct {
	DB *sql.DB
}

// Add a placeholder method for inserting a new record in the sekolah table.
func (m SekolahModel) Insert(sekolah *Sekolah) error {
	// Define the SQL query for inserting a new record in the movies table and returning
	// the system-generated data.
	query := `
INSERT INTO sekolah (
    nama,
    alamat,
    tingkat,
    jumlah_siswa,
    kecamatan_id,
    kelurahan_id,
    latitude,
    longitude,
		sppg_id
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, created_at, version`

	args := []any{
		sekolah.Nama,
		sekolah.Alamat,
		sekolah.Tingkat,
		sekolah.JumlahSiswa,
		sekolah.Kecamatan_ID,
		sekolah.Kelurahan_ID,
		sekolah.Latitude,
		sekolah.Longitude,
		sekolah.SPPGID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&sekolah.ID,
		&sekolah.CreatedAt,
		&sekolah.Version,
	)
}

// Add a placeholder method for fetching a specific record from the sekolah table.
func (m SekolahModel) Get(id int64) (*Sekolah, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT
			s.id,
			s.created_at,
			s.nama,
			s.alamat,
			s.tingkat,
			s.jumlah_siswa,
			s.kecamatan_id,
			k.name AS kecamatan,
			s.kelurahan_id,
			kel.name AS kelurahan,
			s.latitude,
			s.longitude,
			s.sppg_id,
			s.version
		FROM sekolah s
		LEFT JOIN kecamatan k ON k.id = s.kecamatan_id
		LEFT JOIN kelurahan kel ON kel.id = s.kelurahan_id
		WHERE s.id = $1
	`

	var sekolah Sekolah

	// Use the context.WithTimeout() function to create a context.Context which carries a
	// 3-second timeout deadline. Note that we're using the empty context.Background()
	// as the 'parent' context.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&sekolah.ID,
		&sekolah.CreatedAt,
		&sekolah.Nama,
		&sekolah.Alamat,
		&sekolah.Tingkat,
		&sekolah.JumlahSiswa,
		&sekolah.Kecamatan_ID,
		&sekolah.Kecamatan,
		&sekolah.Kelurahan_ID,
		&sekolah.Kelurahan,
		&sekolah.Latitude,
		&sekolah.Longitude,
		&sekolah.SPPGID,
		&sekolah.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &sekolah, nil
}

// Add a placeholder method for updating a specific record in the sekolah table.
func (m SekolahModel) Update(sekolah *Sekolah) error {
	query := `
UPDATE sekolah
SET
	nama = $1,
	alamat = $2,
	tingkat = $3,
	jumlah_siswa = $4,
	kecamatan_id = $5,
	kelurahan_id = $6,
	latitude = $7,
	longitude = $8,
	version = version + 1
WHERE id = $9 AND version = $10
RETURNING version`

	args := []any{
		sekolah.Nama,
		sekolah.Alamat,
		sekolah.Tingkat,
		sekolah.JumlahSiswa,
		sekolah.Kecamatan_ID,
		sekolah.Kelurahan_ID,
		sekolah.Latitude,
		sekolah.Longitude,
		sekolah.ID,
		sekolah.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&sekolah.Version)
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

func (m SekolahModel) Delete(id int64) error {
	// Return an ErrRecordNotFound error if the movie ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM sekolah
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

func (m SekolahModel) GetAll(nama string, tingkat string, kecamatan_id int64, kelurahan_id int64, sppg_id int64, jumlah_siswa int, filters Filters) ([]*Sekolah, Metadata, error) {

	query := fmt.Sprintf(`
	SELECT count(*) OVER(),
		s.id,
		s.created_at,
		s.nama,
		s.alamat,
		s.tingkat,
		s.jumlah_siswa,
		s.kecamatan_id,
		k.name AS kecamatan,
		s.kelurahan_id,
		kel.name AS kelurahan,
		s.latitude,
		s.longitude,
		s.sppg_id,
		s.version
	FROM sekolah s
	LEFT JOIN kecamatan k ON k.id = s.kecamatan_id
	LEFT JOIN kelurahan kel ON kel.id = s.kelurahan_id
	WHERE (LOWER(s.nama) LIKE LOWER('%%' || $1 || '%%') OR $1 = '')
	AND (LOWER(s.tingkat) = LOWER($2) OR $2 = '')
	AND (s.kecamatan_id = $3 OR $3 = 0)
	AND (s.kelurahan_id = $4 OR $4 = 0)
	AND (s.sppg_id = $5 OR $5 = 0)
	AND (s.jumlah_siswa = $6 OR $6 = 0)
	ORDER BY %s %s, id ASC
	LIMIT $7 OFFSET $8`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{nama, tingkat, kecamatan_id, kelurahan_id, sppg_id, jumlah_siswa, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	sekolah_all := []*Sekolah{}

	for rows.Next() {

		var sekolah Sekolah

		err := rows.Scan(
			&totalRecords,
			&sekolah.ID,
			&sekolah.CreatedAt,
			&sekolah.Nama,
			&sekolah.Alamat,
			&sekolah.Tingkat,
			&sekolah.JumlahSiswa,
			&sekolah.Kecamatan_ID,
			&sekolah.Kecamatan,
			&sekolah.Kelurahan_ID,
			&sekolah.Kelurahan,
			&sekolah.Latitude,
			&sekolah.Longitude,
			&sekolah.SPPGID,
			&sekolah.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		sekolah_all = append(sekolah_all, &sekolah)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return sekolah_all, metadata, nil
}
