package models

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel  `bun:"table:users"`
	ID             int64 `bun:",pk,autoincrement"`
	First_name     string
	Last_name      string
	Roles          string
	Hashed_passkey string
	Username       string
	Email          string
}
