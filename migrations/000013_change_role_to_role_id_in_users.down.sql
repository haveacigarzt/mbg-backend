ALTER TABLE users
DROP COLUMN role_id;

ALTER TABLE users
ADD COLUMN role text NOT NULL DEFAULT 'stakeholder';