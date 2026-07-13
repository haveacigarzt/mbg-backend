ALTER TABLE posyandu
DROP CONSTRAINT IF EXISTS fk_posyandu_kecamatan;

ALTER TABLE posyandu
DROP CONSTRAINT IF EXISTS fk_posyandu_kelurahan;

ALTER TABLE posyandu
DROP COLUMN IF EXISTS kecamatan_id;

ALTER TABLE posyandu
DROP COLUMN IF EXISTS kelurahan_id;

ALTER TABLE posyandu
ADD COLUMN IF NOT EXISTS kecamatan TEXT;

ALTER TABLE posyandu
ADD COLUMN IF NOT EXISTS kelurahan TEXT;