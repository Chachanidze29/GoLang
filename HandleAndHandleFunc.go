package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello world")
	})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/Home", Home)
	http.HandleFunc("/About", About)
	http.HandleFunc("/Contact", Contact)
	http.ListenAndServe(":8000", nil)
}

func Home(wr http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(wr, "home.gohtml", "Jimmy choo\n")
	fmt.Fprintln(wr, "Hello from Home")
}

func About(wr http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(wr, "about.gohtml", "Gosha rubchinsky\n")
	fmt.Fprintln(wr, "Hello from About")
}

func Contact(wr http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(wr, "contact.gohtml", "Junya Watanabe\n")
	fmt.Fprintln(wr, "Here you can contact us")
}
