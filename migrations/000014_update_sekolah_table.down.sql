DROP INDEX IF EXISTS sekolah_sppg_id_idx;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_kelurahan_length_check;

ALTER TABLE sekolah
DROP COLUMN IF EXISTS updated_at;

ALTER TABLE sekolah
DROP COLUMN IF EXISTS sppg_id;

ALTER TABLE sekolah
DROP COLUMN IF EXISTS kelurahan;