ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_nama_length_check;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_alamat_length_check;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_tingkat_check;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_sosmed_url_check;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_jumlah_siswa_check;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_kecamatan_length_check;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_latitude_check;

ALTER TABLE sekolah
DROP CONSTRAINT IF EXISTS sekolah_longitude_check;