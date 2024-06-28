CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    poster VARCHAR(255)
);

CREATE TABLE movies_to_watch (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    poster_path VARCHAR(255) NOT NULL,
    time_added TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



