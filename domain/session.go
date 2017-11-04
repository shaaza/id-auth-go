package domain

import "time"

type Session struct {
	Id        string
	UserId    string
	StartTime time.Time
	Expiry    int
	Valid     bool
}
