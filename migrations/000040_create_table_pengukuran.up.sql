CREATE TABLE pengukuran_balita (
    id BIGSERIAL PRIMARY KEY,

    balita_id BIGINT NOT NULL
        REFERENCES balita(penduduk_id) ON DELETE CASCADE,

    tanggal DATE NOT NULL,

    umur_bulan SMALLINT NOT NULL,

    berat_badan NUMERIC(5,2) NOT NULL,
    tinggi_badan NUMERIC(5,2) NOT NULL,

    lingkar_kepala NUMERIC(5,2),
    lila NUMERIC(5,2),

    catatan TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE pengukuran_bumil (
    id BIGSERIAL PRIMARY KEY,

    bumil_id BIGINT NOT NULL
        REFERENCES bumil(penduduk_id) ON DELETE CASCADE,

    tanggal DATE NOT NULL,

    usia_kehamilan_minggu SMALLINT,

    berat_badan NUMERIC(5,2),

    tinggi_badan NUMERIC(5,2),

    lila NUMERIC(5,2),

    hemoglobin NUMERIC(4,2),

    tekanan_darah_sistolik SMALLINT,

    tekanan_darah_diastolik SMALLINT,

    catatan TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE pengukuran_busui (
    id BIGSERIAL PRIMARY KEY,

    busui_id BIGINT NOT NULL
        REFERENCES busui(penduduk_id) ON DELETE CASCADE,

    tanggal DATE NOT NULL,

    berat_badan NUMERIC(5,2),

    tinggi_badan NUMERIC(5,2),

    lila NUMERIC(5,2),

    hemoglobin NUMERIC(4,2),

    catatan TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE pengukuran_anak (
    id BIGSERIAL PRIMARY KEY,

    penduduk_id BIGINT NOT NULL
        REFERENCES penduduk(id) ON DELETE CASCADE,

    tanggal DATE NOT NULL,

    umur_bulan SMALLINT NOT NULL,

    berat_badan NUMERIC(5,2) NOT NULL,

    tinggi_badan NUMERIC(5,2) NOT NULL,

    catatan TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);