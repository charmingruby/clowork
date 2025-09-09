package postgres

const (
	createMessage        = "create message"
	listMessagesByRoomID = "list messages by room id"
)

func messageQueries() map[string]string {
	return map[string]string{
		createMessage: `
		INSERT INTO messages 
			(id, content, room_id, sender_id, created_at)
		VALUES 
			($1, $2, $3, $4, $5)`,
		listMessagesByRoomID: `
		SELECT * FROM messages
		WHERE room_id = $1
		ORDER BY created_at DESC
		OFFSET $2
		LIMIT $3`,
	}
}
