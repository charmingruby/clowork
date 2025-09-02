CREATE TABLE IF NOT EXISTS room_members (
    id VARCHAR PRIMARY KEY,
    room_id VARCHAR NOT NULL,
    user_id VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    role VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,

    CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES rooms(id),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
);