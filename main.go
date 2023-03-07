package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = "8080"
const url = "localhost"

func Index(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("src/templates/index.html")).Execute(w, nil)
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Println("(http://"+url+":"+port+") - Server started on port", port)
	http.ListenAndServe(url+":"+port, nil)
}
