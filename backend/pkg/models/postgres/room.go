package postgres

import (
	"database/sql"
)

type RoomModel struct {
	DB *sql.DB
}

func (r *RoomModel) Create(user_id int, name string, maxPlayers int, isPrivate bool, password string) (int, error) {
	stmt := "INSERT INTO rooms (created_by, name, max_players, is_private, password) values ($1, $2, $3, $4, $5) RETURNING room_id"
	var roomId int
	err := r.DB.QueryRow(stmt, user_id, name, maxPlayers, isPrivate, password).Scan(&roomId)

	if err != nil {
		return 0, err
	}

	return roomId, nil
}
