ALTER TABLE sppg
ADD CONSTRAINT sppg_nama_length_check
CHECK (char_length(nama) <= 200);

ALTER TABLE sppg
ADD CONSTRAINT sppg_alamat_length_check
CHECK (char_length(alamat) <= 500);

ALTER TABLE sppg
ADD CONSTRAINT sppg_kepala_sppg_length_check
CHECK (char_length(kepala_sppg) <= 200);

ALTER TABLE sppg
ADD CONSTRAINT sppg_nomor_telepon_length_check
CHECK (
    nomor_telepon IS NULL OR
    char_length(nomor_telepon) <= 20
);

ALTER TABLE sppg
ADD CONSTRAINT sppg_email_length_check
CHECK (
    email IS NULL OR
    char_length(email) <= 255
);

ALTER TABLE sppg
ADD CONSTRAINT sppg_kapasitas_porsi_check
CHECK (kapasitas_porsi >= 0);

ALTER TABLE sppg
ADD CONSTRAINT sppg_latitude_check
CHECK (
    latitude IS NULL OR
    latitude BETWEEN -90 AND 90
);

ALTER TABLE sppg
ADD CONSTRAINT sppg_longitude_check
CHECK (
    longitude IS NULL OR
    longitude BETWEEN -180 AND 180
);