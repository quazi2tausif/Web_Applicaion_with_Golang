package StoreData

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var tpl *template.Template

type Dictionary struct {
	Word       string
	Definition string
}

func init() {
	tpl = template.Must(template.ParseGlob("WebPage/*.gohtml"))
	http.HandleFunc("/", index)
	http.HandleFunc("/saveData", store)
	http.HandleFunc("/done", done)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, "We are not sure what happned!", http.StatusNotFound)
	}
}

func store(w http.ResponseWriter, req *http.Request) {
	wd := req.FormValue("word")
	dfs := req.FormValue("def")
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, "Dictionary", wd, 0, nil)
	entity := Dictionary{
		Word:       wd,
		Definition: dfs,
	}
	_, err := datastore.Put(ctx, key, &entity)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	http.Redirect(w, req, "/done", http.StatusSeeOther)
}

func done(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "done.gohtml", nil)
	if err != nil {
		http.Error(w, "We are not sure what happned!", http.StatusNotFound)
	}
}
