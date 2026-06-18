// Package main démarre l'application web et configure le serveur HTTP.
package main

import (
	"exempleInjection/app"
	"log"
	"net/http"
)

// main initialise l'application et démarre le serveur HTTP sur le port 3000.
func main() {
	app := app.InitApp()

	log.Printf("Serveur lancé : http://localhost:3000")
	serveErr := http.ListenAndServe(":3000", app.Router)
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %s", serveErr.Error())
	}
}
