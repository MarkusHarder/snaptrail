-- migrate:up
ALTER TABLE users 
ADD  COLUMN version INT,
ADD COLUMN role TEXT;

-- migrate:down
ALTER TABLE users
DROP COLUMN version,
DROP COLUMN role;
