package models

import "time"

type Message struct {
	ID            int       `db:"id_message"`
	Contenu       string    `db:"contenu"`
	DateCreation  time.Time `db:"date_creation"`
	UtilisateurID int       `db:"id_utilisateur"`
	FilID         int       `db:"id_fil"`
}
