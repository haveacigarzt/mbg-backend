CREATE TABLE IF NOT EXISTS posyandu (
    id bigserial PRIMARY KEY,

    created_at timestamptz(0) NOT NULL DEFAULT NOW(),
    updated_at timestamptz(0) NOT NULL DEFAULT NOW(),

    sppg_id bigint NOT NULL
        REFERENCES sppg(id)
        ON DELETE RESTRICT,

    nama text NOT NULL,
    alamat text NOT NULL,

    kecamatan text NOT NULL,
    kelurahan text NOT NULL,

    latitude double precision,
    longitude double precision,

    jumlah_balita integer NOT NULL DEFAULT 0,
    jumlah_ibu_hamil integer NOT NULL DEFAULT 0,

    version integer NOT NULL DEFAULT 1
);

CREATE INDEX posyandu_sppg_id_idx
ON posyandu(sppg_id);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_nama_length_check
CHECK (char_length(nama) <= 200);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_alamat_length_check
CHECK (char_length(alamat) <= 500);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_kecamatan_length_check
CHECK (char_length(kecamatan) <= 100);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_kelurahan_length_check
CHECK (char_length(kelurahan) <= 100);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_jumlah_balita_check
CHECK (jumlah_balita >= 0);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_jumlah_ibu_hamil_check
CHECK (jumlah_ibu_hamil >= 0);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_latitude_check
CHECK (
    latitude IS NULL OR
    latitude BETWEEN -90 AND 90
);

ALTER TABLE posyandu
ADD CONSTRAINT posyandu_longitude_check
CHECK (
    longitude IS NULL OR
    longitude BETWEEN -180 AND 180
);