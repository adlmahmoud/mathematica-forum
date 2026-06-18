package models

import "time"

type Fil struct {
	ID            int       `db:"id_fil"`
	Titre         string    `db:"titre"`
	Statut        string    `db:"statut"`
	DateCreation  time.Time `db:"date_creation"`
	UtilisateurID int       `db:"id_utilisateur"`
}
