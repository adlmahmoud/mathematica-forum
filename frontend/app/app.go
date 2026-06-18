package app

import (
	"log"
	"net/http"
	"os"

	"mathematica-forum/config"
	"mathematica-forum/routers"
)

func Start() {
	config.LoadConfig()

	routers.SetupRoutes()

	fs := http.FileServer(http.Dir("frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

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
