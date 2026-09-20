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
	Penduduk struct {
		ID            int64           `json:"id"`
		NIK           int64           `json:"nik"`
		Nama          string          `json:"nama"`
		JenisKelamin  string          `json:"jenis_kelamin"`
		TanggalLahir  string          `json:"tanggal_lahir"`
		KelurahanID   string          `json:"kelurahan_id"`
		KelurahanNama string          `json:"kelurahan_nama"`
		KecamatanID   int64           `json:"kecamatan_id"`
		Alamat        string          `json:"alamat"`
		NoHP          string          `json:"no_hp"`
		Kategori      KategoriSasaran `json:"kategori"`
	} `json:"penduduk"`
	PesertaDidik *PDResponse `json:"peserta_didik,omitempty"`
	Balita       *BAResponse `json:"balita,omitempty"`
	Bumil        *BMResponse `json:"bumil,omitempty"`
	Busui        *BSResponse `json:"busui,omitempty"`
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

func (m PendudukModel) Get(nik string) (*PendudukResponse, error) {

	query := `
	SELECT 
		p.id,
		p.nik,
		p.nama,
		p.jenis_kelamin,
		p.tanggal_lahir,
		p.kelurahan_id,
		k.name AS kelurahan_nama,
		k.kecamatan_id,
		p.alamat,
		p.no_hp,
		p.kategori,

		-- Peserta Didik
		pd.nisn,
		pd.kelas,
		pd.rombel,
		pd.sekolah_id,
		pd.status_aktif,
		s.nama AS sekolah_nama,

		-- Balita
		ba.ibu_id,
		ibu.nama AS ibu_nama,
		ba.anak_ke,
		ba.berat_lahir,
		ba.panjang_lahir,
		ba.posyandu_id,
		ba.status_aktif,
		pos_ba.nama AS posyandu_nama,

		-- Bumil
		bm.hpht,
		bm.hpl,
		bm.gravida,
		bm.para,
		bm.abortus,
		bm.posyandu_id,
		bm.status_aktif,
		pos_bm.nama AS posyandu_nama,

		-- Busui
		bs.tanggal_persalinan,
		bs.anak_ke,
		bs.asi_eksklusif,
		bs.posyandu_id,
		bs.status_aktif,
		pos_bs.nama AS posyandu_nama

	FROM penduduk p

	JOIN kelurahan k
		ON k.id = p.kelurahan_id

	LEFT JOIN peserta_didik pd
		ON pd.penduduk_id = p.id
		AND p.kategori = 'PESERTA_DIDIK'

	LEFT JOIN sekolah s
		ON s.id = pd.sekolah_id

	LEFT JOIN balita ba
		ON ba.penduduk_id = p.id
		AND p.kategori = 'BALITA'

	LEFT JOIN bumil bm
		ON bm.penduduk_id = p.id
		AND p.kategori = 'BUMIL'

	LEFT JOIN busui bs
		ON bs.penduduk_id = p.id
		AND p.kategori = 'BUSUI'

	LEFT JOIN penduduk ibu
		ON ibu.id = ba.ibu_id

	LEFT JOIN posyandu pos_ba
    ON pos_ba.id = ba.posyandu_id

	LEFT JOIN posyandu pos_bm
			ON pos_bm.id = bm.posyandu_id

	LEFT JOIN posyandu pos_bs
			ON pos_bs.id = bs.posyandu_id

	WHERE
		p.nik = $1
		AND p.deleted_at IS NULL
`

	var response PendudukResponse

	var (
		pdNISN        sql.NullString
		pdKelas       sql.NullString
		pdRombel      sql.NullString
		pdSekolahID   sql.NullInt64
		pdStatusAktif sql.NullBool
		pdSekolahNama sql.NullString

		baIbuID        sql.NullInt64
		baIbuNama      sql.NullString
		baAnakKe       sql.NullInt64
		baBeratLahir   sql.NullInt64
		baPanjangLahir sql.NullInt64
		baPosyanduID   sql.NullInt64
		baStatusAktif  sql.NullBool
		baPosyanduNama sql.NullString

		bmHPHT         sql.NullTime
		bmHPL          sql.NullTime
		bmGravida      sql.NullInt64
		bmPara         sql.NullInt64
		bmAbortus      sql.NullInt64
		bmPosyanduID   sql.NullInt64
		bmStatusAktif  sql.NullBool
		bmPosyanduNama sql.NullString

		bsTanggalPersalinan sql.NullTime
		bsAnakKe            sql.NullInt64
		bsAsiEksklusif      sql.NullBool
		bsPosyanduID        sql.NullInt64
		bsStatusAktif       sql.NullBool
		bsPosyanduNama      sql.NullString
	)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, nik).Scan(
		&response.Penduduk.ID,
		&response.Penduduk.NIK,
		&response.Penduduk.Nama,
		&response.Penduduk.JenisKelamin,
		&response.Penduduk.TanggalLahir,
		&response.Penduduk.KelurahanID,
		&response.Penduduk.KelurahanNama,
		&response.Penduduk.KecamatanID,
		&response.Penduduk.Alamat,
		&response.Penduduk.NoHP,
		&response.Penduduk.Kategori,

		// Peserta Didik
		&pdNISN,
		&pdKelas,
		&pdRombel,
		&pdSekolahID,
		&pdStatusAktif,
		&pdSekolahNama,

		// Balita
		&baIbuID,
		&baIbuNama,
		&baAnakKe,
		&baBeratLahir,
		&baPanjangLahir,
		&baPosyanduID,
		&baStatusAktif,
		&baPosyanduNama,

		// Bumil
		&bmHPHT,
		&bmHPL,
		&bmGravida,
		&bmPara,
		&bmAbortus,
		&bmPosyanduID,
		&bmStatusAktif,
		&bmPosyanduNama,

		// Busui
		&bsTanggalPersalinan,
		&bsAnakKe,
		&bsAsiEksklusif,
		&bsPosyanduID,
		&bsStatusAktif,
		&bsPosyanduNama,
	)

	switch response.Penduduk.Kategori {
	case "PESERTA_DIDIK":
		response.PesertaDidik = &PDResponse{
			NISN:        pdNISN.String,
			Kelas:       pdKelas.String,
			Rombel:      pdRombel.String,
			SekolahID:   pdSekolahID.Int64,
			StatusAktif: pdStatusAktif.Bool,
			SekolahNama: pdSekolahNama.String,
		}

	case "BALITA":
		response.Balita = &BAResponse{
			IbuID:        baIbuID.Int64,
			IbuNama:      baIbuNama.String,
			AnakKe:       int8(baAnakKe.Int64),
			BeratLahir:   int(baBeratLahir.Int64),
			PanjangLahir: int(baPanjangLahir.Int64),
			PosyanduID:   baPosyanduID.Int64,
			StatusAktif:  baStatusAktif.Bool,
			PosyanduNama: baPosyanduNama.String,
		}

	case "BUMIL":
		response.Bumil = &BMResponse{
			HPHT:         bmHPHT.Time,
			HPL:          bmHPL.Time,
			Gravida:      int(bmGravida.Int64),
			Para:         int(bmPara.Int64),
			Abortus:      int(bmAbortus.Int64),
			PosyanduID:   bmPosyanduID.Int64,
			StatusAktif:  bmStatusAktif.Bool,
			PosyanduNama: bmPosyanduNama.String,
		}

	case "BUSUI":
		response.Busui = &BSResponse{
			TanggalPersalinan: bsTanggalPersalinan.Time,
			AnakKe:            int8(bsAnakKe.Int64),
			AsiEksklusif:      bsAsiEksklusif.Bool,
			PosyanduID:        bsPosyanduID.Int64,
			StatusAktif:       bsStatusAktif.Bool,
			PosyanduNama:      bsPosyanduNama.String,
		}
	}

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &response, nil
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
			k.kecamatan_id,
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
			&penduduk.Penduduk.ID,
			&penduduk.Penduduk.NIK,
			&penduduk.Penduduk.Nama,
			&penduduk.Penduduk.JenisKelamin,
			&penduduk.Penduduk.TanggalLahir,
			&penduduk.Penduduk.KelurahanID,
			&penduduk.Penduduk.KelurahanNama,
			&penduduk.Penduduk.KecamatanID,
			&penduduk.Penduduk.Alamat,
			&penduduk.Penduduk.NoHP,
			&penduduk.Penduduk.Kategori,
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
