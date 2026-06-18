// Package app assemble les composants de l'application et configure le routeur.
package app

import (
	"exempleInjection/api"
	"exempleInjection/config"
	"exempleInjection/controllers"
	"exempleInjection/routers"
	"exempleInjection/services"
	"exempleInjection/templates"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

// InitApp initialise l'application, charge les templates et configure le routeur.
func InitApp() *App {
	config.LoadEnv()

	templatesManager := templates.NewTemplatesManager()

	baseURL := config.GetRequiredEnv("BASE_URL")
	productApi := api.InitProductApi(baseURL)

	productService := services.InitProductService(productApi)

	productController := controllers.InitProductController(productService, templatesManager)

	router := mux.NewRouter()
	routers.RegisterAssetsRoutes(router)
	routers.RegisterProductRoutes(router, productController)

	return &App{
		Router: router,
	}
}
