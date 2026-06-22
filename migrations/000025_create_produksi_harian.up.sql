CREATE TABLE IF NOT EXISTS produksi_harian (
    id BIGSERIAL PRIMARY KEY,
    sppg_id BIGINT NOT NULL REFERENCES sppg(id),
    tanggal DATE NOT NULL,
    waktu_mulai TIMESTAMP NOT NULL,
    estimasi_waktu_selesai TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE (sppg_id, tanggal)
);