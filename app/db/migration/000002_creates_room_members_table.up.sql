CREATE TABLE IF NOT EXISTS room_members (
    id VARCHAR PRIMARY KEY,
    nickname VARCHAR NOT NULL,
    hostname VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL
);