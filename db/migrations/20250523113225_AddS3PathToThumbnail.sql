-- migrate:up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
ALTER TABLE thumbnails  DROP data;
ALTER TABLE thumbnails ADD path TEXT;

-- migrate:down
ALTER TABLE thumbnails  ADD data BYTEA;
ALTER TABLE thumbnails DROP path;
