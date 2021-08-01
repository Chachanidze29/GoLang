package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", animeGirl)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8000", nil)
}

func animeGirl(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(rw, "I Love Anime Girls ", `<img src="/resources/anime.jpg">`)
}
