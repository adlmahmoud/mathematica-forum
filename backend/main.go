// Cree le serveur

package main

import (
	"log"
	"mathematica-forum/app"
	"net/http"
)

func main() {
	app := app.InitApp()
	defer app.Close()

	// Lancement du serveur
	log.Printf("Serveur lancé : http://localhost:8080")
	serveErr := http.ListenAndServe(":8080", app.Router)
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %s", serveErr.Error())
	}
}
