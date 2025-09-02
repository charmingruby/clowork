CREATE TABLE IF NOT EXISTS users (
    id VARCHAR PRIMARY KEY,
    nickname VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS users_nickname_uidx ON users(nickname);