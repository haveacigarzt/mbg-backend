package data

import (
	"context"
	"database/sql"
	"time"
)

type Kecamatan struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
}

type KecamatanModel struct {
	DB *sql.DB
}

func (m KecamatanModel) GetAll() ([]*Kecamatan, error) {
	query := `
		SELECT id, created_at, name
		FROM kecamatan
		ORDER BY name
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kecamatanAll []*Kecamatan

	for rows.Next() {
		var kecamatan Kecamatan

		err := rows.Scan(
			&kecamatan.ID,
			&kecamatan.CreatedAt,
			&kecamatan.Name,
		)
		if err != nil {
			return nil, err
		}

		kecamatanAll = append(kecamatanAll, &kecamatan)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return kecamatanAll, nil
}
