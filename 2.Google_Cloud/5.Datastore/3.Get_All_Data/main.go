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
	http.HandleFunc("/allData", done)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, "We are not sure what happned!", http.StatusNotFound)
	}
}

func done(w http.ResponseWriter, req *http.Request) {
	ctx := appengine.NewContext(req)
	q := datastore.NewQuery("Dictionary")
	var entity []Dictionary
	_, err := q.GetAll(ctx, &entity)
	if err != nil {
		http.Error(w, "someting isn't right!", http.StatusNotFound)
	}
	err = tpl.ExecuteTemplate(w, "done.gohtml", entity)
	if err != nil {
		http.Error(w, "We are not sure what happned!", http.StatusNotFound)
	}
}
