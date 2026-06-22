package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"time"

	"github.com/lib/pq"
)

type Driver struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UserID       int64     `json:"user_id"`
	SPPGID       int64     `json:"sppg_id,omitempty"`
	SPPG         SPPG      `json:"sppg,omitzero"`
	Nama         string    `json:"nama"`
	NomorTelepon string    `json:"nomor_telepon"`
	StatusAktif  bool      `json:"status_aktif"`
	Version      int32     `json:"version"`
}

var (
	ErrDuplicatePhoneNumber = errors.New("duplicate phone number")
)

func ValidateDriver(v *validator.Validator, driver *Driver) {
	v.Check(driver.Nama != "", "nama", "must be provided")
	v.Check(
		len(driver.Nama) <= 200,
		"nama",
		"must not be more than 200 bytes long",
	)

	v.Check(driver.NomorTelepon != "", "nomor_telepon", "must be provided")
	v.Check(
		len(driver.NomorTelepon) <= 20,
		"nomor_telepon",
		"must not be more than 20 bytes long",
	)
}

type DriverModel struct {
	DB *sql.DB
}

func (m DriverModel) Insert(driver *Driver) error {
	query := `
		INSERT INTO driver (
    user_id,
    nama,
    nomor_telepon,
    status_aktif
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, version
	`

	args := []any{
		driver.UserID,
		driver.Nama,
		driver.NomorTelepon,
		driver.StatusAktif,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&driver.ID,
		&driver.CreatedAt,
		&driver.Version,
	)
}

func (m DriverModel) InsertTx(ctx context.Context, tx *sql.Tx, driver *Driver) error {
	query := `
		INSERT INTO drivers (
			user_id,
			sppg_id,
			nama,
			nomor_telepon,
			status_aktif
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, version
	`

	args := []any{
		driver.UserID,
		driver.SPPGID,
		driver.Nama,
		driver.NomorTelepon,
		driver.StatusAktif,
	}

	err := tx.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(
		&driver.ID,
		&driver.CreatedAt,
		&driver.Version,
	)
	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) {
			switch pqErr.Constraint {
			case "driver_nomor_telepon_key":
				return ErrDuplicatePhoneNumber
			}
		}

		return err
	}

	return nil
}

func (m DriverModel) GetAll(sppgID int64, statusAktif *bool, filters Filters) ([]*Driver, Metadata, error) {

	query := fmt.Sprintf(`
	SELECT count(*) OVER(),
		id,
		created_at,
		nama,
		nomor_telepon,
		status_aktif,
		user_id,
		sppg_id,
		version
	FROM drivers
	WHERE sppg_id = $1
	AND (status_aktif = $2 OR $2 IS NULL)
	ORDER BY %s %s, id ASC
	LIMIT $3 OFFSET $4`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{sppgID, statusAktif, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	drivers := []*Driver{}

	for rows.Next() {

		var driver Driver

		err := rows.Scan(
			&totalRecords,
			&driver.ID,
			&driver.CreatedAt,
			&driver.Nama,
			&driver.NomorTelepon,
			&driver.StatusAktif,
			&driver.UserID,
			&driver.SPPGID,
			&driver.Version,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		drivers = append(drivers, &driver)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return drivers, metadata, nil
}

func (m DriverModel) Get(id int64) (*Driver, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT
		id,
		created_at,
		nama,
		nomor_telepon,
		status_aktif,
		user_id,
		sppg_id,
		version
		FROM drivers
		WHERE id = $1
	`

	var driver Driver

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&driver.ID,
		&driver.CreatedAt,
		&driver.Nama,
		&driver.NomorTelepon,
		&driver.StatusAktif,
		&driver.UserID,
		&driver.SPPGID,
		&driver.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &driver, nil
}

func (m DriverModel) GetByUserID(user_id int64) (*Driver, error) {
	if user_id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT
		id,
		created_at,
		nama,
		nomor_telepon,
		status_aktif,
		user_id,
		sppg_id,
		version
		FROM drivers
		WHERE user_id = $1
	`

	var driver Driver

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, user_id).Scan(
		&driver.ID,
		&driver.CreatedAt,
		&driver.Nama,
		&driver.NomorTelepon,
		&driver.StatusAktif,
		&driver.UserID,
		&driver.SPPGID,
		&driver.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &driver, nil
}

func (m DriverModel) GetByUserIDJoinSPPG(user_id int64) (*Driver, error) {
	if user_id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT
    d.id,
    d.created_at,
    d.nama,
    d.nomor_telepon,
    d.status_aktif,
    d.user_id,
    d.version,

    s.id,
    s.nama,
    s.alamat
		FROM drivers d
		JOIN sppg s ON s.id = d.sppg_id
		WHERE d.user_id = $1
	`

	var driver Driver

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, user_id).Scan(
		&driver.ID,
		&driver.CreatedAt,
		&driver.Nama,
		&driver.NomorTelepon,
		&driver.StatusAktif,
		&driver.UserID,
		&driver.Version,

		&driver.SPPG.ID,
		&driver.SPPG.Nama,
		&driver.SPPG.Alamat,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &driver, nil
}

func (m DriverModel) Update(driver *Driver) error {
	query := `
UPDATE drivers
SET
	nama = $1,
	nomor_telepon = $2,
	status_aktif = $3,
	version = version + 1
WHERE id = $4 AND version = $5
RETURNING version`

	args := []any{
		driver.Nama,
		driver.NomorTelepon,
		driver.StatusAktif,
		driver.ID,
		driver.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&driver.Version)

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

func (m DriverModel) DeleteTx(ctx context.Context, tx *sql.Tx, id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `
	DELETE FROM drivers
	WHERE id = $1`

	result, err := tx.ExecContext(ctx, query, id)
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
