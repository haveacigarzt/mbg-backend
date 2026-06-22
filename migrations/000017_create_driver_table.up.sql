CREATE TABLE IF NOT EXISTS driver (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    user_id BIGINT NOT NULL,
    sppg_id BIGINT NOT NULL,

    nama TEXT NOT NULL,
    nomor_telepon TEXT NOT NULL,

    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    version INTEGER NOT NULL DEFAULT 1,

    CONSTRAINT driver_user_id_key
        UNIQUE (user_id),
    
    CONSTRAINT driver_nomor_telepon_key
      UNIQUE (nomor_telepon),

    CONSTRAINT driver_user_id_fkey
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT driver_sppg_id_fkey
        FOREIGN KEY (sppg_id)
        REFERENCES sppg(id)
        ON DELETE RESTRICT,

    CONSTRAINT driver_version_check
        CHECK (version >= 1),

    CONSTRAINT driver_nama_check
        CHECK (char_length(trim(nama)) > 0),

    CONSTRAINT driver_nomor_telepon_check
        CHECK (char_length(trim(nomor_telepon)) > 0)
);

CREATE INDEX idx_driver_user_id ON driver(user_id);
CREATE INDEX idx_driver_sppg_id ON driver(sppg_id);