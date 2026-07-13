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

    kecamatan_id BIGINT NOT NULL,
    kelurahan_id BIGINT NOT NULL,

    kapasitas_porsi integer NOT NULL DEFAULT 0,
    status_aktif boolean NOT NULL DEFAULT true,

    version integer NOT NULL DEFAULT 1,

    CONSTRAINT fk_sppg_kecamatan
        FOREIGN KEY (kecamatan_id)
        REFERENCES kecamatan(id),

    CONSTRAINT fk_sppg_kelurahan
        FOREIGN KEY (kelurahan_id)
        REFERENCES kelurahan(id)
);
