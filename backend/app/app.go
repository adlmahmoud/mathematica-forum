package app

import (
	"database/sql"

	"mathematica-forum/config"
	"mathematica-forum/controllers"
	"mathematica-forum/repositories"
	"mathematica-forum/routers"
	"mathematica-forum/services"

	"github.com/gorilla/mux"
)

type App struct {
	Db     *sql.DB
	Router *mux.Router
}

func InitApp() *App {
	// Charegement des variables d'environnements
	config.LoadEnv()

	// Initilisation de la connexion à la base de données
	db := config.InitDB()

	// Initilisation des repositories
	filRepository := repositories.InitFilRepository(db)
	utilisateurRepository := repositories.InitUtilisateurRepository(db)

	// Initilisation des services
	filService := services.InitFilService(filRepository)
	utilisateurService := services.InitUtilisateurService(utilisateurRepository)

	// Initilisation des controllers
	filController := controllers.InitFilController(filService)
	utilisateurController := controllers.InitUtilisateurController(utilisateurService)

	// Enregistrement des routes (avec ajout du préfix "/api/...")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	routers.RegisterFilRoutes(router, filController)
	routers.RegisterUtilisateurRoutes(router, utilisateurController)

	return &App{
		Db:     db,
		Router: router,
	}
}

func (a *App) Close() {
	if a.Db != nil {
		a.Db.Close()
	}
}
