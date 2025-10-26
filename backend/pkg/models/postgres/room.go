package postgres

import (
	"database/sql"
)

type RoomModel struct {
	DB *sql.DB
}

func (r *RoomModel) Create(user_id int) (int, error) {
	stmt := "INSERT INTO rooms (created_by) values ($1)"
	result, err := r.DB.Exec(stmt, user_id)

	if err != nil {
		return 0, err
	}

	room_id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(room_id), nil
}
