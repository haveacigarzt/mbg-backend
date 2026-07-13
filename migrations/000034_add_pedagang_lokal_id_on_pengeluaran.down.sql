ALTER TABLE pengeluaran_harian
DROP CONSTRAINT IF EXISTS pengeluaran_harian_pedagang_lokal_id_fkey;

ALTER TABLE pengeluaran_harian
DROP COLUMN IF EXISTS pedagang_lokal_id;