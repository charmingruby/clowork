CREATE TABLE IF NOT EXISTS rooms (
    id VARCHAR PRIMARY KEY,
    name VARCHAR NOT NULL,
    topic VARCHAR,
    is_dm BOOLEAN NOT NULL,
    created_by VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (created_by) REFERENCES users (id)
);

CREATE INDEX IF NOT EXISTS rooms_name_uidx ON rooms(name);