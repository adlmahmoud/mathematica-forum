// Cree le serveur

package main

import (
	"log"
	"mathematica-forum/app"
	"net/http"
)

func main() {
	application := app.InitApp()
	defer application.Close()

	// Lancement du serveur
	log.Printf("Serveur lancé : http://localhost:8080")
	serveErr := http.ListenAndServe(":8080", application.Router)
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %s", serveErr.Error())
	}
}
