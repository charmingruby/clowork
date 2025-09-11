CREATE TABLE IF NOT EXISTS rooms (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL UNIQUE,
    topic VARCHAR,
    created_at TIMESTAMP NOT NULL
);

CREATE INDEX IF NOT EXISTS rooms_name_uidx ON rooms(name);