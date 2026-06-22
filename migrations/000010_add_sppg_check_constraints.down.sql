ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_longitude_check;

ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_latitude_check;

ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_kapasitas_porsi_check;

ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_email_length_check;

ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_nomor_telepon_length_check;

ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_kepala_sppg_length_check;

ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_alamat_length_check;

ALTER TABLE sppg
DROP CONSTRAINT IF EXISTS sppg_nama_length_check;