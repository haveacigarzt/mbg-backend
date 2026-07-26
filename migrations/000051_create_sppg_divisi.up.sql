CREATE TABLE sppg_divisi (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),

    sppg_id BIGINT NOT NULL REFERENCES sppg(id) ON DELETE CASCADE,
    divisi_id BIGINT NOT NULL REFERENCES divisi_sppg(id),

    jumlah_sdm INTEGER NOT NULL DEFAULT 0,

    version INTEGER NOT NULL DEFAULT 1,

    CONSTRAINT sppg_divisi_jumlah_sdm_check CHECK (jumlah_sdm >= 0),
    CONSTRAINT sppg_divisi_unique UNIQUE (sppg_id, divisi_id)
);