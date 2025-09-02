package postgres

const (
	findUserByID       = "find user by id"
	findUserByNickname = "find user by nickname"
	createUser         = "create user"
)

func userQueries() map[string]string {
	return map[string]string{
		findUserByID:       `SELECT * FROM users WHERE id = $1`,
		findUserByNickname: `SELECT * FROM users WHERE nickname = $1`,
		createUser:         `INSERT INTO users (id, nickname, password, created_at) VALUES ($1, $2, $3, $4)`,
	}
}
