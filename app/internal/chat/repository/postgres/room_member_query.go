package postgres

const (
	createRoomMember         = "create room member"
	roomMemberExistsInRoom   = "room member exists in room"
	findRoomMemberByIDInRoom = "find room member by id in room"
	updateRoomMemberStatus   = "update room member status"
)

func roomMemberQueries() map[string]string {
	return map[string]string{
		createRoomMember: `
		INSERT INTO room_members 
			(id, nickname, hostname, room_id, created_at)
		VALUES 
			($1, $2, $3, $4, $5)`,
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
