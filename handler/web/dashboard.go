package web

import (
	"a21hc3NpZ25tZW50/client"
	"embed"
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

	// TODO: answer here
	dash := path.Join("views", "main", "dashboard.html")
	header := path.Join("views", "general", "header.html")

	var tmpl = template.Must(template.ParseFS(d.embed, dash, header)).Funcs(funcMap)
	// tmpl, err := template.New("dashboard.html").Funcs(funcMap).ParseFS(d.embed, dash, header)
	// if err != nil {
	// 	log.Println("error parse template: ", err.Error())
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	exc := tmpl.Execute(w, dataTemplate)
	if exc != nil {		
		http.Error(w, exc.Error(), http.StatusInternalServerError)
		return
	}
	// end of answer
}
