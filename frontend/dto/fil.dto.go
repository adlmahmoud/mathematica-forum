package dto

type CreateFilRequest struct {
	Titre string `json:"titre"`
	Tags  []int  `json:"tags,omitempty"`
}

type FilResponse struct {
	ID        int      `json:"id"`
	Titre     string   `json:"titre"`
	Auteur    string   `json:"auteur"`
	Etat      string   `json:"etat"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
}
