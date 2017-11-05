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

func (db *DB) GetValidSession(username string) (domain.Session, *domain.Error) {
	session := domain.Session{}
	rows, err := db.Instance.Query("SELECT id FROM sessions WHERE user_id = $1 AND valid = true", username)
	if err != nil {
		log.Fatal(err)
		return session, domain.NewError(err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()

	rowCnt := 0
	for rows.Next() {
		err := rows.Scan(&session.Id)
		if err != nil {
			log.Fatal(err)
		}
		rowCnt++
	}

	if rowCnt == 0 {
		return session, domain.NewError("given session ID was not found", http.StatusNotFound)
	}

	return session, nil
}

func (db *DB) InvalidateSession(id string) (int64, *domain.Error) {
	statement, err := db.Instance.Prepare("UPDATE sessions SET valid = false WHERE id = $1")
	if err != nil {
		log.Fatal(err)
		return 0, domain.NewError(err.Error(), http.StatusInternalServerError)
	}

	result, err := statement.Exec(id)
	if err != nil {
		log.Fatal(err)
		return 0, domain.NewError(err.Error(), http.StatusNotFound)
	}

	rowCount, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
		return 0, domain.NewError(err.Error(), http.StatusInternalServerError)
	}

	return rowCount, nil
}
