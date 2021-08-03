package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8000", nil)
}

func foo(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Method", req.Method)
	fn := ""
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fn = h.Filename
		if !(strings.HasSuffix(fn, "jpg") || strings.HasSuffix(fn, "jpeg") || strings.HasSuffix(fn, "png")) {
			fmt.Fprintln(rw, "Wrong Type")
			return
		}
		df, err := os.Create("./assets/" + fn)
		if err != nil {
			panic(err)
		}
		defer df.Close()
		_, err = io.Copy(df, f)
		if err != nil {
			panic(err)
		}
	}
	tpl.ExecuteTemplate(rw, "index.gohtml", fn)
}
