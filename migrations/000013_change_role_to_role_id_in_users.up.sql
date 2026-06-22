ALTER TABLE users
DROP COLUMN role;

ALTER TABLE users
ADD COLUMN role_id bigint NOT NULL
REFERENCES roles(id);