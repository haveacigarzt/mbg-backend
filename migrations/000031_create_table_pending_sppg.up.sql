-- +migrate Up

CREATE TABLE sppg_invitations (
    id BIGSERIAL PRIMARY KEY,

    token TEXT NOT NULL UNIQUE,

    nama_sppg TEXT NOT NULL,

    expires_at TIMESTAMPTZ,

    used_at TIMESTAMPTZ,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE pending_sppg (
    id BIGSERIAL PRIMARY KEY,

    -- data akun
    name TEXT NOT NULL,
    email CITEXT NOT NULL UNIQUE,
    password_hash BYTEA NOT NULL,

    -- data sppg
    nama TEXT NOT NULL,
    alamat TEXT NOT NULL,
    sosmed_url TEXT[],
    kepala_sppg TEXT,
    nomor_telepon TEXT,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,
    kapasitas_porsi INTEGER NOT NULL DEFAULT 0,
    kecamatan_id BIGINT,
    kelurahan_id BIGINT,

    -- informasi undangan
    sppg_invitations_id BIGINT NOT NULL,

    status TEXT NOT NULL DEFAULT 'pending'
      CHECK (status IN ('pending', 'approved', 'rejected')),
    approved_by BIGINT,
    approved_at TIMESTAMP(0) WITH TIME ZONE,
    rejected_reason TEXT,

    -- audit
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),

    CONSTRAINT pending_sppg_nama_length_check
        CHECK (char_length(nama) <= 200),

    CONSTRAINT pending_sppg_alamat_length_check
        CHECK (char_length(alamat) <= 500),

    CONSTRAINT pending_sppg_kepala_sppg_length_check
        CHECK (kepala_sppg IS NULL OR char_length(kepala_sppg) <= 200),

    CONSTRAINT pending_sppg_nomor_telepon_length_check
        CHECK (nomor_telepon IS NULL OR char_length(nomor_telepon) <= 20),

    CONSTRAINT pending_sppg_kapasitas_porsi_check
        CHECK (kapasitas_porsi >= 0),

    CONSTRAINT pending_sppg_latitude_check
        CHECK (
            latitude IS NULL OR
            (latitude >= -90 AND latitude <= 90)
        ),

    CONSTRAINT pending_sppg_longitude_check
        CHECK (
            longitude IS NULL OR
            (longitude >= -180 AND longitude <= 180)
        ),

    CONSTRAINT fk_pending_sppg_invitation
      FOREIGN KEY (sppg_invitations_id)
      REFERENCES sppg_invitations(id),
    
    CONSTRAINT fk_pending_sppg_approved_by
      FOREIGN KEY (approved_by)
      REFERENCES users(id),

    CONSTRAINT fk_pending_sppg_kecamatan
        FOREIGN KEY (kecamatan_id)
        REFERENCES kecamatan(id),

    CONSTRAINT fk_pending_sppg_kelurahan
        FOREIGN KEY (kelurahan_id)
        REFERENCES kelurahan(id)
);

CREATE INDEX idx_pending_sppg_kecamatan_id
    ON pending_sppg(kecamatan_id);

CREATE INDEX idx_pending_sppg_kelurahan_id
    ON pending_sppg(kelurahan_id);