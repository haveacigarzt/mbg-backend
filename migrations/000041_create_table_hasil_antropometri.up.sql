CREATE TYPE indikator_antropometri AS ENUM (
    'TB_U',
    'BB_U',
    'BB_TB',
    'IMT_U',
    'LK_U'      -- Lingkar Kepala menurut Umur (opsional)
);

CREATE TABLE hasil_antropometri (
    id BIGSERIAL PRIMARY KEY,

    jenis_pengukuran VARCHAR(20) NOT NULL
        CHECK (jenis_pengukuran IN ('BALITA', 'ANAK')),

    pengukuran_id BIGINT NOT NULL,

    indikator indikator_antropometri NOT NULL,

    zscore NUMERIC(5,2),

    status VARCHAR(30),

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);