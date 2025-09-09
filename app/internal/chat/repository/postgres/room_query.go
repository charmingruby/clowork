package postgres

const (
	createRoom     = "create room"
	findRoomByName = "find room by name"
	findRoomByID   = "find room by id"
	listRooms      = "list rooms"
)

func roomQueries() map[string]string {
	return map[string]string{
		createRoom: `
		INSERT INTO rooms 
			(id, name, topic, created_at)
		VALUES 
			($1, $2, $3, $4)`,
		findRoomByName: `
		SELECT * FROM rooms
		WHERE name = $1`,
		findRoomByID: `
		SELECT * FROM rooms
		WHERE id = $1`,
		listRooms: `
		SELECT * FROM rooms
		ORDER BY created_at DESC
		LIMIT $1
		OFFSET $2`,
	}
}
