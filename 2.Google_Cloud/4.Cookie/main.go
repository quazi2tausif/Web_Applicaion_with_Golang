package CookiesFoundSample

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

var un string

func init() {
	tpl = template.Must(template.ParseGlob("WebPage/*.gohtml"))
	http.HandleFunc("/", index)
	http.HandleFunc("/data", data)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, "We are not sure what happened!", http.StatusNotFound)
	}
}

func data(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(un)
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	err = tpl.ExecuteTemplate(w, "data.gohtml", c)
	if err != nil {
		http.Error(w, "We are not sure what happened!", http.StatusNotFound)
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	un = req.FormValue("uname")
	up := req.FormValue("upass")
	if up == un {
		http.SetCookie(w, &http.Cookie{
			Name:  un,
			Value: "Cookie Is Activated",
		})
		http.Redirect(w, req, "/data", http.StatusSeeOther)
	}
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie(un)
	if err != nil {
		http.Redirect(w, req, "/data", http.StatusSeeOther)
		return
	}
	c.MaxAge = -1 // delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
