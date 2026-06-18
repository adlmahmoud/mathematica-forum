package models

import "time"

type Utilisateur struct {
	ID              int       `json:"id_utilisateur" db:"id_utilisateur"`
	Username        string    `json:"nom_utilisateur" db:"nom_utilisateur"`
	Email           string    `json:"email" db:"email"`
	PasswordHash    string    `json:"mot_de_passe_hash" db:"mot_de_passe_hash"`
	Sel             string    `json:"sel" db:"sel"`
	IsAdmin         bool      `json:"is_admin" db:"is_admin"`
	IsBanni         bool      `json:"is_banni" db:"is_banni"`
	DateInscription time.Time `json:"date_inscription" db:"date_inscription"`
}
