CREATE TABLE IF NOT EXISTS alokasi_harian (
    id BIGSERIAL PRIMARY KEY,
    sppg_id BIGINT NOT NULL REFERENCES sppg(id),
    tanggal DATE NOT NULL,
    jumlah BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE (sppg_id, tanggal)
);