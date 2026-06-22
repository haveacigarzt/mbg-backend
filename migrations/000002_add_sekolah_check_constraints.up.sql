ALTER TABLE sekolah
ADD CONSTRAINT sekolah_nama_length_check
CHECK (char_length(nama) <= 200);

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_alamat_length_check
CHECK (char_length(alamat) <= 500);

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_tingkat_check
CHECK (tingkat IN ('SD', 'SMP', 'SMA'));

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_sosmed_url_check
CHECK (
    sosmed_url IS NULL OR
    char_length(sosmed_url) <= 500
);

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_jumlah_siswa_check
CHECK (jumlah_siswa > 0);

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_kecamatan_length_check
CHECK (char_length(kecamatan) <= 100);

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_latitude_check
CHECK (latitude >= -90 AND latitude <= 90);

ALTER TABLE sekolah
ADD CONSTRAINT sekolah_longitude_check
CHECK (longitude >= -180 AND longitude <= 180);