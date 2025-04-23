-- migrate:up
CREATE TABLE hellos (
	id SERIAL PRIMARY KEY,
	text TEXT NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP
);

INSERT INTO hellos (text, created_at, updated_at) VALUES

('This is a hello message from the database!', 
 '2024-11-05 16:30:00', 
 '2024-11-05 17:20:00');

-- migrate:down
DROP TABLE hellos; 
