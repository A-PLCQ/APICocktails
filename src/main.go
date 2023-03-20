package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"src/packages"
)

const port = "8082"
const url = "localhost"

type Mitems struct {
	packages.Items
}

var items Mitems

func Index(w http.ResponseWriter, r *http.Request) {
	// Liste des cocktails commençant par "ma"
	//items := packages.GetCocktailsByName("search.php?s=margarita")

	// Liste des cocktails contenant l'ingrédient "vodka"
	items := packages.GetItems("filter.php?i=Vodka")

	fmt.Println(items)

	template.Must(template.ParseFiles("templates/index.html")).Execute(w, items)
}

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	http.HandleFunc("/", Index)

	fmt.Println("(http://"+url+":"+port+") - Server started on port", port)
	http.ListenAndServe(":"+port, nil)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
