package data

import (
	"context"
	"database/sql"
	"fmt"
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
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Version     int32     `json:"version"`
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
