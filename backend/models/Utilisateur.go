package models

import "time"

type Utilisateur struct {
	ID              int       `db:"id_utilisateur"`
	Username        string    `db:"nom_utilisateur"`
	Email           string    `db:"email"`
	PasswordHash    string    `db:"mot_de_passe_hash"`
	IsAdmin         bool      `db:"is_admin"`
	IsBanni         bool      `db:"is_banni"`
	DateInscription time.Time `db:"date_inscription"`
}
