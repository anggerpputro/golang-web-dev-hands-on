package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)

}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "HOME SWEET HOME")
}

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "bark bark")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hi, My name is Angger. Nice to meet you!")
}
