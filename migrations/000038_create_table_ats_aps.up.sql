CREATE TYPE pendidikan_terakhir AS ENUM (
    'BELUM_SEKOLAH',
    'TK',
    'SD',
    'SMP',
    'SMA',
    'SMK',
    'PAKET_A',
    'PAKET_B',
    'PAKET_C'
);

CREATE TABLE ats (
    penduduk_id BIGINT PRIMARY KEY REFERENCES penduduk(id) ON DELETE CASCADE,

    pendidikan_terakhir pendidikan_terakhir NOT NULL,

    alasan_tidak_sekolah TEXT
);

CREATE TABLE aps (
    penduduk_id BIGINT PRIMARY KEY REFERENCES penduduk(id) ON DELETE CASCADE,

    sekolah_id BIGINT REFERENCES sekolah(id),

    kelas_terakhir VARCHAR(30),

    tanggal_putus DATE,

    alasan_putus TEXT
);