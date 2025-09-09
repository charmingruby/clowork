package postgres

const (
	createRoomMember         = "create room member"
	findRoomMemberByIDInRoom = "find room member by id in room"
	listMembersByRoomID      = "list members by room id"
	roomMemberExistsInRoom   = "room member exists in room"
	updateRoomMemberStatus   = "update room member status"
)

func roomMemberQueries() map[string]string {
	return map[string]string{
		createRoomMember: `
		INSERT INTO room_members 
			(id, nickname, hostname, room_id, created_at)
		VALUES 
			($1, $2, $3, $4, $5)`,
		listMembersByRoomID: `
		SELECT * FROM room_members
		WHERE room_id = $1
		ORDER BY created_at DESC
		OFFSET $2
		LIMIT $3`,
		roomMemberExistsInRoom: `
		SELECT * FROM room_members
		WHERE 
			room_id = $1 AND 
			nickname = $2 AND 
			hostname = $3`,
		findRoomMemberByIDInRoom: `
		SELECT * FROM room_members
		WHERE 
			room_id = $1 AND 
			id = $2`,
		updateRoomMemberStatus: `
		UPDATE room_members
		SET 
			status = $1,
			updated_at = $2
		WHERE 
			id = $3`,
	}
}
