ALTER TABLE balita
ALTER COLUMN berat_lahir TYPE NUMERIC(5,2)
USING berat_lahir::NUMERIC(5,2);

ALTER TABLE balita
ALTER COLUMN panjang_lahir TYPE NUMERIC(5,2)
USING panjang_lahir::NUMERIC(5,2);