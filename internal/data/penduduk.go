package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"regexp"
	"time"

	"github.com/lib/pq"
)

type KategoriSasaran string

const (
	KategoriBUMIL        KategoriSasaran = "BUMIL"
	KategoriBALITA       KategoriSasaran = "BALITA"
	KategoriBUSUI        KategoriSasaran = "BUSUI"
	KategoriPesertaDidik KategoriSasaran = "PESERTA_DIDIK"
	KategoriATS          KategoriSasaran = "ATS"
	KategoriAPS          KategoriSasaran = "APS"
)

type Penduduk struct {
	ID        int64      `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`

	NIK          string          `json:"nik"`
	Nama         string          `json:"nama"`
	JenisKelamin string          `json:"jenis_kelamin"`
	TanggalLahir time.Time       `json:"tanggal_lahir"`
	KelurahanID  int64           `json:"kelurahan_id"`
	Alamat       string          `json:"alamat"`
	NoHP         string          `json:"no_hp"`
	Kategori     KategoriSasaran `json:"kategori"`
}

type PendudukResponse struct {
	ID            int64           `json:"id"`
	NIK           string          `json:"nik"`
	Nama          string          `json:"nama"`
	JenisKelamin  string          `json:"jenis_kelamin"`
	TanggalLahir  time.Time       `json:"tanggal_lahir"`
	KelurahanID   int64           `json:"kelurahan_id"`
	KelurahanNama string          `json:"kelurahan_nama"`
	Alamat        string          `json:"alamat"`
	NoHP          string          `json:"no_hp"`
	Kategori      KategoriSasaran `json:"kategori"`
}

type PendudukInput struct {
	NIK          string          `json:"nik"`
	Nama         string          `json:"nama"`
	JenisKelamin string          `json:"jenis_kelamin"`
	TanggalLahir string          `json:"tanggal_lahir"`
	KelurahanID  int64           `json:"kelurahan_id"`
	Alamat       string          `json:"alamat"`
	NoHP         string          `json:"no_hp"`
	Kategori     KategoriSasaran `json:"kategori"`
}

func ValidatePendudukInput(v *validator.Validator, input *PendudukInput) {
	if input.NIK != "" {
		v.Check(len(input.NIK) == 16, "nik", "harus terdiri dari 16 digit")
		v.Check(regexp.MustCompile(`^\d{16}$`).MatchString(input.NIK), "nik", "hanya boleh berisi angka")
	}

	v.Check(input.Nama != "", "nama", "harus diisi")
	v.Check(len(input.Nama) <= 255, "nama", "maksimal 255 karakter")

	v.Check(
		input.JenisKelamin == "L" || input.JenisKelamin == "P",
		"jenis_kelamin",
		"harus bernilai L atau P",
	)

	v.Check(input.TanggalLahir != "", "tanggal_lahir", "harus diisi")

	v.Check(input.KelurahanID > 0, "kelurahan_id", "harus diisi")

	if input.NoHP != "" {
		v.Check(len(input.NoHP) <= 20, "no_hp", "maksimal 20 karakter")
	}

	v.Check(
		input.Kategori == KategoriBUMIL ||
			input.Kategori == KategoriBALITA ||
			input.Kategori == KategoriBUSUI ||
			input.Kategori == KategoriPesertaDidik ||
			input.Kategori == KategoriATS ||
			input.Kategori == KategoriAPS,
		"kategori",
		"tidak valid",
	)
}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type PendudukModel struct {
	DB *sql.DB
}

func (m PendudukModel) InsertTx(ctx context.Context, tx *sql.Tx, penduduk *Penduduk) error {
	// Define the SQL query for inserting a new record in the movies table and returning
	// the system-generated data.
	query := `
INSERT INTO penduduk (
		nik,
    nama,
    jenis_kelamin,
    tanggal_lahir,
    kelurahan_id,
    alamat,
    no_hp,
		kategori
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, created_at`

	args := []any{
		penduduk.NIK,
		penduduk.Nama,
		penduduk.JenisKelamin,
		penduduk.TanggalLahir,
		penduduk.KelurahanID,
		penduduk.Alamat,
		penduduk.NoHP,
		penduduk.Kategori,
	}

	err := tx.QueryRowContext(ctx, query, args...).Scan(
		&penduduk.ID,
		&penduduk.CreatedAt,
	)
	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) &&
			pqErr.Code == "23505" &&
			pqErr.Constraint == "penduduk_nik_key" {
			return ErrDuplicateNIK
		}

		return err
	}
	return nil
}

func (m PendudukModel) GetAllIbu(posyandu_id int64, nama string, filters Filters) ([]*PendudukResponse, Metadata, error) {

	query := fmt.Sprintf(`
		SELECT count(*) OVER(),
			p.id,
			p.nik,
			p.nama,
			p.jenis_kelamin,
			p.tanggal_lahir,
			p.kelurahan_id,
			k.name,
			p.alamat,
			p.no_hp,
			p.kategori
		FROM penduduk p
		LEFT JOIN bumil b
			ON b.penduduk_id = p.id
			AND p.kategori = 'BUMIL'
		LEFT JOIN busui bs
			ON bs.penduduk_id = p.id
			AND p.kategori = 'BUSUI'
		JOIN kelurahan k
			ON k.id = p.kelurahan_id
		WHERE
			(LOWER(p.nama) LIKE LOWER('%%' || $1 || '%%') OR $1 = '')
			AND (
				b.posyandu_id = $2 OR
				bs.posyandu_id = $2
			)
			AND p.deleted_at IS NULL
		ORDER BY %s %s, p.id ASC
		LIMIT $3 OFFSET $4`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{nama, posyandu_id, filters.limit(), filters.offset()}
	// args := []any{nama, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	penduduk_all := []*PendudukResponse{}

	for rows.Next() {

		var penduduk PendudukResponse

		err := rows.Scan(
			&totalRecords,
			&penduduk.ID,
			&penduduk.NIK,
			&penduduk.Nama,
			&penduduk.JenisKelamin,
			&penduduk.TanggalLahir,
			&penduduk.KelurahanID,
			&penduduk.KelurahanNama,
			&penduduk.Alamat,
			&penduduk.NoHP,
			&penduduk.Kategori,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		penduduk_all = append(penduduk_all, &penduduk)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return penduduk_all, metadata, nil
}
