package models

import "time"

type Message struct {
	ID            int       `db:"id"`
	Contenu       string    `db:"contenu"`
	CreatedAt     time.Time `db:"created_at"`
	UtilisateurID int       `db:"utilisateur_id"`
	FilID         int       `db:"fil_id"`
}
