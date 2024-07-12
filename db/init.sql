CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE movies_to_watch (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    poster_path VARCHAR(255) NOT NULL,
    time_added TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    watched BOOLEAN DEFAULT FALSE,
    user_id INT NOT NULL REFERENCES users(id)
);

-- read only user creation
CREATE USER movie_db_ro WITH PASSWORD 'DevPass123';

GRANT CONNECT ON DATABASE my_movie_db TO movie_db_ro;
GRANT USAGE ON SCHEMA public TO movie_db_ro;

GRANT SELECT ON ALL TABLES IN SCHEMA public TO movie_db_ro;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO movie_db_ro;

-- read write user creation
CREATE USER movie_db_rw WITH PASSWORD 'DevPass123';

GRANT CONNECT ON DATABASE my_movie_db TO movie_db_rw;
GRANT USAGE ON SCHEMA public TO movie_db_rw;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO movie_db_rw;

ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO movie_db_rw;

