package dto

type CreateMessageRequest struct {
	Contenu string `json:"contenu"`
	FilID   int    `json:"fil_id"`
}

type MessageResponse struct {
	ID        int    `json:"id"`
	Contenu   string `json:"contenu"`
	Auteur    string `json:"auteur"`
	Score     int    `json:"score"`
	CreatedAt string `json:"created_at"`
}

type ReactionRequest struct {
	MessageID int    `json:"message_id"`
	Type      string `json:"type"`
}
