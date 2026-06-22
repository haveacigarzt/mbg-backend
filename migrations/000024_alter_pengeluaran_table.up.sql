ALTER TABLE pengeluaran_harian
DROP COLUMN subtotal;

ALTER TABLE pengeluaran_harian
ALTER COLUMN jumlah TYPE BIGINT
USING jumlah::BIGINT;

ALTER TABLE pengeluaran_harian
ADD COLUMN subtotal BIGINT GENERATED ALWAYS AS (
    jumlah * harga_satuan
) STORED;