package main

import (
	"html/template"
	"net/http"
)

const port = "8082"
const url = "localhost"

func Index(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("src/templates/index.html")).Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(url+":"+port, nil)
}
