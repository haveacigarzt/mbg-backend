-- Rename kolom
ALTER TABLE sekolah
RENAME COLUMN tingkat TO kategori;

-- Hapus constraint lama terlebih dahulu
ALTER TABLE sekolah
DROP CONSTRAINT sekolah_tingkat_check;

-- Baru update data
UPDATE sekolah
SET kategori = CASE kategori
    WHEN 'SD' THEN 'SD/MI'
    WHEN 'SMP' THEN 'SMP/MTs'
    WHEN 'SMA' THEN 'SMA/SMK/MA'
END;

-- Tambahkan constraint baru
ALTER TABLE sekolah
ADD CONSTRAINT sekolah_kategori_check
CHECK (
    kategori IN (
        'PAUD/TK',
        'SD/MI',
        'SMP/MTs',
        'SMA/SMK/MA',
        'YAYASAN'
    )
);