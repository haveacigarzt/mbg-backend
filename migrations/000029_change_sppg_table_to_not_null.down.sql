UPDATE sppg
SET
    kecamatan_id = 1,
    kelurahan_id = 1
WHERE
    kecamatan_id IS NULL
    OR kelurahan_id IS NULL;

ALTER TABLE sppg
    ALTER COLUMN nama SET NOT NULL,
    ALTER COLUMN alamat SET NOT NULL,
    ALTER COLUMN kecamatan_id SET NOT NULL,
    ALTER COLUMN kelurahan_id SET NOT NULL;