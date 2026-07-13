CREATE TABLE IF NOT EXISTS sekolah (
    id bigserial PRIMARY KEY,
    created_at timestamptz NOT NULL DEFAULT NOW(),
    nama text NOT NULL,
    alamat text NOT NULL,
    tingkat text NOT NULL,
    jumlah_siswa integer NOT NULL,
    kecamatan text NOT NULL,
    latitude double precision NOT NULL,
    longitude double precision NOT NULL,
    sosmed_url text
);