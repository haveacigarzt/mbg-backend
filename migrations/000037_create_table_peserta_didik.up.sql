CREATE TABLE peserta_didik (
    penduduk_id BIGINT PRIMARY KEY REFERENCES penduduk(id) ON DELETE CASCADE,

    sekolah_id BIGINT NOT NULL REFERENCES sekolah(id),

    nisn CHAR(10) UNIQUE,

    kelas VARCHAR(20),

    rombel VARCHAR(20),

    status_aktif BOOLEAN DEFAULT TRUE
);