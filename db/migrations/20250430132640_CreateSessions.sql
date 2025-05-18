-- migrate:up
CREATE TABLE sessions (
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	subtitle TEXT,
	description TEXT,
	date TIMESTAMP,
	published BOOLEAN,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);


-- migrate:down
DROP TABLE sessions; 

-- migrate:down

