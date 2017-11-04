-- +migrate Up
CREATE TABLE users (
    id uuid PRIMARY KEY,
    username varchar(255) UNIQUE,
    password text,
    first_name varchar(255),
    last_name varchar(255),
    phone_number varchar(255)
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY,
    user_id uuid,
    start_time TIMESTAMP,
    expiry_seconds INT,
    valid BOOLEAN
);

-- +migrate Down
DROP TABLE users;
DROP TABLE sessions;