package data

import (
	"context"
	"database/sql"
	"errors"
	"mbg/internal/validator"
	"time"
)

func ValidateTracking(v *validator.Validator, tracking *Tracking) {
	v.Check(
		tracking.PengirimanID > 0,
		"pengiriman_id",
		"must be a valid pengiriman id",
	)

	v.Check(
		tracking.Latitude >= -90 && tracking.Latitude <= 90,
		"latitude",
		"must be between -90 and 90",
	)

	v.Check(
		tracking.Longitude >= -180 && tracking.Longitude <= 180,
		"longitude",
		"must be between -180 and 180",
	)

	v.Check(
		tracking.Speed >= 0,
		"speed",
		"must not be negative",
	)

	v.Check(
		tracking.Accuracy >= 0,
		"accuracy",
		"must not be negative",
	)

	v.Check(
		!(tracking.Latitude == 0 && tracking.Longitude == 0),
		"location",
		"must be a valid GPS coordinate",
	)
}

type TrackingModel struct {
	DB *sql.DB
}

type Tracking struct {
	ID           int64     `json:"id"`
	PengirimanID int64     `json:"pengiriman_id"`
	Latitude     float64   `json:"latitude"`
	Longitude    float64   `json:"longitude"`
	Speed        float64   `json:"speed"`
	Accuracy     float64   `json:"accuracy"`
	CreatedAt    time.Time `json:"created_at"`
	TujuanType   string    `json:"tujuan_type,omitempty"`
	TujuanNama   string    `json:"tujuan_nama,omitempty"`
	TujuanLat    float64   `json:"tujuan_lat,omitempty"`
	TujuanLng    float64   `json:"tujuan_lng,omitempty"`
	SPPGNama     *string   `json:"sppg_nama,omitempty"`
	DriverNama   *string   `json:"driver_nama,omitempty"`
}

func (m TrackingModel) Insert(tracking *Tracking) error {
	query := `
		INSERT INTO tracking (
			pengiriman_id,
			latitude,
			longitude,
			speed,
			accuracy
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	args := []any{
		tracking.PengirimanID,
		tracking.Latitude,
		tracking.Longitude,
		tracking.Speed,
		tracking.Accuracy,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&tracking.ID,
		&tracking.CreatedAt,
	)
}

type Tujuan struct {
	Nama      string  `json:"nama"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type TrackingGroup struct {
	PengirimanID int64       `json:"pengiriman_id"`
	Tujuan       Tujuan      `json:"tujuan"`
	Tracks       []*Tracking `json:"tracks"`
}

func (m TrackingModel) Get(pengiriman_id int64) (*Tracking, error) {

	query := `
	SELECT
    t.id,
    t.created_at,
    t.pengiriman_id,
    t.latitude,
    t.longitude,
    t.speed,
    t.accuracy,
		g.nama AS sppg_nama,
		d.nama AS driver_nama,

    CASE
        WHEN p.tujuan_type = 'sekolah'
            THEN s.nama
        WHEN p.tujuan_type = 'posyandu'
            THEN pos.nama
    END AS tujuan_nama,

    CASE
        WHEN p.tujuan_type = 'sekolah'
            THEN s.latitude
        WHEN p.tujuan_type = 'posyandu'
            THEN pos.latitude
    END AS tujuan_lat,

    CASE
        WHEN p.tujuan_type = 'sekolah'
            THEN s.longitude
        WHEN p.tujuan_type = 'posyandu'
            THEN pos.longitude
    END AS tujuan_lng

		FROM tracking t
		LEFT JOIN pengiriman p
				ON p.id = t.pengiriman_id

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

		WHERE t.id = $1
	`
	var tracking Tracking

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, pengiriman_id).Scan(
		&tracking.ID,
		&tracking.CreatedAt,
		&tracking.PengirimanID,
		&tracking.Latitude,
		&tracking.Longitude,
		&tracking.Speed,
		&tracking.Accuracy,
		&tracking.SPPGNama,
		&tracking.DriverNama,
		&tracking.TujuanNama,
		&tracking.TujuanLat,
		&tracking.TujuanLng,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &tracking, nil
}

func (m TrackingModel) GetAll(pengiriman_ids []int64, filters Filters) ([]*Tracking, Metadata, error) {

	query := `
	SELECT
    count(*) OVER(),
    t.id,
    t.created_at,
    t.pengiriman_id,
    t.latitude,
    t.longitude,
    t.speed,
    t.accuracy,
		g.nama AS sppg_nama,
		d.nama AS driver_nama,

    CASE
        WHEN p.tujuan_type = 'sekolah' THEN s.nama
        WHEN p.tujuan_type = 'posyandu' THEN pos.nama
    END AS tujuan_nama,

    CASE
        WHEN p.tujuan_type = 'sekolah' THEN s.latitude
        WHEN p.tujuan_type = 'posyandu' THEN pos.latitude
    END AS tujuan_lat,

    CASE
        WHEN p.tujuan_type = 'sekolah' THEN s.longitude
        WHEN p.tujuan_type = 'posyandu' THEN pos.longitude
    END AS tujuan_lng

		FROM (
				SELECT DISTINCT ON (pengiriman_id) *
				FROM tracking
				WHERE pengiriman_id = ANY($1::bigint[])
				ORDER BY pengiriman_id, created_at DESC
		) t

		LEFT JOIN pengiriman p
				ON p.id = t.pengiriman_id

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

		ORDER BY t.created_at DESC
		LIMIT $2 OFFSET $3
			`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{pengiriman_ids, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}
	defer rows.Close()

	totalRecords := 0

	tracking_all := []*Tracking{}

	for rows.Next() {

		var tracking Tracking

		err := rows.Scan(
			&totalRecords,
			&tracking.ID,
			&tracking.CreatedAt,
			&tracking.PengirimanID,
			&tracking.Latitude,
			&tracking.Longitude,
			&tracking.Speed,
			&tracking.Accuracy,
			&tracking.SPPGNama,
			&tracking.DriverNama,
			&tracking.TujuanNama,
			&tracking.TujuanLat,
			&tracking.TujuanLng,
		)
		if err != nil {
			return nil, Metadata{}, err
		}
		tracking_all = append(tracking_all, &tracking)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return tracking_all, metadata, nil
}

func (m TrackingModel) DeleteByPengirimanTx(ctx context.Context, tx *sql.Tx, id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM tracking
	WHERE pengiriman_id = $1`

	_, err := tx.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func nullString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

func nullFloat(nf sql.NullFloat64) float64 {
	if nf.Valid {
		return nf.Float64
	}
	return 0
}
