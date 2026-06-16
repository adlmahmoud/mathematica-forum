package routers

import (
	"mathematica-forum/controllers"

	"github.com/gorilla/mux"
)

func RegisterFilRoutes(r *mux.Router, filController *controllers.FilControllers) {
	r.HandleFunc("/fils", filController.ReadAll).Methods("GET")
	r.HandleFunc("/fils/{id}", filController.ReadById).Methods("GET")
	r.HandleFunc("/fils", filController.Create).Methods("POST")
	r.HandleFunc("/fils/{id}", filController.UpdateById).Methods("PUT")
	r.HandleFunc("/fils/{id}", filController.DeleteById).Methods("DELETE")
}

func RegisterUtilisateurRoutes(r *mux.Router, utilisateurController *controllers.UtilisateurControllers) {
	r.HandleFunc("/utilisateurs", utilisateurController.ReadAll).Methods("GET")
	r.HandleFunc("/utilisateurs/{id}", utilisateurController.ReadById).Methods("GET")
	r.HandleFunc("/utilisateurs", utilisateurController.Create).Methods("POST")
	r.HandleFunc("/utilisateurs/{id}", utilisateurController.UpdateById).Methods("PUT")
	r.HandleFunc("/utilisateurs/{id}", utilisateurController.DeleteById).Methods("DELETE")
}
