package main

import (
	"html/template"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", "Avto")
	if err != nil {
		panic(err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Welcome)
	http.ListenAndServe(":8000", mux)
}
