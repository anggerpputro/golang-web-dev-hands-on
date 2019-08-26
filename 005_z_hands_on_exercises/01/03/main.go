package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./views/*.gohtml"))
}

func main() {

	http.Handle("/", http.HandlerFunc(home))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(w, "HOME SWEET HOME")
	tpl.ExecuteTemplate(w, "index.gohtml", "HOME SWEET HOME")
}

func dog(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(w, "bark bark")
	tpl.ExecuteTemplate(w, "index.gohtml", "bark bark")
}

func me(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(w, "Hi, My name is Angger. Nice to meet you!")
	tpl.ExecuteTemplate(w, "index.gohtml", "Hi, My name is Angger. Nice to meet you!")
}
