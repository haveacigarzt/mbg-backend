CREATE TABLE divisi_sppg (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),

    nama TEXT NOT NULL,
    urutan INTEGER NOT NULL,

    CONSTRAINT divisi_sppg_nama_unique UNIQUE (nama),
    CONSTRAINT divisi_sppg_urutan_unique UNIQUE (urutan),
    CONSTRAINT divisi_sppg_nama_length_check CHECK (char_length(nama) <= 100),
    CONSTRAINT divisi_sppg_urutan_check CHECK (urutan > 0)
);

INSERT INTO divisi_sppg (nama, urutan)
VALUES
    ('Persiapan', 1),
    ('Produksi (Masak)', 2),
    ('Pemorsian', 3),
    ('Distribusi & Kurir', 4),
    ('Pencucian Alat', 5),
    ('Kebersihan', 6),
    ('Keamanan (Security)', 7),
    ('Asisten Lapangan (ASLAP)', 8);