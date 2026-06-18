// Une fois que le service renvois le token JWT au controleur, il place dans un cookie HTTP
package controllers

import (
	"mathematica-forum/dto"
	"mathematica-forum/services"
	"mathematica-forum/templates"
	"net/http"
)

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	templates.Templates.ExecuteTemplate(w, "login.html", nil)
}

func ProcessLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	req := dto.LoginRequest{
		Identifiant: r.FormValue("identifiant"),
		Password:    r.FormValue("password"),
	}

	token, err := services.Login(req)
	if err != nil {
		data := map[string]string{
			"Message": err.Error(),
		}
		templates.Templates.ExecuteTemplate(w, "error.html", data)
		return
	}
	// Le fameux cookie
	cookie := &http.Cookie{
		Name:     "jwt_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
