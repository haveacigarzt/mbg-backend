ALTER TABLE sekolah
ADD COLUMN kelurahan text NOT NULL DEFAULT '';

ALTER TABLE sekolah
ADD COLUMN sppg_id bigint
REFERENCES sppg(id)
ON DELETE SET NULL;

ALTER TABLE sekolah
ADD COLUMN updated_at timestamptz(0) NOT NULL DEFAULT NOW();


-- Constraints

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_kelurahan_length_check
CHECK (char_length(kelurahan) <= 100);


-- Index

CREATE INDEX sekolah_sppg_id_idx
ON sekolah(sppg_id);


-- Hapus default setelah data lama terisi

ALTER TABLE sekolah
ALTER COLUMN kelurahan DROP DEFAULT;