package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type PedagangLokal struct {
	ID          int64     `json:"id"`
	Nama        string    `json:"nama"`
	Alamat      string    `json:"alamat"`
	NoHP        string    `json:"no_hp"`
	Longitude   float64   `json:"longitude"`
	Latitude    float64   `json:"latitude"`
	JenisProduk string    `json:"jenis_produk"`
	SPPGID      int64     `json:"sppg_id"`
	SPPGNama    *string   `json:"sppg_nama,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     int32     `json:"version"`
}

func ValidatePedagangLokal(v *validator.Validator, input *PedagangLokal) {
	v.Check(input.Nama != "", "nama", "harus diisi")
	v.Check(len(input.Nama) <= 255, "nama", "maksimal 255 karakter")

	v.Check(input.Alamat != "", "alamat", "harus diisi")
	v.Check(len(input.Alamat) <= 255, "alamat", "maksimal 255 karakter")

	v.Check(input.NoHP != "", "no_hp", "harus diisi")
	v.Check(len(input.NoHP) <= 20, "no_hp", "maksimal 20 karakter")

	v.Check(input.JenisProduk != "", "jenis_produk", "harus diisi")
	v.Check(len(input.JenisProduk) <= 100, "jenis_produk", "maksimal 100 karakter")

	v.Check(input.Latitude >= -90 && input.Latitude <= 90,
		"latitude", "harus berada di antara -90 dan 90")

	v.Check(input.Longitude >= -180 && input.Longitude <= 180,
		"longitude", "harus berada di antara -180 dan 180")
}

type PedagangLokalModel struct {
	DB *sql.DB
}

func (m PedagangLokalModel) GetAll(
	nama string,
	jenisProduk string,
	sppg string,
	filters Filters,
) ([]*PedagangLokal, Metadata, error) {

	query := fmt.Sprintf(`
	SELECT
		count(*) OVER(),
		p.id,
		p.created_at,
		p.nama,
		p.alamat,
		p.no_hp,
		p.longitude,
		p.latitude,
		p.jenis_produk,
		p.sppg_id,
		s.nama,
		p.updated_at,
		p.version
	FROM pedagang_lokal p
	LEFT JOIN sppg s ON s.id = p.sppg_id
	WHERE
		(LOWER(p.nama) LIKE LOWER('%%' || $1 || '%%') OR $1 = '')
	AND (LOWER(p.jenis_produk) LIKE LOWER('%%' || $2 || '%%') OR $2 = '')
	AND (LOWER(s.nama) LIKE LOWER('%%' || $3 || '%%') OR $3 = '')
	ORDER BY %s %s, p.id ASC
	LIMIT $4 OFFSET $5`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{
		nama,
		jenisProduk,
		sppg,
		filters.limit(),
		filters.offset(),
	}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0
	pedagang := []*PedagangLokal{}

	for rows.Next() {
		var p PedagangLokal

		err := rows.Scan(
			&totalRecords,
			&p.ID,
			&p.CreatedAt,
			&p.Nama,
			&p.Alamat,
			&p.NoHP,
			&p.Longitude,
			&p.Latitude,
			&p.JenisProduk,
			&p.SPPGID,
			&p.SPPGNama,
			&p.UpdatedAt,
			&p.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pedagang = append(pedagang, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pedagang, metadata, nil
}

func (m PedagangLokalModel) Get(id int64) (*PedagangLokal, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
	SELECT
		p.id,
		p.created_at,
		p.nama,
		p.alamat,
		p.no_hp,
		p.longitude,
		p.latitude,
		p.jenis_produk,
		p.sppg_id,
		s.nama,
		p.updated_at,
		p.version
	FROM pedagang_lokal p
	LEFT JOIN sppg s ON s.id = p.sppg_id
	WHERE p.id = $1`

	var pedagangLokal PedagangLokal

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&pedagangLokal.ID,
		&pedagangLokal.CreatedAt,
		&pedagangLokal.Nama,
		&pedagangLokal.Alamat,
		&pedagangLokal.NoHP,
		&pedagangLokal.Longitude,
		&pedagangLokal.Latitude,
		&pedagangLokal.JenisProduk,
		&pedagangLokal.SPPGID,
		&pedagangLokal.SPPGNama,
		&pedagangLokal.UpdatedAt,
		&pedagangLokal.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &pedagangLokal, nil

}

func (m PedagangLokalModel) Insert(pedagangLokal *PedagangLokal) error {
	query := `
		INSERT INTO pedagang_lokal (
			nama,
			alamat,
			no_hp,
			longitude,
			latitude,
			jenis_produk,
			sppg_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, version
	`

	args := []any{
		pedagangLokal.Nama,
		pedagangLokal.Alamat,
		pedagangLokal.NoHP,
		pedagangLokal.Longitude,
		pedagangLokal.Latitude,
		pedagangLokal.JenisProduk,
		pedagangLokal.SPPGID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&pedagangLokal.ID,
		&pedagangLokal.CreatedAt,
		&pedagangLokal.Version,
	)
}

func (m PedagangLokalModel) Update(pedagangLokal *PedagangLokal) error {
	query := `
		UPDATE pedagang_lokal SET 
			nama = $1,
			alamat= $2,
			no_hp= $3,
			longitude= $4,
			latitude= $5,
			jenis_produk= $6,
			version = version + 1
		WHERE id = $7 AND version = $8
		RETURNING version`

	args := []any{
		pedagangLokal.Nama,
		pedagangLokal.Alamat,
		pedagangLokal.NoHP,
		pedagangLokal.Longitude,
		pedagangLokal.Latitude,
		pedagangLokal.JenisProduk,
		pedagangLokal.ID,
		pedagangLokal.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&pedagangLokal.Version)

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

func (m PedagangLokalModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
	DELETE FROM pedagang_lokal
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
