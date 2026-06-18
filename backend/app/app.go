package app

import (
	"database/sql"

	"mathematica-forum/config"
	"mathematica-forum/controllers"
	"mathematica-forum/repositories"
	"mathematica-forum/routes"
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
	messageRepository := repositories.InitMessageRepository(db)
	reactionRepository := repositories.InitReactionRepository(db)
	tagRepository := repositories.InitTagRepository(db)

	// Initilisation des services
	filService := services.InitFilService(filRepository)
	utilisateurService := services.InitUtilisateurService(utilisateurRepository)
	authService := services.InitAuthService(utilisateurRepository)
	messageService := services.InitMessageService(messageRepository)
	reactionService := services.InitReactionService(reactionRepository)
	tagService := services.InitTagService(tagRepository)

	// Initilisation des controllers
	filController := controllers.InitFilController(filService)
	utilisateurController := controllers.InitUtilisateurController(utilisateurService)
	authController := controllers.InitAuthController(authService)
	messageController := controllers.InitMessageController(messageService)
	reactionController := controllers.InitReactionController(reactionService)
	tagController := controllers.InitTagController(tagService)

	// Enregistrement des routes (avec ajout du préfix "/api/...")
	router := mux.NewRouter().PathPrefix("/api").Subrouter()

	routes.RegisterAuthRoutes(router, authController)
	routes.RegisterFilRoutes(router, filController)
	routes.RegisterUtilisateurRoutes(router, utilisateurController)
	routes.RegisterMessageRoutes(router, messageController)
	routes.RegisterReactionRoutes(router, reactionController)
	routes.RegisterTagRoutes(router, tagController)

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
