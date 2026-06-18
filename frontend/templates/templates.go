package templates

import (
	"html/template"
	"log"
)

var Templates *template.Template

func InitTemplates() {
	var err error
	Templates, err = template.ParseGlob("frontend/templates/*.html")
	if err != nil {
		log.Fatalf("Erreur chargement templates : %v", err)
	}
	log.Println("Templates chargés avec succès.")
}
