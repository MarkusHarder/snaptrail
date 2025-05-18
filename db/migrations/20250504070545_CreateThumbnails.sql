-- migrate:up
CREATE TABLE thumbnails (
	id SERIAL PRIMARY KEY,
  session_id  INTEGER REFERENCES sessions(id) ON DELETE CASCADE,
  filename TEXT,
  mime_type TEXT,
  data BYTEA,
  camera_model TEXT,
  make TEXT,
  lens_model TEXT,
  exposure TEXT,
  date_time TEXT,
  aperture DOUBLE PRECISION,
  iso INTEGER,
  focal_length DOUBLE PRECISION,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- migrate:down

DROP TABLE thumbnails;
