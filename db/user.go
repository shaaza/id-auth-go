package db

import (
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"micro-auth/domain"
)

const HASHING_COST = 10

func (db *DB) NewUser(user domain.User) (int64, error) {
	statement, err := db.Instance.Prepare("INSERT INTO users(id, username, password, first_name, last_name, phone_number) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), HASHING_COST)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	result, err := statement.Exec(uuid.NewV4(), user.Username, hashedPass, user.FirstName, user.LastName, user.PhoneNumber)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return 0, err
	}

	return rowCount, nil
}
