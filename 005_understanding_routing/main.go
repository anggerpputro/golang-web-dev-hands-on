package main

import (
	"fmt"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from hotdog!")
}

func cat(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello from hotcat!")
}

func main() {

	http.HandleFunc("/dog", dog)
	http.HandleFunc("/cat", cat)

	http.ListenAndServe(":8080", nil)

}
