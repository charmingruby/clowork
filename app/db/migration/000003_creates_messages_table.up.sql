CREATE TABLE IF NOT EXISTS messages (
    id VARCHAR PRIMARY KEY,
    content TEXT NOT NULL,
    room_id VARCHAR NOT NULL,
    sender_id VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,

    CONSTRAINT fk_room FOREIGN KEY (room_id) REFERENCES rooms(id),
    CONSTRAINT fk_user FOREIGN KEY (sender_id) REFERENCES users(id)
);