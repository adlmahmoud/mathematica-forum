package models

// Type possible : "like", "dislike"
type Reaction struct {
	ID            int    `db:"id"`
	Type          string `db:"type"`
	UtilisateurID int    `db:"utilisateur_id"`
	MessageID     int    `db:"message_id"`
}
