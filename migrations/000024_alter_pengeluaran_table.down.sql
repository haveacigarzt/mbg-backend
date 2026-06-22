ALTER TABLE pengeluaran_harian
DROP COLUMN subtotal;

ALTER TABLE pengeluaran_harian
ALTER COLUMN jumlah TYPE NUMERIC(12,2)
USING jumlah::NUMERIC(12,2);

ALTER TABLE pengeluaran_harian
ADD COLUMN subtotal BIGINT GENERATED ALWAYS AS (
    (jumlah * harga_satuan)::BIGINT
) STORED;