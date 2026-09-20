package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Color    bool     `json:"color,omitempty"`
	Actors   []string `json:"actors"`
	cacheKey string
}

func main() {
	movies := []Movie{
		{
			Title:    "The Matrix",
			Year:     1999,
			Color:    true,
			Actors:   []string{"Keanu Reeves", "Carrie-Anne Moss"},
			cacheKey: "matrix-1999",
		},
	}

	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	var decoded []Movie
	if err := json.Unmarshal(data, &decoded); err != nil {
		panic(err)
	}
	fmt.Printf("decoded: %q (%d), actors=%d\n",
		decoded[0].Title, decoded[0].Year, len(decoded[0].Actors))
}
