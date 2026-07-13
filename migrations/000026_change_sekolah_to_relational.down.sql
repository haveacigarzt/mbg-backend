-- Hapus foreign key dulu (wajib urutan benar)
ALTER TABLE sekolah DROP CONSTRAINT IF EXISTS fk_sekolah_kecamatan;
ALTER TABLE sekolah DROP CONSTRAINT IF EXISTS fk_sekolah_kelurahan;

-- Hapus kolom relational
ALTER TABLE sekolah DROP COLUMN IF EXISTS kecamatan_id;
ALTER TABLE sekolah DROP COLUMN IF EXISTS kelurahan_id;

-- Kembalikan ke model lama (TEXT)
ALTER TABLE sekolah ADD COLUMN IF NOT EXISTS kecamatan TEXT;
ALTER TABLE sekolah ADD COLUMN IF NOT EXISTS kelurahan TEXT;