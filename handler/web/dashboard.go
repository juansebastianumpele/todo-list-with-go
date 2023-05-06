package web

import (
	"a21hc3NpZ25tZW50/client"
	"embed"
	"log"
	"net/http"
	"path"
	"text/template"
)

type DashboardWeb interface {
	Dashboard(w http.ResponseWriter, r *http.Request)
}

type dashboardWeb struct {
	categoryClient client.CategoryClient
	embed          embed.FS
}

func NewDashboardWeb(catClient client.CategoryClient, embed embed.FS) *dashboardWeb {
	return &dashboardWeb{catClient, embed}
}

func (d *dashboardWeb) Dashboard(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	categories, err := d.categoryClient.GetCategories(userId.(string))
	if err != nil {
		log.Println("error get cat: ", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dataTemplate = map[string]interface{}{
		"categories": categories,
	}

	var funcMap = template.FuncMap{
		"categoryInc": func(catId int) int {
			return catId + 1
		},
		"categoryDec": func(catId int) int {
			return catId - 1
		},
	}

	// ignore this
	_ = dataTemplate
	_ = funcMap
	//
	headerPath := path.Join("views", "general", "header.html")
	loginPath := path.Join("views", "main", "dashboard.html")
	// tmpl, err := template.ParseFS(d.embed, loginPath, headerPath)
	var tmpl = template.Must(template.New("dashboard.html").
		Funcs(funcMap).
		ParseFS(d.embed, loginPath, headerPath))
	if err := tmpl.Execute(w, dataTemplate); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// TODO: answer here
}
