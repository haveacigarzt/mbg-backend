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

type SPPG struct {
	ID             int64     `json:"id,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UserID         int64     `json:"user_id,omitempty"`
	Nama           string    `json:"nama,omitempty"`
	Alamat         string    `json:"alamat,omitempty"`
	SosmedURL      []string  `json:"sosmed_url,omitempty"`
	KepalaSPPG     string    `json:"kepala_sppg,omitempty"`
	NomorTelepon   string    `json:"nomor_telepon,omitempty"`
	Email          string    `json:"email,omitempty"`
	Latitude       float64   `json:"latitude,omitempty"`
	Longitude      float64   `json:"longitude,omitempty"`
	Kecamatan      string    `json:"kecamatan,omitempty"`
	Kelurahan      string    `json:"kelurahan,omitempty"`
	Kecamatan_ID   int64     `json:"kecamatan_id,omitempty"`
	Kelurahan_ID   int64     `json:"kelurahan_id,omitempty"`
	KapasitasPorsi int       `json:"kapasitas_porsi,omitempty"`
	StatusAktif    bool      `json:"status_aktif,omitempty"`
	Version        int32     `json:"version,omitempty"`
}

func ValidateSPPG(v *validator.Validator, sppg *SPPG) {
	v.Check(sppg.Nama != "", "nama", "must be provided")
	v.Check(len(sppg.Nama) <= 200, "nama", "must not be more than 200 bytes long")

	v.Check(sppg.Alamat != "", "alamat", "must be provided")
	v.Check(len(sppg.Alamat) <= 500, "alamat", "must not be more than 500 bytes long")

	v.Check(sppg.KepalaSPPG != "", "kepala_sppg", "must be provided")
	v.Check(
		len(sppg.KepalaSPPG) <= 200,
		"kepala_sppg",
		"must not be more than 200 bytes long",
	)

	v.Check(sppg.Kecamatan_ID != 0, "kecamatan", "must be provided")

	v.Check(sppg.Kelurahan_ID != 0, "kelurahan", "must be provided")

	v.Check(
		sppg.NomorTelepon == "" || len(sppg.NomorTelepon) <= 20,
		"nomor_telepon",
		"must not be more than 20 bytes long",
	)

	v.Check(
		sppg.Email == "" || validator.Matches(sppg.Email, validator.EmailRX),
		"email",
		"must be a valid email address",
	)

	v.Check(
		sppg.KapasitasPorsi >= 0,
		"kapasitas_porsi",
		"must be greater than or equal to 0",
	)

	v.Check(
		sppg.Latitude >= -90 && sppg.Latitude <= 90,
		"latitude",
		"must be between -90 and 90",
	)

	v.Check(
		sppg.Longitude >= -180 && sppg.Longitude <= 180,
		"longitude",
		"must be between -180 and 180",
	)

	for _, url := range sppg.SosmedURL {
		v.Check(
			validator.IsValidURL(url),
			"sosmed_url",
			"must contain valid URLs",
		)
	}
}

// func ValidateSPPGForUpdate(v *validator.Validator, sppg *SPPG) {

// 	if sppg.Nama != "" {
// 		v.Check(len(sppg.Nama) <= 200, "nama", "must not be more than 200 bytes long")
// 	}

// 	if sppg.Alamat != "" {
// 		v.Check(len(sppg.Alamat) <= 500, "alamat", "must not be more than 500 bytes long")
// 	}

// 	if sppg.KepalaSPPG != "" {
// 		v.Check(len(sppg.KepalaSPPG) <= 200, "kepala_sppg", "must not be more than 200 bytes long")
// 	}

// 	if sppg.Kecamatan != "" {
// 		v.Check(len(sppg.Kecamatan) <= 100, "kecamatan", "must not be more than 100 bytes long")
// 	}

// 	if sppg.Kelurahan != "" {
// 		v.Check(len(sppg.Kelurahan) <= 100, "kelurahan", "must not be more than 100 bytes long")
// 	}

// 	if sppg.NomorTelepon != "" {
// 		v.Check(len(sppg.NomorTelepon) <= 20, "nomor_telepon", "must not be more than 20 bytes long")
// 	}

// 	if sppg.Email != "" {
// 		v.Check(
// 			len(sppg.Email) <= 255,
// 			"email",
// 			"must not be more than 255 bytes long",
// 		)

// 		v.Check(
// 			validator.Matches(sppg.Email, validator.EmailRX),
// 			"email",
// 			"must be a valid email address",
// 		)
// 	}

// 	if sppg.KapasitasPorsi != 0 {
// 		v.Check(
// 			sppg.KapasitasPorsi >= 0,
// 			"kapasitas_porsi",
// 			"must be greater than or equal to 0",
// 		)
// 	}

// 	// Latitude validasi hanya jika dikirim (hindari 0 dianggap valid input)
// 	if sppg.Latitude != 0 {
// 		v.Check(
// 			sppg.Latitude >= -90 && sppg.Latitude <= 90,
// 			"latitude",
// 			"must be between -90 and 90",
// 		)
// 	}

// 	if sppg.Longitude != 0 {
// 		v.Check(
// 			sppg.Longitude >= -180 && sppg.Longitude <= 180,
// 			"longitude",
// 			"must be between -180 and 180",
// 		)
// 	}

// 	if sppg.SosmedURL != nil {
// 		for _, url := range sppg.SosmedURL {
// 			v.Check(
// 				validator.IsValidURL(url),
// 				"sosmed_url",
// 				"must be a valid URL",
// 			)
// 		}
// 	}
// }

type SPPGModel struct {
	DB *sql.DB
}

func (m SPPGModel) Insert(sppg *SPPG) error {
	query := `
		INSERT INTO sppg (
    user_id,
    nama,
    alamat,
    sosmed_url,
    kepala_sppg,
    nomor_telepon,
    email,
    latitude,
    longitude,
    kecamatan_id,
    kelurahan_id,
    kapasitas_porsi,
    status_aktif
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
		RETURNING id, created_at, version
	`

	args := []any{
		sppg.UserID,
		sppg.Nama,
		sppg.Alamat,
		sppg.SosmedURL,
		sppg.KepalaSPPG,
		sppg.NomorTelepon,
		sppg.Email,
		sppg.Latitude,
		sppg.Longitude,
		sppg.Kecamatan_ID,
		sppg.Kelurahan_ID,
		sppg.KapasitasPorsi,
		sppg.StatusAktif,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&sppg.ID,
		&sppg.CreatedAt,
		&sppg.Version,
	)
}

func (m SPPGModel) GetAll(nama string, kecamatan_id int64, kelurahan_id int64, status_aktif *bool, filters Filters) ([]*SPPG, Metadata, error) {

	query := fmt.Sprintf(`
	SELECT count(*) OVER(),
		s.id,
		s.created_at,
		s.nama,
		s.alamat,
		s.kepala_sppg,
		s.nomor_telepon,
		s.email,
		s.kapasitas_porsi,
		s.kecamatan_id,
		k.name AS kecamatan,
		s.kelurahan_id,
		kel.name AS kelurahan,
		s.sosmed_url,
		s.status_aktif,
		s.latitude,
		s.longitude
	FROM sppg s
	LEFT JOIN kecamatan k ON k.id = s.kecamatan_id
	LEFT JOIN kelurahan kel ON kel.id = s.kelurahan_id
	WHERE (LOWER(s.nama) LIKE LOWER('%%' || $1 || '%%') OR $1 = '')
	AND (s.kecamatan_id = $2 OR $2 = 0)
	AND (s.kelurahan_id = $3 OR $3 = 0)
	AND (s.status_aktif = $4 OR $4 IS NULL)
	ORDER BY %s %s, s.id ASC
	LIMIT $5 OFFSET $6`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	args := []any{
		nama,
		kecamatan_id,
		kelurahan_id,
		status_aktif,
		filters.limit(),
		filters.offset(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	sppg_all := []*SPPG{}

	for rows.Next() {

		var sppg SPPG

		err := rows.Scan(
			&totalRecords,
			&sppg.ID,
			&sppg.CreatedAt,
			&sppg.Nama,
			&sppg.Alamat,
			&sppg.KepalaSPPG,
			&sppg.NomorTelepon,
			&sppg.Email,
			&sppg.KapasitasPorsi,
			&sppg.Kecamatan_ID,
			&sppg.Kecamatan,
			&sppg.Kelurahan_ID,
			&sppg.Kelurahan,
			pq.Array(&sppg.SosmedURL),
			&sppg.StatusAktif,
			&sppg.Latitude,
			&sppg.Longitude,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		sppg_all = append(sppg_all, &sppg)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return sppg_all, metadata, nil
}

func (m SPPGModel) Get(id int64) (*SPPG, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT
			s.id,
			s.created_at,
			s.nama,
			s.alamat,
			s.kepala_sppg,
			s.nomor_telepon,
			s.email,
			s.kapasitas_porsi,
			s.kecamatan_id,
			k.name AS kecamatan,
			s.kelurahan_id,
			kel.name AS kelurahan,
			s.sosmed_url,
			s.status_aktif,
			s.latitude,
			s.longitude,
			s.user_id,
			s.version
			FROM sppg s
			LEFT JOIN kecamatan k ON k.id = s.kecamatan_id
			LEFT JOIN kelurahan kel ON kel.id = s.kelurahan_id
		WHERE s.id = $1
	`

	var sppg SPPG

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&sppg.ID,
		&sppg.CreatedAt,
		&sppg.Nama,
		&sppg.Alamat,
		&sppg.KepalaSPPG,
		&sppg.NomorTelepon,
		&sppg.Email,
		&sppg.KapasitasPorsi,
		&sppg.Kecamatan_ID,
		&sppg.Kecamatan,
		&sppg.Kelurahan_ID,
		&sppg.Kelurahan,
		pq.Array(&sppg.SosmedURL),
		&sppg.StatusAktif,
		&sppg.Latitude,
		&sppg.Longitude,
		&sppg.UserID,
		&sppg.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &sppg, nil
}

func (m SPPGModel) GetByUserID(user_id int64) (*SPPG, error) {
	if user_id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
		SELECT
			id,
			created_at,
			nama,
			alamat,
			kepala_sppg,
			nomor_telepon,
			email,
			kapasitas_porsi,
			kecamatan_id,
			kelurahan_id,
			sosmed_url,
			status_aktif,
			latitude,
			longitude,
			user_id,
			version
		FROM sppg
		WHERE user_id = $1
	`

	var sppg SPPG

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, user_id).Scan(
		&sppg.ID,
		&sppg.CreatedAt,
		&sppg.Nama,
		&sppg.Alamat,
		&sppg.KepalaSPPG,
		&sppg.NomorTelepon,
		&sppg.Email,
		&sppg.KapasitasPorsi,
		&sppg.Kecamatan_ID,
		&sppg.Kelurahan_ID,
		pq.Array(&sppg.SosmedURL),
		&sppg.StatusAktif,
		&sppg.Latitude,
		&sppg.Longitude,
		&sppg.UserID,
		&sppg.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &sppg, nil
}

func (m SPPGModel) Update(sppg *SPPG) error {
	query := `
UPDATE sppg
SET
	nama = $1,
    alamat = $2,
    sosmed_url = $3,
    kepala_sppg = $4,
    nomor_telepon = $5,
    email = $6,
    latitude = $7,
    longitude = $8,
    kecamatan_id = $9,
    kelurahan_id = $10,
    kapasitas_porsi = $11,
    status_aktif = $12,
	version = version + 1
WHERE id = $13 AND version = $14
RETURNING version`

	args := []any{
		sppg.Nama,
		sppg.Alamat,
		sppg.SosmedURL,
		sppg.KepalaSPPG,
		sppg.NomorTelepon,
		sppg.Email,
		sppg.Latitude,
		sppg.Longitude,
		sppg.Kecamatan_ID,
		sppg.Kelurahan_ID,
		sppg.KapasitasPorsi,
		sppg.StatusAktif,
		sppg.ID,
		sppg.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&sppg.Version)
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
