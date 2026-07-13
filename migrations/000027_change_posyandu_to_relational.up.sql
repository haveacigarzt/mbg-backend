ALTER TABLE posyandu
DROP COLUMN IF EXISTS kecamatan;

ALTER TABLE posyandu
DROP COLUMN IF EXISTS kelurahan;

ALTER TABLE posyandu
ADD COLUMN IF NOT EXISTS kecamatan_id BIGINT NOT NULL;

ALTER TABLE posyandu
ADD COLUMN IF NOT EXISTS kelurahan_id BIGINT NOT NULL;

ALTER TABLE posyandu
DROP CONSTRAINT IF EXISTS fk_posyandu_kecamatan;

ALTER TABLE posyandu
DROP CONSTRAINT IF EXISTS fk_posyandu_kelurahan;

ALTER TABLE posyandu
ADD CONSTRAINT fk_posyandu_kecamatan
FOREIGN KEY (kecamatan_id) REFERENCES kecamatan(id);

ALTER TABLE posyandu
ADD CONSTRAINT fk_posyandu_kelurahan
FOREIGN KEY (kelurahan_id) REFERENCES kelurahan(id);