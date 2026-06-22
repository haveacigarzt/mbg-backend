CREATE TABLE IF NOT EXISTS pengeluaran_harian (
    id BIGSERIAL PRIMARY KEY,
    alokasi_harian_id BIGINT NOT NULL REFERENCES alokasi_harian(id) ON DELETE CASCADE,

    produk VARCHAR(255) NOT NULL,
    jumlah NUMERIC(12,2) NOT NULL,
    satuan VARCHAR(50) NOT NULL,
    harga_satuan BIGINT NOT NULL,

    subtotal BIGINT GENERATED ALWAYS AS (
        (jumlah * harga_satuan)::BIGINT
    ) STORED,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);