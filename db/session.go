package db

import (
	"log"
	"micro-auth/domain"
	"net/http"
)

func (db *DB) NewSession(session domain.Session) (int64, *domain.Error) {
	statement, err := db.Instance.Prepare("INSERT INTO sessions(id, user_id, start_time, expiry_seconds, valid) VALUES($1, $2, $3, $4, $5)")
	if err != nil {
		log.Fatal(err)
		return 0, domain.NewError(err.Error(), http.StatusInternalServerError)
	}

	result, err := statement.Exec(session.Id, session.UserId, session.StartTime, session.Expiry, session.Valid)
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
