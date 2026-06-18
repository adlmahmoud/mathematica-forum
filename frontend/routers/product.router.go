// Package routers configure les routes de l'application.
package routers

import (
	"exempleInjection/controllers"

	"github.com/gorilla/mux"
)

// RegisterProductRoutes enregistre les routes liées aux produits pour l'interface web.
func RegisterProductRoutes(r *mux.Router, productController *controllers.ProductControllers) {
	r.HandleFunc("/", productController.DisplayList).Methods("GET")
	r.HandleFunc("/product/create", productController.CreateForm).Methods("GET")
	r.HandleFunc("/product/{id}", productController.DisplayById).Methods("GET")
	r.HandleFunc("/product", productController.Create).Methods("POST")
}
