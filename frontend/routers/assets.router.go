// Package routers configure les routes HTTP de l'application.
package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterAssetsRoutes expose le répertoire ./assets/ pour servir des fichiers statiques.
func RegisterAssetsRoutes(r *mux.Router) {
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
}
