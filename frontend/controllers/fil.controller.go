package controllers

import (
	"forum-frontend/services"
	"forum-frontend/templates"
	"net/http"
)

func ShowHome(w http.ResponseWriter, r *http.Request) {
	fils, err := services.GetAllFils()
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des fils", http.StatusInternalServerError)
		return
	}

	templates.Templates.ExecuteTemplate(w, "home.html", fils)
}
