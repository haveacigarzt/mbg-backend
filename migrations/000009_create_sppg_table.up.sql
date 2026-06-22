CREATE TABLE IF NOT EXISTS sppg (
    id bigserial PRIMARY KEY,
    user_id bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,

    created_at timestamptz(0) NOT NULL DEFAULT NOW(),
    updated_at timestamptz(0) NOT NULL DEFAULT NOW(),

    nama text NOT NULL,
    alamat text NOT NULL,
    sosmed_url text[],

    kepala_sppg text NOT NULL,
    nomor_telepon text,
    email text,

    latitude double precision,
    longitude double precision,

    kecamatan text NOT NULL,
    kelurahan text NOT NULL,

    kapasitas_porsi integer NOT NULL DEFAULT 0,
    status_aktif boolean NOT NULL DEFAULT true,

    version integer NOT NULL DEFAULT 1
);