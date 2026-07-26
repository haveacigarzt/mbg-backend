-- Hapus constraint baru
ALTER TABLE sekolah
DROP CONSTRAINT sekolah_kategori_check;

-- Ubah data kembali
UPDATE sekolah
SET kategori = CASE kategori
    WHEN 'SD/MI' THEN 'SD'
    WHEN 'SMP/MTs' THEN 'SMP'
    WHEN 'SMA/SMK/MA' THEN 'SMA'
    ELSE kategori
END;

-- Rename kolom kembali
ALTER TABLE sekolah
RENAME COLUMN kategori TO tingkat;

-- Tambahkan constraint lama
ALTER TABLE sekolah
ADD CONSTRAINT sekolah_tingkat_check
CHECK (
    tingkat IN (
        'SD',
        'SMP',
        'SMA'
    )
);