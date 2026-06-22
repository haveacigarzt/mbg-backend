CREATE TABLE tracking (
    id BIGSERIAL PRIMARY KEY,

    pengiriman_id BIGINT NOT NULL
        REFERENCES pengiriman(id) ON DELETE CASCADE,

    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,

    speed DOUBLE PRECISION,
    accuracy DOUBLE PRECISION,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);