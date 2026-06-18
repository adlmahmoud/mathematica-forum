package app

import (
	"database/sql"
	"net/http"

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

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func InitApp() *App {
	config.LoadEnv()

	db := config.InitDB()

	filRepository := repositories.InitFilRepository(db)
	utilisateurRepository := repositories.InitUtilisateurRepository(db)
	messageRepository := repositories.InitMessageRepository(db)
	reactionRepository := repositories.InitReactionRepository(db)
	tagRepository := repositories.InitTagRepository(db)

	filService := services.InitFilService(filRepository)
	utilisateurService := services.InitUtilisateurService(utilisateurRepository)
	authService := services.InitAuthService(db)
	messageService := services.InitMessageService(messageRepository)
	reactionService := services.InitReactionService(reactionRepository)
	tagService := services.InitTagService(tagRepository)

	filController := controllers.InitFilController(filService)
	utilisateurController := controllers.InitUtilisateurController(utilisateurService)
	authController := controllers.InitAuthController(authService)
	messageController := controllers.InitMessageController(messageService)
	reactionController := controllers.InitReactionController(reactionService)
	tagController := controllers.InitTagController(tagService)

	baseRouter := mux.NewRouter()
	baseRouter.Use(corsMiddleware)

	apiRouter := baseRouter.PathPrefix("/api").Subrouter()

	routes.RegisterAuthRoutes(apiRouter, authController)
	routes.RegisterFilRoutes(apiRouter, filController)
	routes.RegisterUtilisateurRoutes(apiRouter, utilisateurController)
	routes.RegisterMessageRoutes(apiRouter, messageController)
	routes.RegisterReactionRoutes(apiRouter, reactionController)
	routes.RegisterTagRoutes(apiRouter, tagController)

	return &App{
		Db:     db,
		Router: baseRouter,
	}
}

func (a *App) Close() {
	if a.Db != nil {
		a.Db.Close()
	}
}
