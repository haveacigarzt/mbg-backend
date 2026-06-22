ALTER TABLE sppg
ADD COLUMN kecamatan_id BIGINT,
ADD COLUMN kelurahan_id BIGINT;

-- update data lama dengan nilai default (misalnya 0 atau NULL)

ALTER TABLE sppg
ALTER COLUMN kecamatan_id SET NOT NULL,
ALTER COLUMN kelurahan_id SET NOT NULL;

ALTER TABLE sppg
ADD CONSTRAINT fk_sppg_kecamatan
FOREIGN KEY (kecamatan_id)
REFERENCES kecamatan(id);

ALTER TABLE sppg
ADD CONSTRAINT fk_sppg_kelurahan
FOREIGN KEY (kelurahan_id)
REFERENCES kelurahan(id);





ALTER TABLE sekolah
ADD COLUMN kecamatan_id BIGINT,
ADD COLUMN kelurahan_id BIGINT;

-- update data lama dengan nilai default (misalnya 0 atau NULL)

ALTER TABLE sekolah
ALTER COLUMN kecamatan_id SET NOT NULL,
ALTER COLUMN kelurahan_id SET NOT NULL;

ALTER TABLE sekolah
ADD CONSTRAINT fk_sekolah_kecamatan
FOREIGN KEY (kecamatan_id)
REFERENCES kecamatan(id);

ALTER TABLE sekolah
ADD CONSTRAINT fk_sekolah_kelurahan
FOREIGN KEY (kelurahan_id)
REFERENCES kelurahan(id);

-- Setelah route dan model sudah diperbarui, kita bisa menghapus kolom kecamatan dan kelurahan yang lama

ALTER TABLE sekolah
DROP COLUMN IF EXISTS kecamatan,
DROP COLUMN IF EXISTS kelurahan;





ALTER TABLE posyandu
ADD COLUMN kecamatan_id BIGINT,
ADD COLUMN kelurahan_id BIGINT;

-- update data lama dengan nilai default (misalnya 0 atau NULL)

ALTER TABLE posyandu
ALTER COLUMN kecamatan DROP NOT NULL;

ALTER TABLE posyandu
ALTER COLUMN kelurahan DROP NOT NULL;

ALTER TABLE posyandu
ALTER COLUMN kecamatan_id SET NOT NULL,
ALTER COLUMN kelurahan_id SET NOT NULL;

ALTER TABLE posyandu
ADD CONSTRAINT fk_posyandu_kecamatan
FOREIGN KEY (kecamatan_id)
REFERENCES kecamatan(id);

ALTER TABLE posyandu
ADD CONSTRAINT fk_posyandu_kelurahan
FOREIGN KEY (kelurahan_id)
REFERENCES kelurahan(id);

-- Setelah route dan model sudah diperbarui, kita bisa menghapus kolom kecamatan dan kelurahan yang lama

ALTER TABLE posyandu
DROP COLUMN IF EXISTS kecamatan,
DROP COLUMN IF EXISTS kelurahan;