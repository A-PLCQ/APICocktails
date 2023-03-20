package packages

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Item struct {
	Name         string `json:"strDrink"`
	Instructions string `json:"strInstructions"`
	Image        string `json:"strDrinkThumb"`
}

type Items struct {
	Items []Item
}

func GetItems(request string) Items {
	timeClient := http.Client{
		Timeout: time.Second * 2,
	}
	url := "http://www.thecocktaildb.com/api/json/v1/1/" + request
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	res, _ := timeClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var data map[string][]Item
	err := json.Unmarshal(body, &data)
	fmt.Println(err)
	fmt.Println()

	items := Items{Items: data["drinks"]}

	return items
}
