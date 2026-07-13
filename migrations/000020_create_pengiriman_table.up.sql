CREATE TABLE IF NOT EXISTS pengiriman (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    sppg_id BIGINT NOT NULL,
    driver_id BIGINT NULL,

    tujuan_type TEXT NOT NULL,
    tujuan_id BIGINT NOT NULL,

    status TEXT NOT NULL DEFAULT 'menunggu',

    waktu_berangkat TIMESTAMPTZ,
    waktu_selesai TIMESTAMPTZ,

    version INTEGER NOT NULL DEFAULT 1,

    CONSTRAINT pengiriman_sppg_id_fkey
        FOREIGN KEY (sppg_id)
        REFERENCES sppg(id)
        ON DELETE RESTRICT,

    CONSTRAINT pengiriman_driver_id_fkey
        FOREIGN KEY (driver_id)
        REFERENCES drivers(id)
        ON DELETE RESTRICT,

    CONSTRAINT pengiriman_version_check
        CHECK (version >= 1),

    CONSTRAINT pengiriman_status_check
        CHECK (
            status IN (
                'menunggu',
                'berangkat',
                'sampai',
                'dibatalkan'
            )
        ),

    CONSTRAINT pengiriman_tujuan_type_check
        CHECK (
            tujuan_type IN (
                'sekolah',
                'posyandu'
            )
        )
);

CREATE INDEX idx_pengiriman_sppg_id
    ON pengiriman(sppg_id);

CREATE INDEX idx_pengiriman_driver_id
    ON pengiriman(driver_id);

CREATE INDEX idx_pengiriman_tujuan
    ON pengiriman(tujuan_type, tujuan_id);

CREATE INDEX idx_pengiriman_status
    ON pengiriman(status);
