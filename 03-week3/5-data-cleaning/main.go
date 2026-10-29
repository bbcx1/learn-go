package main

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"os"
)

type Review struct {
	Username       string `json:"username"`
	Score          int    `json:"score"`
	RestaurantName string `json:"restaurant_name"`
	Comment        string `json:"comment"`
}

func main() {
	path := "reviews.json"
	byteVal, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error read the file")
		return
	}

	var reviews []Review

	err = json.Unmarshal(byteVal, &reviews)
	if err != nil {
		fmt.Println("Failed unmarshal")
		return
	}

	for i := range reviews {
		// cleaning data
		// empty string -> N/A
		// empty int -> 0 (don't need to handle it because zero value of int is already 0)
		// -1 -> leave it as it is, there is no instruction on this.

		if reviews[i].Username == "" {
			reviews[i].Username = "N/A"
		}

		if reviews[i].RestaurantName == "" {
			reviews[i].RestaurantName = "N/A"
		}

		if reviews[i].Comment == "" {
			reviews[i].Comment = "N/A"
		}
	}

	newFile, err := os.Create("clean_reviews.json")
	if err != nil {
		fmt.Println("Error creating file")
	}
	defer newFile.Close()

	err = json.MarshalWrite(newFile, reviews, jsontext.WithIndent("  "))
	if err != nil {
		fmt.Println("Error write")
	}
}
