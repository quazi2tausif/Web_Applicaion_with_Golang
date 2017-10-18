package PassingData

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
	http.HandleFunc("/", index)
}

type user struct {
	UserName     string
	UserPassword string
}

func index(w http.ResponseWriter, req *http.Request) {
	var u user
	if req.Method == "POST" {
		un := req.FormValue("uname")
		up := req.FormValue("upass")
		u = user{
			UserName:     un,
			UserPassword: up,
		}
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}
