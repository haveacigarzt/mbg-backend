package data

import (
	"database/sql"
)

type SummaryPenerimaManfaat struct {
	TotalPenerimaManfaat uint32 `json:"total_penerima_manfaat"`
	TKPAUD               uint32 `json:"tk_paud"`
	SDMI                 uint32 `json:"sd_mi"`
	SMPMTS               uint32 `json:"smp_mts"`
	SMASMKMA             uint32 `json:"sma_smk_ma"`
	Yayasan              uint32 `json:"yayasan"`
	Balita               uint32 `json:"balita"`
	Bumil                uint32 `json:"bumil"`
	Busui                uint32 `json:"busui"`
}

type SummaryOperasional struct {
	TotalDapur    uint32 `json:"total_dapur"`
	DapurAktif    uint32 `json:"dapur_aktif"`
	DapurNonaktif uint32 `json:"dapur_nonaktif"`
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

type SummaryLapanganPekerjaan struct {
	Divisi    string `json:"divisi"`
	JumlahSDM uint32 `json:"jumlah_sdm"`
}

type SummaryModel struct {
	DB *sql.DB
}

func (m SummaryModel) GetSummaryPenerimaManfaat() (*SummaryPenerimaManfaat, error) {
	query := `
		WITH sekolah_summary AS (
				SELECT
						COALESCE(SUM(jumlah_siswa), 0) AS peserta_didik,

						COALESCE(SUM(jumlah_siswa)
								FILTER (WHERE kategori = 'PAUD/TK'), 0) AS tk_paud,

						COALESCE(SUM(jumlah_siswa)
								FILTER (WHERE kategori = 'SD/MI'), 0) AS sd_mi,

						COALESCE(SUM(jumlah_siswa)
								FILTER (WHERE kategori = 'SMP/MTs'), 0) AS smp_mts,

						COALESCE(SUM(jumlah_siswa)
								FILTER (WHERE kategori = 'SMA/SMK/MA'), 0) AS sma_smk_ma,

						COALESCE(SUM(jumlah_siswa)
								FILTER (WHERE kategori = 'YAYASAN'), 0) AS yayasan

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

				tk_paud,
				sd_mi,
				smp_mts,
				sma_smk_ma,
				yayasan,

				balita,
				bumil,
				busui

		FROM sekolah_summary, posyandu_summary;
		`

	var summary SummaryPenerimaManfaat

	err := m.DB.QueryRow(query).Scan(
		&summary.TotalPenerimaManfaat,
		&summary.TKPAUD,
		&summary.SDMI,
		&summary.SMPMTS,
		&summary.SMASMKMA,
		&summary.Yayasan,
		&summary.Balita,
		&summary.Bumil,
		&summary.Busui,
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
			COUNT(*) FILTER (WHERE status_aktif = TRUE) AS dapur_aktif,
			COUNT(*) FILTER (WHERE status_aktif = FALSE) AS dapur_nonaktif
		FROM sppg;
	`

	err := m.DB.QueryRow(queryOperasional).Scan(
		&summary.Operasional.TotalDapur,
		&summary.Operasional.DapurAktif,
		&summary.Operasional.DapurNonaktif,
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

func (m SummaryModel) GetLapanganPekerjaan() (*[]SummaryLapanganPekerjaan, error) {
	query := `
	SELECT
    d.nama AS divisi,
    COALESCE(SUM(sd.jumlah_sdm), 0) AS jumlah_sdm
FROM divisi_sppg d
LEFT JOIN sppg_divisi sd
    ON sd.divisi_id = d.id
LEFT JOIN sppg s
    ON s.id = sd.sppg_id
    AND s.status_aktif = TRUE
GROUP BY d.id, d.nama, d.urutan
ORDER BY d.urutan;`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	summary := []SummaryLapanganPekerjaan{}

	for rows.Next() {
		var item SummaryLapanganPekerjaan

		err := rows.Scan(
			&item.Divisi,
			&item.JumlahSDM,
		)
		if err != nil {
			return nil, err
		}

		summary = append(summary, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &summary, nil
}
