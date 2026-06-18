package routes

import (
	"mathematica-forum/controllers"
	"mathematica-forum/middleware"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, authController *controllers.AuthControllers) {
	r.HandleFunc("/auth/login", authController.Login).Methods("POST")
}

func RegisterFilRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/fils", filController.ReadAll).Methods("GET")
	r.HandleFunc("/fils/{id}", filController.ReadById).Methods("GET")

	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/fils", filController.Create).Methods("POST")
	protected.HandleFunc("/fils/{id}", filController.UpdateById).Methods("PUT")
	protected.HandleFunc("/fils/{id}", filController.DeleteById).Methods("DELETE")
}

func RegisterUtilisateurRoutes(r *mux.Router, utilisateurController *controllers.UtilisateurControllers) {
	r.HandleFunc("/utilisateurs", utilisateurController.ReadAll).Methods("GET")
	r.HandleFunc("/utilisateurs/{id}", utilisateurController.ReadById).Methods("GET")
	r.HandleFunc("/utilisateurs", utilisateurController.Create).Methods("POST")

	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/utilisateurs/{id}", utilisateurController.UpdateById).Methods("PUT")
	protected.HandleFunc("/utilisateurs/{id}", utilisateurController.DeleteById).Methods("DELETE")
	protected.HandleFunc("/utilisateurs/{id}/ban", utilisateurController.BanUser).Methods("POST")
}

func RegisterMessageRoutes(r *mux.Router, messageController *controllers.MessageControllers) {
	r.HandleFunc("/messages", messageController.ReadAll).Methods("GET")
	r.HandleFunc("/messages/{id}", messageController.ReadById).Methods("GET")
	r.HandleFunc("/fils/{filId}/messages", messageController.ReadByFilId).Methods("GET")

	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/messages", messageController.Create).Methods("POST")
	protected.HandleFunc("/messages/{id}", messageController.UpdateById).Methods("PUT")
	protected.HandleFunc("/messages/{id}", messageController.DeleteById).Methods("DELETE")
}

func RegisterReactionRoutes(r *mux.Router, reactionController *controllers.ReactionController) {
	r.HandleFunc("/messages/{messageID}/reactions", reactionController.GetReactionsByMessage).Methods("GET")
	r.HandleFunc("/messages/{messageID}/score", reactionController.GetScore).Methods("GET")

	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/messages/{messageID}/reactions", reactionController.CreateReaction).Methods("POST")
	protected.HandleFunc("/reactions/{reactionID}", reactionController.DeleteReaction).Methods("DELETE")
}

func RegisterTagRoutes(r *mux.Router, tagController *controllers.TagController) {
	r.HandleFunc("/tags", tagController.GetAllTags).Methods("GET")
	r.HandleFunc("/tags/{tagID}", tagController.GetTagById).Methods("GET")
	r.HandleFunc("/tags/{tagID}/fils", tagController.GetFilsByTag).Methods("GET")

	protected := r.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)
	protected.HandleFunc("/tags", tagController.CreateTag).Methods("POST")
	protected.HandleFunc("/tags/{tagID}", tagController.DeleteTag).Methods("DELETE")
	protected.HandleFunc("/fils/{filID}/tags", tagController.AddTagToFil).Methods("POST")
}
