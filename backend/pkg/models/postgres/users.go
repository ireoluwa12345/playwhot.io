package postgres

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"playwhot.io/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(username, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (username, email, hash) VALUES ($1, $2, $3)`

	_, err = u.DB.Exec(stmt, username, email, string(hashedPassword))

	if err != nil {
		// var mySQLError *mysql.MySQLError
		// if errors.As(err, &mySQLError) {
		// 	if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message,
		// 		"users_uc_email") {
		// 		return models.ErrDuplicateEmail
		// 	}
		// 	if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "users_uc_username") {
		// 		return models.ErrDuplicateUsername
		// 	}
		// }
		return err
	}
	return nil
}

func (u *UserModel) Authenticate(username string, password string) (int, string, error) {
	var id int
	var hash []byte
	var email string
	// Use index-optimized query with LIMIT 1
	stmt := `SELECT id, email, hash FROM users WHERE username = $1 LIMIT 1`
	row := u.DB.QueryRow(stmt, username)
	err := row.Scan(&id, &email, &hash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, "", models.ErrInvalidCredentials
		} else {
			return 0, "", err
		}
	}

	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, "", models.ErrInvalidCredentials
		} else {
			return 0, "", err
		}
	}

	return id, email, nil
}
