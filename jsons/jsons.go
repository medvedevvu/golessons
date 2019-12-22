package main1

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1941, Color: false,
		Actors: []string{"Morgunov", "Nikulin"},
	},
	{Title: "Tree musketas", Year: 1980, Color: false,
		Actors: []string{"Xazanov", "Kostlolevskii"},
	},
	{Title: "New Afon", Year: 1984, Color: true,
		Actors: []string{"Xazanov", "Kostlolevskii"},
	},
}

func main() {
	//data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatal("Сбой маршалинга JSON: %s", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }

	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatal("Сбой демаршалинга JSON: %s", err)
	}
	fmt.Println(titles)
}
