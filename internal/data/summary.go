package data

import (
	"database/sql"
)

type SummaryPenerimaManfaat struct {
	TotalPenerimaManfaat uint32 `json:"total_penerima_manfaat"`
	TotalPesertaDidik    uint32 `json:"total_peserta_didik"`
	TotalBumil           uint32 `json:"total_bumil"`
	TotalBusui           uint32 `json:"total_busui"`
	TotalBalita          uint32 `json:"total_balita"`
	TotalAPS             uint32 `json:"total_aps"`
	TotalATS             uint32 `json:"total_ats"`
}

type SummaryOperasional struct {
	TotalDapur         uint32 `json:"total_dapur"`
	TotalDapurAktif    uint32 `json:"total_dapur_aktif"`
	TotalDapurNonaktif uint32 `json:"total_dapur_nonaktif"`
}

type SebaranDapurPerKecamatan struct {
	KecamatanID int64  `json:"kecamatan_id"`
	Kecamatan   string `json:"kecamatan"`
	JumlahDapur uint32 `json:"jumlah_dapur"`
}

type SummaryDapur struct {
	Operasional SummaryOperasional         `json:"operasional"`
	Sebaran     []SebaranDapurPerKecamatan `json:"sebaran"`
}

type SummaryModel struct {
	DB *sql.DB
}

func (m SummaryModel) GetSummaryPenerimaManfaat() (*SummaryPenerimaManfaat, error) {
	query := `
		WITH sekolah_summary AS (
				SELECT COALESCE(SUM(jumlah_siswa), 0) AS peserta_didik
				FROM sekolah
				WHERE deleted_at IS NULL
		),
		posyandu_summary AS (
				SELECT
						COALESCE(SUM(jumlah_balita), 0) AS balita,
						COALESCE(SUM(jumlah_ibu_hamil), 0) AS bumil,
						COALESCE(SUM(jumlah_ibu_menyusui), 0) AS busui
				FROM posyandu
				WHERE deleted_at IS NULL
		)
		SELECT
				peserta_didik + balita + bumil + busui AS total_penerima_manfaat,
				peserta_didik AS total_peserta_didik,
				balita AS total_balita,
				bumil AS total_bumil,
				busui AS total_busui
		FROM sekolah_summary, posyandu_summary;
	`

	var summary SummaryPenerimaManfaat

	err := m.DB.QueryRow(query).Scan(
		&summary.TotalPenerimaManfaat,
		&summary.TotalPesertaDidik,
		&summary.TotalBalita,
		&summary.TotalBumil,
		&summary.TotalBusui,
	)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}

func (m SummaryModel) GetSummaryDapur() (*SummaryDapur, error) {
	var summary SummaryDapur

	// Summary operasional
	queryOperasional := `
		SELECT
			COUNT(*) AS total_dapur,
			COUNT(*) FILTER (WHERE status_aktif = TRUE) AS total_dapur_aktif,
			COUNT(*) FILTER (WHERE status_aktif = FALSE) AS total_dapur_nonaktif
		FROM sppg;
	`

	err := m.DB.QueryRow(queryOperasional).Scan(
		&summary.Operasional.TotalDapur,
		&summary.Operasional.TotalDapurAktif,
		&summary.Operasional.TotalDapurNonaktif,
	)
	if err != nil {
		return nil, err
	}

	// Sebaran per kecamatan
	querySebaran := `
		SELECT
			k.id,
			k.name,
			COUNT(s.id) AS jumlah_dapur
		FROM kecamatan k
		LEFT JOIN sppg s
			ON s.kecamatan_id = k.id
		GROUP BY
			k.id,
			k.name
		ORDER BY
			jumlah_dapur DESC,
			k.name;
	`

	rows, err := m.DB.Query(querySebaran)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item SebaranDapurPerKecamatan

		err := rows.Scan(
			&item.KecamatanID,
			&item.Kecamatan,
			&item.JumlahDapur,
		)
		if err != nil {
			return nil, err
		}

		summary.Sebaran = append(summary.Sebaran, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &summary, nil
}
