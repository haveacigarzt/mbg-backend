INSERT INTO roles (name, permissions)
VALUES
(
    'posyandu',
    ARRAY[
        'posyandu:read',
        'posyandu:write',
        'balita:read',
        'balita:write',
        'bumil:read',
        'bumil:write',
        'busui:read',
        'busui:write',
        'aps:read',
        'aps:write',
        'ats:read',
        'ats:write',
        'pengukuran_balita:read',
        'pengukuran_balita:write',
        'pengukuran_bumil:read',
        'pengukuran_bumil:write',
        'pengukuran_busui:read',
        'pengukuran_busui:write',
        'pengukuran_aps:read',
        'pengukuran_aps:write',
        'pengukuran_ats:read',
        'pengukuran_ats:write'
    ]
),
(
    'sekolah',
    ARRAY[
        'sekolah:read',
        'sekolah:write',
        'peserta_didik:read',
        'peserta_didik:write',
        'pengukuran_peserta_didik:read',
        'pengukuran_peserta_didik:write'
    ]
);

ALTER TABLE posyandu
ADD COLUMN IF NOT EXISTS user_id BIGINT;

ALTER TABLE posyandu
ADD CONSTRAINT uq_posyandu_user
UNIQUE (user_id);

ALTER TABLE posyandu
ADD CONSTRAINT fk_posyandu_user
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE sekolah
ADD COLUMN IF NOT EXISTS user_id BIGINT,
ADD COLUMN IF NOT EXISTS updated_at timestamptz(0) DEFAULT NOW();

ALTER TABLE sekolah
ADD CONSTRAINT uq_sekolah_user
UNIQUE (user_id);

ALTER TABLE sekolah
ADD CONSTRAINT fk_sekolah_user
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;