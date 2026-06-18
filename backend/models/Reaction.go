package models

type Reaction struct {
	ID            int    `db:"id_reaction"`
	Type          string `db:"type_reaction"`
	UtilisateurID int    `db:"id_utilisateur"`
	MessageID     int    `db:"id_message"`
}
