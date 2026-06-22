CREATE TABLE IF NOT EXISTS kecamatan (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  name text NOT NULL UNIQUE,
  version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS kelurahan (
  id bigserial PRIMARY KEY,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  kecamatan_id bigint NOT NULL REFERENCES kecamatan(id) ON DELETE CASCADE,
  name text NOT NULL, 
  version integer NOT NULL DEFAULT 1,
  UNIQUE(kecamatan_id, name)
);

CREATE INDEX kelurahan_kecamatan_id_idx
ON kelurahan(kecamatan_id);