CREATE TABLE IF NOT EXISTS members (
    id VARCHAR PRIMARY KEY,
    nickname VARCHAR NOT NULL,
    hostname VARCHAR NOT NULL,
    room_id VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL,

    CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES rooms (id)
);