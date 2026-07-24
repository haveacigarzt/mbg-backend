CREATE TABLE bumil (
    penduduk_id BIGINT PRIMARY KEY
        REFERENCES penduduk(id) ON DELETE CASCADE,

    hpht DATE,                          -- Hari Pertama Haid Terakhir
    hpl DATE,                           -- Hari Perkiraan Lahir

    gravida SMALLINT,                   -- Kehamilan ke-
    para SMALLINT,                      -- Jumlah persalinan
    abortus SMALLINT,                   -- Riwayat keguguran

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE busui (
    penduduk_id BIGINT PRIMARY KEY
        REFERENCES penduduk(id) ON DELETE CASCADE,

    tanggal_persalinan DATE,

    anak_ke SMALLINT,

    asi_eksklusif BOOLEAN,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE balita (
    penduduk_id BIGINT PRIMARY KEY
        REFERENCES penduduk(id) ON DELETE CASCADE,

    ibu_id BIGINT
        REFERENCES penduduk(id),

    anak_ke SMALLINT,

    berat_lahir NUMERIC(5,2),

    panjang_lahir NUMERIC(5,2),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);