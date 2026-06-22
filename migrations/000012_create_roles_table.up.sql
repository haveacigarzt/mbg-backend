CREATE TABLE roles (
    id bigserial PRIMARY KEY,
    name text NOT NULL UNIQUE,
    permissions text[] NOT NULL DEFAULT '{}'
);

INSERT INTO roles (name, permissions)
VALUES
(
    'admin',
    ARRAY[
        'akunsppg:write',
        'akunsppg:read',
        'sppg:read',
        'sekolah:read',
        'posyandu:read',
        'pengiriman:read',
        'tracking:read'
    ]
),
(
    'stakeholder',
    ARRAY[
        'sppg:read',
        'sekolah:read',
        'posyandu:read',
        'pengiriman:read',
        'tracking:read'
    ]
),
(
    'sppg',
    ARRAY[
        'sppg:read',
        'sppg:write',
        'sekolah:read',
        'sekolah:write',
        'posyandu:read',
        'posyandu:write',
        'pengiriman:read',
        'tracking:read',
        'driver:read',
        'driver:write'
    ]
),
(
    'driver',
    ARRAY[
        'pengiriman:read',
        'pengiriman:write',
        'tracking:read',
        'tracking:write'
    ]
);