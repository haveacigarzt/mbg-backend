CREATE TABLE pedagang_lokal (
    id BIGSERIAL PRIMARY KEY,

    nama TEXT NOT NULL,
    alamat TEXT NOT NULL,
    no_hp TEXT NOT NULL,

    longitude DOUBLE PRECISION NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,

    jenis_produk TEXT NOT NULL,

    sppg_id BIGINT NOT NULL
        REFERENCES sppg(id)
        ON DELETE CASCADE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    version INTEGER NOT NULL DEFAULT 1
);

CREATE INDEX idx_pedagang_lokal_sppg
ON pedagang_lokal(sppg_id);

CREATE INDEX idx_pedagang_lokal_jenis_produk
ON pedagang_lokal(jenis_produk);