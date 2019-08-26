package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type handler int

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// err := req.ParseForm()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// data := struct {
	// 	Method        string
	// 	URL           *url.URL
	// 	Header        map[string][]string
	// 	ContentLength int64
	// 	Submissions   map[string][]string
	// }{
	// 	Method:        req.Method,
	// 	URL:           req.URL,
	// 	Header:        req.Header,
	// 	ContentLength: req.ContentLength,
	// 	Submissions:   req.PostForm,
	// }

	// tpl.ExecuteTemplate(w, "index.gohtml", data)

	fmt.Fprintln(w, "Hello from handler!")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
