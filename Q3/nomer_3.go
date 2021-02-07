package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"refactory/models"
	"strings"
)

func main() {
	fmt.Println(BrownColour())
}

var path string = "api_2.json"

func ItemMeetingRoom() []models.Room {
	var room []models.Room

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &room)

	var search string = "meeting room"
	var arr []models.Room
	for _, key := range room {
		if strings.ToLower(key.Placements.Name) == search {
			arr = append(arr, key)
		}
	}
	return arr

}

func ElectronicsDevice() []models.Room {
	var room []models.Room

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &room)

	var search string = "electronic"
	var arr []models.Room
	for _, key := range room {
		if strings.ToLower(key.Type) == search {
			arr = append(arr, key)
		}
	}
	return arr
}

func Furniture() []models.Room {
	var room []models.Room

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &room)

	var search string = "furniture"
	var arr []models.Room
	for _, key := range room {
		if strings.ToLower(key.Type) == search {
			arr = append(arr, key)
		}
	}
	return arr
}

// func PurchasedJanuary() []models.Room {
// 	var room []models.Room

// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = json.Unmarshal(data, &room)

// 	purchasedTime := time.Parse()

// }

func BrownColour() []models.Room {
	var room []models.Room

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(data, &room)

	var search string = "brown"

	var arr []models.Room
	for _, key := range room {
		find := findBrown(search, key.Tags)
		if find {
			arr = append(arr, key)
		}
	}
	return arr
}

func findBrown(query string, data []string) bool {
	var find bool = false
	for i := range data {
		if strings.ToLower(data[i]) == query {
			find = true
		}
	}
	return find
}
