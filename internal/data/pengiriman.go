package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"
)

type PengirimanModel struct {
	DB *sql.DB
}

type Pengiriman struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`

	TujuanType      string  `json:"tujuan_type"`
	TujuanID        int64   `json:"tujuan_id"`
	TujuanNama      string  `json:"tujuan_nama,omitempty"`
	TujuanAlamat    string  `json:"tujuan_alamat,omitempty"`
	TujuanLatitude  float64 `json:"tujuan_latitude,omitempty"`
	TujuanLongitude float64 `json:"tujuan_longitude,omitempty"`

	DriverID   *int64  `json:"driver_id"`
	DriverNama *string `json:"driver_nama"`

	Status         string     `json:"status"`
	WaktuBerangkat *time.Time `json:"waktu_berangkat"`
	WaktuSelesai   *time.Time `json:"waktu_selesai"`

	SPPGID   int64   `json:"sppg_id"`
	SPPGNama *string `json:"sppg_nama,omitempty"`

	Version int32 `json:"version"`
}

type TujuanInput struct {
	TujuanType string `json:"tujuan_type"`
	TujuanID   int64  `json:"tujuan_id"`
}

type CreatePengirimanInput struct {
	Tujuan []TujuanInput `json:"tujuan"`
	Status string        `json:"status"`
	SPPGID int64         `json:"sppg_id"`
}

func ValidatePengiriman(v *validator.Validator, pengiriman *Pengiriman) {
	v.Check(
		pengiriman.SPPGID > 0,
		"sppg_id",
		"must be provided",
	)

	v.Check(
		pengiriman.TujuanType != "",
		"tujuan_type",
		"must be provided",
	)

	v.Check(
		validator.PermittedValue(
			pengiriman.TujuanType,
			"sekolah",
			"posyandu",
		),
		"tujuan_type",
		"must be either sekolah or posyandu",
	)

	v.Check(
		pengiriman.TujuanID > 0,
		"tujuan_id",
		"must be provided",
	)

	v.Check(
		pengiriman.Status != "",
		"status",
		"must be provided",
	)

	v.Check(
		validator.PermittedValue(
			pengiriman.Status,
			"menunggu",
			"berangkat",
			"sampai",
			"dibatalkan",
		),
		"status",
		"must be a valid status",
	)

	if pengiriman.WaktuBerangkat != nil &&
		pengiriman.WaktuSelesai != nil {
		v.Check(
			!pengiriman.WaktuSelesai.Before(*pengiriman.WaktuBerangkat),
			"waktu_selesai",
			"must be after waktu_berangkat",
		)
	}
}

func ValidateInputPengiriman(v *validator.Validator, input *CreatePengirimanInput) {
	v.Check(
		input.SPPGID > 0,
		"sppg_id",
		"must be provided",
	)

	v.Check(
		len(input.Tujuan) > 0,
		"tujuan",
		"must contain at least one destination",
	)

	v.Check(
		input.Status != "",
		"status",
		"must be provided",
	)

	v.Check(
		validator.PermittedValue(input.Status, "menunggu"),
		"status",
		"must be a valid status",
	)

	for i, tujuan := range input.Tujuan {
		v.Check(
			tujuan.TujuanID > 0,
			fmt.Sprintf("tujuan[%d].tujuan_id", i),
			"must be provided",
		)

		v.Check(
			validator.PermittedValue(
				tujuan.TujuanType,
				"sekolah",
				"posyandu",
			),
			fmt.Sprintf("tujuan[%d].tujuan_type", i),
			"must be either sekolah or posyandu",
		)
	}
}

func ValidatePengirimanUpdate(v *validator.Validator, pengiriman *Pengiriman) {
	v.Check(
		pengiriman.SPPGID > 0,
		"sppg_id",
		"must be provided",
	)

	v.Check(
		pengiriman.TujuanType != "",
		"tujuan_type",
		"must be provided",
	)

	v.Check(
		validator.PermittedValue(
			pengiriman.TujuanType,
			"sekolah",
			"posyandu",
		),
		"tujuan_type",
		"must be either sekolah or posyandu",
	)

	v.Check(
		pengiriman.TujuanID > 0,
		"tujuan_id",
		"must be provided",
	)

	v.Check(
		pengiriman.Status != "",
		"status",
		"must be provided",
	)

	v.Check(
		validator.PermittedValue(
			pengiriman.Status,
			"menunggu",
			"berangkat",
			"sampai",
			"dibatalkan",
		),
		"status",
		"must be a valid status",
	)

	if pengiriman.WaktuBerangkat != nil &&
		pengiriman.WaktuSelesai != nil {
		v.Check(
			!pengiriman.WaktuSelesai.Before(*pengiriman.WaktuBerangkat),
			"waktu_selesai",
			"must be after waktu_berangkat",
		)
	}
}

func (m PengirimanModel) Insert(pengiriman *Pengiriman) error {
	query := `
INSERT INTO pengiriman (
    sppg_id,
    tujuan_type,
    tujuan_id,
    status,
    waktu_berangkat,
    waktu_selesai
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, version`

	args := []any{
		pengiriman.SPPGID,
		pengiriman.TujuanType,
		pengiriman.TujuanID,
		pengiriman.Status,
		pengiriman.WaktuBerangkat,
		pengiriman.WaktuSelesai,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&pengiriman.ID,
		&pengiriman.CreatedAt,
		&pengiriman.Version,
	)
}

func (m PengirimanModel) InsertTx(ctx context.Context, tx *sql.Tx, pengiriman *Pengiriman) error {
	query := `
INSERT INTO pengiriman (
    sppg_id,
    tujuan_type,
    tujuan_id,
    status,
    waktu_berangkat,
    waktu_selesai
)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, version`

	args := []any{
		pengiriman.SPPGID,
		pengiriman.TujuanType,
		pengiriman.TujuanID,
		pengiriman.Status,
		pengiriman.WaktuBerangkat,
		pengiriman.WaktuSelesai,
	}

	err := tx.QueryRowContext(ctx, query, args...).Scan(
		&pengiriman.ID,
		&pengiriman.CreatedAt,
		&pengiriman.Version,
	)
	if err != nil {
		return err
	}
	return nil
}

func (m PengirimanModel) GetAll(status string, sppg_id int64, start *time.Time, end *time.Time, tujuan_type string, filters Filters) ([]*Pengiriman, Metadata, error) {
	query := fmt.Sprintf(`
	SELECT
		count(*) OVER(),
		p.id,
		p.created_at,
		p.tujuan_type,
		p.tujuan_id,
		p.driver_id,
		d.nama AS driver_nama,
		p.status,
		p.waktu_berangkat,
		p.waktu_selesai,
		p.sppg_id,
		g.nama AS sppg_nama,
		p.version,

		CASE
				WHEN p.tujuan_type = 'sekolah'
						THEN COALESCE(s.nama, '')
				WHEN p.tujuan_type = 'posyandu'
						THEN COALESCE(pos.nama, '')
				ELSE ''
		END AS tujuan_nama,

		CASE
				WHEN p.tujuan_type = 'sekolah'
						THEN COALESCE(s.alamat, '')
				WHEN p.tujuan_type = 'posyandu'
						THEN COALESCE(pos.alamat, '')
				ELSE ''
		END AS tujuan_alamat

	FROM pengiriman p

	LEFT JOIN drivers d
		ON d.id = p.driver_id

	LEFT JOIN sekolah s
		ON p.tujuan_type = 'sekolah'
		AND s.id = p.tujuan_id

	LEFT JOIN posyandu pos
		ON p.tujuan_type = 'posyandu'
		AND pos.id = p.tujuan_id

	LEFT JOIN sppg g
		ON p.sppg_id = g.id

	WHERE (p.sppg_id = $1 OR $1 = 0)
	AND (LOWER(p.status) = LOWER($2) OR $2 = '')
	AND (LOWER(p.tujuan_type) = LOWER($3) OR $3 = '')
	AND (
    $4::timestamp IS NULL
    OR (
        p.created_at >= $4::timestamp
        AND p.created_at < $5::timestamp
			)
	)

	ORDER BY %s %s, p.id ASC
	LIMIT $6 OFFSET $7`,
		filters.sortColumn(),
		filters.sortDirection(),
	)
	// WHERE LOWER(status) = LOWER($1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{sppg_id, status, tujuan_type, start, end, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	pengiriman_all := []*Pengiriman{}

	for rows.Next() {

		var pengiriman Pengiriman

		err := rows.Scan(
			&totalRecords,
			&pengiriman.ID,
			&pengiriman.CreatedAt,
			&pengiriman.TujuanType,
			&pengiriman.TujuanID,
			&pengiriman.DriverID,
			&pengiriman.DriverNama,
			&pengiriman.Status,
			&pengiriman.WaktuBerangkat,
			&pengiriman.WaktuSelesai,
			&pengiriman.SPPGID,
			&pengiriman.SPPGNama,
			&pengiriman.Version,
			&pengiriman.TujuanNama,
			&pengiriman.TujuanAlamat,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		pengiriman_all = append(pengiriman_all, &pengiriman)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return pengiriman_all, metadata, nil
}

func (m PengirimanModel) Get(id int64) (*Pengiriman, error) {
	query := `
	SELECT id, created_at, tujuan_type, tujuan_id, status, waktu_berangkat, waktu_selesai, sppg_id, driver_id, version
	FROM pengiriman
	WHERE id = $1`

	var pengiriman Pengiriman

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&pengiriman.ID,
		&pengiriman.CreatedAt,
		&pengiriman.TujuanType,
		&pengiriman.TujuanID,
		&pengiriman.Status,
		&pengiriman.WaktuBerangkat,
		&pengiriman.WaktuSelesai,
		&pengiriman.SPPGID,
		&pengiriman.DriverID,
		&pengiriman.Version,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &pengiriman, nil
}

func (m PengirimanModel) GetAktifByDriverID(driverID int64) (*Pengiriman, error) {
	// fmt.Println(driverID)
	query := `
	SELECT 
		p.id,
	 	p.created_at,
		p.tujuan_type, 
		p.tujuan_id, 
		p.driver_id, 
		d.nama AS driver_nama, 
		p.status, 
		p.waktu_berangkat, 
		p.waktu_selesai, 
		p.sppg_id, 
		p.version,

	CASE
    WHEN p.tujuan_type = 'sekolah' THEN s.nama
    WHEN p.tujuan_type = 'posyandu' THEN pos.nama
	END AS tujuan_nama,
	
	CASE
    WHEN p.tujuan_type = 'sekolah' THEN s.alamat
    WHEN p.tujuan_type = 'posyandu' THEN pos.alamat
	END AS tujuan_alamat,

	CASE
			WHEN p.tujuan_type = 'sekolah' THEN s.latitude
			WHEN p.tujuan_type = 'posyandu' THEN pos.latitude
	END AS tujuan_latitude,

	CASE
			WHEN p.tujuan_type = 'sekolah' THEN s.longitude
			WHEN p.tujuan_type = 'posyandu' THEN pos.longitude
	END AS tujuan_longitude

	FROM pengiriman p

	LEFT JOIN drivers d
		ON d.id = p.driver_id

	LEFT JOIN sekolah s
		ON p.tujuan_type = 'sekolah'
		AND s.id = p.tujuan_id

	LEFT JOIN posyandu pos
		ON p.tujuan_type = 'posyandu'
		AND pos.id = p.tujuan_id

	WHERE driver_id = $1
	AND status = 'berangkat'
	`

	var pengiriman Pengiriman

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, driverID).Scan(
		&pengiriman.ID,
		&pengiriman.CreatedAt,
		&pengiriman.TujuanType,
		&pengiriman.TujuanID,
		&pengiriman.DriverID,
		&pengiriman.DriverNama,
		&pengiriman.Status,
		&pengiriman.WaktuBerangkat,
		&pengiriman.WaktuSelesai,
		&pengiriman.SPPGID,
		&pengiriman.Version,
		&pengiriman.TujuanNama,
		&pengiriman.TujuanAlamat,
		&pengiriman.TujuanLatitude,
		&pengiriman.TujuanLongitude,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, nil
		default:
			return nil, err
		}
	}

	return &pengiriman, nil
}

func (m PengirimanModel) GetIsDelivering(driverID int64) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM pengiriman
			WHERE driver_id = $1
			AND status = 'berangkat'
		)
`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var exists bool

	err := m.DB.QueryRowContext(ctx, query, driverID).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (m PengirimanModel) CountByDriverID(id int64) (int, error) {
	query := `
	SELECT COUNT(*)
	FROM pengiriman
	WHERE driver_id = $1`

	var count int

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (m PengirimanModel) Update(pengiriman *Pengiriman) error {
	query := `
	UPDATE pengiriman
	SET status = $1, waktu_berangkat = $2, waktu_selesai = $3, driver_id = $4, version = version + 1
	WHERE id = $5 AND version = $6
	RETURNING version`

	args := []any{
		pengiriman.Status,
		pengiriman.WaktuBerangkat,
		pengiriman.WaktuSelesai,
		pengiriman.DriverID,
		pengiriman.ID,
		pengiriman.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&pengiriman.Version)
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

func (m PengirimanModel) UpdateTx(ctx context.Context, tx *sql.Tx, pengiriman *Pengiriman) error {
	query := `
	UPDATE pengiriman
	SET status = $1, waktu_berangkat = $2, waktu_selesai = $3, driver_id = $4, version = version + 1
	WHERE id = $5 AND version = $6
	RETURNING version`

	args := []any{
		pengiriman.Status,
		pengiriman.WaktuBerangkat,
		pengiriman.WaktuSelesai,
		pengiriman.DriverID,
		pengiriman.ID,
		pengiriman.Version,
	}

	err := tx.QueryRowContext(ctx, query, args...).Scan(&pengiriman.Version)
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
