package routers

import (
	"forum-frontend/controllers"
	"forum-frontend/middleware"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			controllers.ShowLogin(w, r)
		} else if r.Method == http.MethodPost {
			controllers.ProcessLogin(w, r)
		}
	})

	http.HandleFunc("/home", middleware.RequireAuth(controllers.ShowHome))
}
