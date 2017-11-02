-- +migrate Up
CREATE TABLE users (
    id int PRIMARY KEY,
    username varchar(255),
    password varchar(255)
);
-- +migrate Down
DROP TABLE users;
