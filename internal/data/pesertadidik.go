package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mbg/internal/validator"
	"regexp"
	"strings"
	"time"

	"github.com/lib/pq"
)

type PesertaDidik struct {
	PendudukID int64  `json:"penduduk_id"`
	SekolahID  int64  `json:"sekolah_id"`
	Nama       string `json:"nama,omitempty"` // hasil JOIN dari tabel penduduk
	NISN       string `json:"nisn"`
	Kelas      string `json:"kelas"`
	Rombel     string `json:"rombel"`

	StatusAktif bool `json:"status_aktif"`
}

type PDResponse struct {
	NISN        string `json:"nisn"`
	Kelas       string `json:"kelas"`
	Rombel      string `json:"rombel"`
	SekolahID   int64  `json:"sekolah_id"`
	StatusAktif bool   `json:"status_aktif"`
	SekolahNama string `json:"sekolah_nama"`
}

type PesertaDidikResponse struct {
	Penduduk struct {
		ID            int64  `json:"id"`
		NIK           string `json:"nik"`
		Nama          string `json:"nama"`
		JenisKelamin  string `json:"jenis_kelamin"`
		TanggalLahir  string `json:"tanggal_lahir"`
		Umur          string `json:"umur"`
		KelurahanID   int64  `json:"kelurahan_id"`
		KelurahanNama string `json:"kelurahan_nama"`
		KecamatanID   int64  `json:"kecamatan_id"`
		KecamatanNama string `json:"kecamatan_nama"`
		Alamat        string `json:"alamat"`
		NoHP          string `json:"no_hp"`
	} `json:"penduduk"`
	PesertaDidik PDResponse `json:"peserta_didik"`
}

func ValidatePesertaDidik(v *validator.Validator, pd *PesertaDidik) {
	v.Check(pd.PendudukID > 0, "penduduk_id", "harus diisi")
	v.Check(pd.SekolahID > 0, "sekolah_id", "harus diisi")

	if pd.NISN != "" {
		v.Check(len(pd.NISN) == 10, "nisn", "harus terdiri dari 10 digit")
		v.Check(regexp.MustCompile(`^\d{10}$`).MatchString(pd.NISN), "nisn", "harus terdiri dari 10 digit angka")
	}

	if pd.Kelas != "" {
		v.Check(len(pd.Kelas) <= 20, "kelas", "maksimal 20 karakter")
	}

	if pd.Rombel != "" {
		v.Check(len(pd.Rombel) <= 20, "rombel", "maksimal 20 karakter")
	}
}

// Define a MovieModel struct type which wraps a sql.DB connection pool.
type PesertaDidikModel struct {
	DB *sql.DB
}

func (m PesertaDidikModel) InsertTx(ctx context.Context, tx *sql.Tx, pesertaDidik *PesertaDidik) error {
	query := `
INSERT INTO peserta_didik (
    penduduk_id,
    sekolah_id,
    nisn,
    kelas,
    rombel
)
VALUES ($1, $2, $3, $4, $5)
`

	_, err := tx.ExecContext(ctx, query,
		pesertaDidik.PendudukID,
		pesertaDidik.SekolahID,
		pesertaDidik.NISN,
		pesertaDidik.Kelas,
		pesertaDidik.Rombel,
	)

	if err != nil {
		var pqErr *pq.Error

		if errors.As(err, &pqErr) &&
			pqErr.Code == "23505" &&
			pqErr.Constraint == "peserta_didik_nisn_key" {
			return ErrDuplicateNISN
		}

		return err
	}
	return nil
}

func (m PesertaDidikModel) GetAll(sekolah_id int64, nama string, filters Filters) ([]*PesertaDidikResponse, Metadata, error) {

	query := fmt.Sprintf(`
		SELECT count(*) OVER(),
			p.id,
			p.nik,
			p.nama AS penduduk_nama,
			p.jenis_kelamin AS penduduk_jenis_kelamin,
			p.tanggal_lahir AS penduduk_tanggal_lahir,
			EXTRACT(YEAR FROM AGE(p.tanggal_lahir)) AS penduduk_umur,
			p.kelurahan_id,
			k.name,
			k.kecamatan_id,
			kc.name,
			p.alamat,
			p.no_hp,
		
			pd.nisn,
			pd.kelas AS peserta_didik_kelas,
			pd.rombel AS peserta_didik_rombel,
			pd.status_aktif,
			pd.sekolah_id,
			s.nama
	FROM peserta_didik pd
	JOIN penduduk p ON p.id = pd.penduduk_id
	JOIN sekolah s ON s.id = pd.sekolah_id
	JOIN kelurahan k ON k.id = p.kelurahan_id
	JOIN kecamatan kc ON kc.id = k.kecamatan_id
	WHERE (LOWER(p.nama) LIKE LOWER('%%' || $1 || '%%') OR $1 = '')
  AND pd.sekolah_id = $2
	AND pd.status_aktif = TRUE
  AND p.deleted_at IS NULL
	ORDER BY %s %s, p.id ASC
	LIMIT $3 OFFSET $4`,
		filters.sortColumn(),
		filters.sortDirection(),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{nama, sekolah_id, filters.limit(), filters.offset()}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, Metadata{}, err
	}

	defer rows.Close()

	totalRecords := 0
	peserta_didik_all := []*PesertaDidikResponse{}

	for rows.Next() {

		var peserta_didik PesertaDidikResponse

		err := rows.Scan(
			&totalRecords,
			&peserta_didik.Penduduk.ID,
			&peserta_didik.Penduduk.NIK,
			&peserta_didik.Penduduk.Nama,
			&peserta_didik.Penduduk.JenisKelamin,
			&peserta_didik.Penduduk.TanggalLahir,
			&peserta_didik.Penduduk.Umur,
			&peserta_didik.Penduduk.KelurahanID,
			&peserta_didik.Penduduk.KelurahanNama,
			&peserta_didik.Penduduk.KecamatanID,
			&peserta_didik.Penduduk.KecamatanNama,
			&peserta_didik.Penduduk.Alamat,
			&peserta_didik.Penduduk.NoHP,
			&peserta_didik.PesertaDidik.NISN,
			&peserta_didik.PesertaDidik.Kelas,
			&peserta_didik.PesertaDidik.Rombel,
			&peserta_didik.PesertaDidik.StatusAktif,
			&peserta_didik.PesertaDidik.SekolahID,
			&peserta_didik.PesertaDidik.SekolahNama,
		)
		if err != nil {
			return nil, Metadata{}, err
		}

		peserta_didik_all = append(peserta_didik_all, &peserta_didik)
	}

	if err = rows.Err(); err != nil {
		return nil, Metadata{}, err
	}

	metadata := calculateMetadata(totalRecords, filters.Page, filters.PageSize)

	return peserta_didik_all, metadata, nil
}

func (m PesertaDidikModel) GetByNISN(nisn string) (*PesertaDidikResponse, error) {
	query := `
		SELECT
			p.id,
			p.nik,
			p.nama AS penduduk_nama,
			p.jenis_kelamin AS penduduk_jenis_kelamin,
			p.tanggal_lahir AS penduduk_tanggal_lahir,
			EXTRACT(YEAR FROM AGE(p.tanggal_lahir)) AS penduduk_umur,
			p.kelurahan_id,
			k.name,
			kc.id,
			kc.name,
			p.alamat,
			p.no_hp,

			pd.nisn,
			pd.kelas,
			pd.rombel,
			pd.status_aktif,
			pd.sekolah_id,
			s.nama AS sekolah_nama

		FROM peserta_didik pd
		JOIN penduduk p ON p.id = pd.penduduk_id
		JOIN sekolah s ON s.id = pd.sekolah_id
		JOIN kelurahan k ON k.id = p.kelurahan_id
		JOIN kecamatan kc ON kc.id = k.kecamatan_id

		WHERE pd.nisn = $1
	`
	var peserta_didik PesertaDidikResponse
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, nisn).Scan(
		// Peserta Didik
		&peserta_didik.Penduduk.ID,
		&peserta_didik.Penduduk.NIK,
		&peserta_didik.Penduduk.Nama,
		&peserta_didik.Penduduk.JenisKelamin,
		&peserta_didik.Penduduk.TanggalLahir,
		&peserta_didik.Penduduk.Umur,
		&peserta_didik.Penduduk.KelurahanID,
		&peserta_didik.Penduduk.KelurahanNama,
		&peserta_didik.Penduduk.KecamatanID,
		&peserta_didik.Penduduk.KecamatanNama,
		&peserta_didik.Penduduk.Alamat,
		&peserta_didik.Penduduk.NoHP,
		&peserta_didik.PesertaDidik.NISN,
		&peserta_didik.PesertaDidik.Kelas,
		&peserta_didik.PesertaDidik.Rombel,
		&peserta_didik.PesertaDidik.StatusAktif,
		&peserta_didik.PesertaDidik.SekolahID,
		&peserta_didik.PesertaDidik.SekolahNama,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &peserta_didik, nil

}

func (m PesertaDidikModel) GetByID(id int64) (*PesertaDidikResponse, error) {
	query := `
		SELECT
			p.id,
			p.nik,
			p.nama AS penduduk_nama,
			p.jenis_kelamin AS penduduk_jenis_kelamin,
			p.tanggal_lahir AS penduduk_tanggal_lahir,
			EXTRACT(YEAR FROM AGE(p.tanggal_lahir)) AS penduduk_umur,
			p.kelurahan_id,
			k.name,
			kc.id,
			kc.name,
			p.alamat,
			p.no_hp,

			pd.nisn,
			pd.kelas,
			pd.rombel,
			pd.status_aktif,
			pd.sekolah_id,
			s.nama AS sekolah_nama

		FROM penduduk p
		JOIN peserta_didik pd ON pd.penduduk_id = p.id
		JOIN sekolah s ON s.id = pd.sekolah_id
		JOIN kelurahan k ON k.id = p.kelurahan_id
		JOIN kecamatan kc ON kc.id = k.kecamatan_id

		WHERE p.id = $1
	`
	var peserta_didik PesertaDidikResponse
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		// Peserta Didik
		&peserta_didik.Penduduk.ID,
		&peserta_didik.Penduduk.NIK,
		&peserta_didik.Penduduk.Nama,
		&peserta_didik.Penduduk.JenisKelamin,
		&peserta_didik.Penduduk.TanggalLahir,
		&peserta_didik.Penduduk.Umur,
		&peserta_didik.Penduduk.KelurahanID,
		&peserta_didik.Penduduk.KelurahanNama,
		&peserta_didik.Penduduk.KecamatanID,
		&peserta_didik.Penduduk.KecamatanNama,
		&peserta_didik.Penduduk.Alamat,
		&peserta_didik.Penduduk.NoHP,
		&peserta_didik.PesertaDidik.NISN,
		&peserta_didik.PesertaDidik.Kelas,
		&peserta_didik.PesertaDidik.Rombel,
		&peserta_didik.PesertaDidik.StatusAktif,
		&peserta_didik.PesertaDidik.SekolahID,
		&peserta_didik.PesertaDidik.SekolahNama,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &peserta_didik, nil

}

func (m PesertaDidikModel) Update(pesertaDidik *PesertaDidikResponse) error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queryPenduduk := `
		UPDATE penduduk
		SET
			nik = $1,
			nama = $2,
			jenis_kelamin = $3,
			tanggal_lahir = $4,
			kelurahan_id = $5,
			alamat = $6,
			no_hp = $7
		WHERE id = $8
	`

	argsPenduduk := []any{
		pesertaDidik.Penduduk.NIK,
		pesertaDidik.Penduduk.Nama,
		pesertaDidik.Penduduk.JenisKelamin,
		pesertaDidik.Penduduk.TanggalLahir,
		pesertaDidik.Penduduk.KelurahanID,
		pesertaDidik.Penduduk.Alamat,
		pesertaDidik.Penduduk.NoHP,
		pesertaDidik.Penduduk.ID,
	}

	if _, err := tx.Exec(queryPenduduk, argsPenduduk...); err != nil {
		return err
	}

	queryPesertaDidik := `
		UPDATE peserta_didik
		SET
			nisn = $1,
			kelas = $2,
			rombel = $3,
			status_aktif = $4
		WHERE penduduk_id = $5
	`

	argsPesertaDidik := []any{
		pesertaDidik.PesertaDidik.NISN,
		pesertaDidik.PesertaDidik.Kelas,
		pesertaDidik.PesertaDidik.Rombel,
		pesertaDidik.PesertaDidik.StatusAktif,
		pesertaDidik.Penduduk.ID,
	}

	if _, err := tx.Exec(queryPesertaDidik, argsPesertaDidik...); err != nil {
		return err
	}

	return tx.Commit()
}

func (m PesertaDidikModel) UpdatePDOnly(pesertaDidik *PesertaDidik) error {
	query := `
		UPDATE peserta_didik
		SET
			kelas = $1,
			rombel = $2,
			sekolah_id = $3,
			status_aktif = true
		WHERE penduduk_id = $4
		AND status_aktif = false
	`

	args := []any{
		pesertaDidik.Kelas,
		pesertaDidik.Rombel,
		pesertaDidik.SekolahID,
		pesertaDidik.PendudukID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrEditConflict
	}

	return nil
}

func (m PesertaDidikModel) UpdateFrom3B(pesertaDidik *PesertaDidik, jenis KategoriSasaran) error {
	tx, err := m.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	queryPenduduk := `
		UPDATE penduduk
		SET
			kategori = 'PESERTA_DIDIK'
		WHERE id = $1
	`

	argsPenduduk := []any{
		pesertaDidik.PendudukID,
	}

	if _, err := tx.Exec(queryPenduduk, argsPenduduk...); err != nil {
		return err
	}

	table3B := strings.ToLower(string(jenis))

	query3B := fmt.Sprintf(`
    DELETE FROM %s
    WHERE penduduk_id = $1
`, table3B)

	if _, err := tx.Exec(query3B, pesertaDidik.PendudukID); err != nil {
		return err
	}

	queryPesertaDidik := `
	INSERT INTO peserta_didik (
			penduduk_id,
			sekolah_id,
			nisn,
			kelas,
			rombel
	)
	VALUES ($1, $2, $3, $4, $5)
	`

	argsPesertaDidik := []any{
		pesertaDidik.PendudukID,
		pesertaDidik.SekolahID,
		pesertaDidik.NISN,
		pesertaDidik.Kelas,
		pesertaDidik.Rombel,
	}

	if _, err := tx.Exec(queryPesertaDidik, argsPesertaDidik...); err != nil {
		return err
	}

	return tx.Commit()
}

func (m PesertaDidikModel) ValidatePesertaDidikInclude(sekolahID, pendudukID int64) error {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM peserta_didik
			WHERE sekolah_id = $1
			  AND penduduk_id = $2
		)
	`

	var exists bool

	err := m.DB.QueryRow(query, sekolahID, pendudukID).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("peserta didik tidak terdaftar di sekolah ini")
	}

	return nil
}

func (m PesertaDidikModel) ValidatePesertaDidikExclude(sekolahID, pendudukID int64) error {
	query := `
		SELECT EXISTS (
			SELECT 1
			FROM peserta_didik pd
			WHERE pd.penduduk_id = $1
			AND pd.sekolah_id <> $2
		)
	`

	var exists bool

	err := m.DB.QueryRow(query, pendudukID, sekolahID).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("peserta didik sudah terdaftar di sekolah lain")
	}

	return nil
}
