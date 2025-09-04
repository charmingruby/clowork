CREATE TABLE IF NOT EXISTS room_members (
    id VARCHAR PRIMARY KEY,
    nickname VARCHAR NOT NULL,
    host_name VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES rooms(id)
);