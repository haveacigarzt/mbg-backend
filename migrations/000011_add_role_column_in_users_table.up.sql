ALTER TABLE users
ADD COLUMN role text NOT NULL DEFAULT 'stakeholder';

ALTER TABLE users
ADD CONSTRAINT users_role_check
CHECK (role IN ('admin', 'stakeholder', 'sppg', 'driver'));