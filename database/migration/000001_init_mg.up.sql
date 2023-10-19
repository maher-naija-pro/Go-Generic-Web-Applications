CREATE TABLE
    users (
        id SERIAL PRIMARY KEY,
        first_name VARCHAR(50) NOT NULL,
        last_name VARCHAR(50) NOT NULL,
        email VARCHAR(50) NOT NULL,
        password VARCHAR(500) NOT NULL,
        access_level BIGINT NOT NULL DEFAULT (1)
    );

CREATE INDEX id_user ON users (id);