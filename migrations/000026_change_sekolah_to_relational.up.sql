ALTER TABLE sekolah
DROP COLUMN IF EXISTS kecamatan;

ALTER TABLE sekolah
DROP COLUMN IF EXISTS kelurahan;

ALTER TABLE sekolah
ADD COLUMN IF NOT EXISTS kecamatan_id BIGINT;

ALTER TABLE sekolah
ADD COLUMN IF NOT EXISTS kelurahan_id BIGINT;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS fk_sekolah_kecamatan;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS fk_sekolah_kelurahan;

ALTER TABLE sekolah
ADD CONSTRAINT fk_sekolah_kecamatan
FOREIGN KEY (kecamatan_id) REFERENCES kecamatan(id);

ALTER TABLE sekolah
ADD CONSTRAINT fk_sekolah_kelurahan
FOREIGN KEY (kelurahan_id) REFERENCES kelurahan(id);