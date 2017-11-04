package db

import (
	"github.com/satori/go.uuid"
	"log"
	"micro-auth/domain"
	"net/http"
)

func (db *DB) NewUser(user domain.User) (int64, *domain.Error) {
	statement, err := db.Instance.Prepare("INSERT INTO users(id, username, password, first_name, last_name, phone_number) VALUES($1, $2, $3, $4, $5, $6)")
	if err != nil {
		log.Fatal(err)
		return 0, domain.NewError(err.Error(), http.StatusInternalServerError)
	}

	result, err := statement.Exec(uuid.NewV4(), user.Username, user.Password, user.FirstName, user.LastName, user.PhoneNumber)
	if err != nil {
		log.Fatal(err)
		return 0, domain.NewError(err.Error(), http.StatusConflict)
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return 0, domain.NewError(err.Error(), http.StatusInternalServerError)
	}

	return rowCount, nil
}

func (db *DB) GetUser(username string) (domain.User, *domain.Error) {
	user := domain.User{}
	rows, err := db.Instance.Query("SELECT username, password, first_name, last_name, phone_number FROM users WHERE username = $1", username)
	if err != nil {
		log.Fatal(err)
		return user, domain.NewError(err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	rowCnt := 0
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Password, &user.FirstName, &user.LastName, &user.PhoneNumber)
		if err != nil {
			log.Fatal(err)
		}
		rowCnt++
	}

	if rowCnt == 0 {
		return user, domain.NewError("user with given username was not found", http.StatusNotFound)
	}

	return user, nil
}
