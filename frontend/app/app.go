package app

import (
	"log"
	"net/http"
	"os"

	"mathematica-forum/config"
	"mathematica-forum/routers"
	"mathematica-forum/templates"
)

func Start() {
	// Charger la configuration
	config.LoadEnv()

	// Initialiser les templates (Nouveau !)
	templates.InitTemplates()

	// Configurer les routes
	routers.SetupRoutes()

	// Configurer les fichiers statiques (CSS/JS)
	fs := http.FileServer(http.Dir("frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Lancer le serveur
	port := os.Getenv("FRONTEND_PORT")
	if port == "" {
		port = "8081"
	}

	log.Println("Le serveur frontend est démarré sur http://localhost:" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Erreur fatale du serveur : ", err)
	}
}
