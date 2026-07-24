CREATE TYPE kategori_sasaran AS ENUM (
    'BUMIL',
    'BALITA',
    'BUSUI',
    'PESERTA_DIDIK',
    'ATS',
    'APS'
);

CREATE TABLE penduduk (
    id BIGSERIAL PRIMARY KEY,

    nik CHAR(16) UNIQUE,

    nama VARCHAR(255) NOT NULL,

    jenis_kelamin CHAR(1)
        CHECK (jenis_kelamin IN ('L','P')),

    tanggal_lahir DATE,

    kelurahan_id BIGINT REFERENCES kelurahan(id),

    alamat TEXT,

    no_hp VARCHAR(20),

    kategori kategori_sasaran NOT NULL,

    created_at TIMESTAMP DEFAULT now(),

    updated_at TIMESTAMP DEFAULT now(),

    deleted_at TIMESTAMP
);