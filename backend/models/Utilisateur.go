package models

import "time"

type Utilisateur struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	IsAdmin      bool      `db:"is_admin"`
	IsBanni      bool      `db:"is_banni"`
	CreatedAt    time.Time `db:"created_at"`
}
