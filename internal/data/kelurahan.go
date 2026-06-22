package data

import (
	"context"
	"database/sql"
	"time"
)

type Kelurahan struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	KecamatanID int32     `json:"kecamatan_id"`
}

type KelurahanModel struct {
	DB *sql.DB
}

func (m KelurahanModel) GetAll(id int64) ([]*Kelurahan, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT id, created_at, name, kecamatan_id
		FROM kelurahan
		WHERE kecamatan_id = $1
		ORDER BY name
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var kelurahanAll []*Kelurahan

	for rows.Next() {
		var kelurahan Kelurahan

		err := rows.Scan(
			&kelurahan.ID,
			&kelurahan.CreatedAt,
			&kelurahan.Name,
			&kelurahan.KecamatanID,
		)
		if err != nil {
			return nil, err
		}

		kelurahanAll = append(kelurahanAll, &kelurahan)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return kelurahanAll, nil
}
