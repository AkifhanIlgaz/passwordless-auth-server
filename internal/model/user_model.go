package model

import "time"

type User struct {
	ID        string    `redis:"-" json:"id"`
	Email     string    `json:"email" redis:"email"`
	Phone     string    `json:"phone" redis:"phone"`
	LastLogin time.Time `json:"last_login" redis:"last_login"`
}
