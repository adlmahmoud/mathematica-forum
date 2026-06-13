package models

import "time"

type Fil struct {
	ID            int       `db:"id"`
	Titre         string    `db:"titre"`
	Statut        string    `db:"statut"`
	CreatedAt     time.Time `db:"created_at"`
	UtilisateurID int       `db:"utilisateur_id"`
}
