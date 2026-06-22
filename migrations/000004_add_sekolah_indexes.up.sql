CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE INDEX IF NOT EXISTS sekolah_nama_trgm_idx
ON sekolah
USING GIN (nama gin_trgm_ops);