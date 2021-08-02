package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
}

func foo(rw http.ResponseWriter, req *http.Request) {
	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		fmt.Println("file", f, "header", h, "error", err)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		s = string(bs)
	}
	rw.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(rw, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}
