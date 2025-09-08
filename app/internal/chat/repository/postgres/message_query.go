package postgres

const (
	createMessage = "create message"
)

func messageQueries() map[string]string {
	return map[string]string{
		createMessage: `
		INSERT INTO messages 
			(id, content, room_id, sender_id, created_at)
		VALUES 
			($1, $2, $3, $4, $5)`,
	}
}
