ALTER TABLE pengeluaran_harian
ADD COLUMN pedagang_lokal_id BIGINT;

ALTER TABLE pengeluaran_harian
ADD CONSTRAINT pengeluaran_harian_pedagang_lokal_id_fkey
FOREIGN KEY (pedagang_lokal_id)
REFERENCES pedagang_lokal(id)
ON DELETE SET NULL;