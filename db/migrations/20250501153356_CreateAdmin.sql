-- migrate:up

CREATE TABLE users(
	id SERIAL PRIMARY KEY,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

-- migrate:down
DROP TABLE users; 
