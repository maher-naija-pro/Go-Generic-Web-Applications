CREATE TABLE users
(
    id             BIGINT PRIMARY KEY,
    first_name       VARCHAR(50) NOT NULL,
    last_name       VARCHAR(50) NOT NULL,
    email       VARCHAR(50) NOT NULL,
    password   VARCHAR(50) NOT NULL,
    access_level   BIGINT  NOT NULL DEFAULT (1)
);


