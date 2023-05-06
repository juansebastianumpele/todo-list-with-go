package web

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/entity"
	"embed"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"text/template"
	"time"
)

type AuthWeb interface {
	Login(w http.ResponseWriter, r *http.Request)
	LoginProcess(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	RegisterProcess(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type authWeb struct {
	userClient client.UserClient
	embed      embed.FS
}

func NewAuthWeb(userClient client.UserClient, embed embed.FS) *authWeb {
	return &authWeb{userClient, embed}
}

func (a *authWeb) Login(w http.ResponseWriter, r *http.Request) {
	headerPath := path.Join("views", "general", "header.html")
	loginPath := path.Join("views", "auth", "login.html")
	tmpl, err := template.ParseFS(a.embed, loginPath, headerPath)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	// TODO: answer here
}

func (a *authWeb) LoginProcess(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	userId, status, err := a.userClient.Login(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if status == 200 {
		http.SetCookie(w, &http.Cookie{
			Name:   "user_id",
			Value:  fmt.Sprintf("%d", userId),
			Path:   "/",
			MaxAge: 31536000,
			Domain: "",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func (a *authWeb) Register(w http.ResponseWriter, r *http.Request) {
	headerPath := path.Join("views", "general", "header.html")
	loginPath := path.Join("views", "auth", "register.html")
	tmpl, err := template.ParseFS(a.embed, loginPath, headerPath)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	// TODO: answer here
}

func (a *authWeb) RegisterProcess(w http.ResponseWriter, r *http.Request) {
	fullname := r.FormValue("fullname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	userId, status, err := a.userClient.Register(fullname, email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if status == 200 {
		http.SetCookie(w, &http.Cookie{
			Name:   "user_id",
			Value:  fmt.Sprintf("%d", userId),
			Path:   "/",
			MaxAge: 31536000,
			Domain: "",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}
}

func (a *authWeb) Logout(w http.ResponseWriter, r *http.Request) {
	c := &http.Cookie{}
	c.Name = "user_id"
	c.Value = ""
	c.Expires = time.Now()
	http.SetCookie(w, c)

	http.Redirect(w, r, "/login", http.StatusFound)
	// TODO: answer here
}
