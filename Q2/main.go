package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"refactory/models"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello World")
	//jawab := phoneNumberFind()
	jawab3 := publishedBeforeAugust2019()
	fmt.Println(jawab3)

}

var path string = "api.json"

func phoneNumberFind() []models.Biodata {
	var biodata []models.Biodata
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &biodata)

	var arr []models.Biodata
	for _, key := range biodata {
		//fmt.Println(key.Profile.Phones)
		if len(key.Profile.Phones) == 0 {
			arr = append(arr, key)
		}
	}
	return arr
}

func findArticles() []models.Biodata {
	var biodata []models.Biodata
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &biodata)

	var arr []models.Biodata
	for _, key := range biodata {
		if len(key.Articles) == 0 {
			arr = append(arr, key)
		}
	}
	return arr
}

func findAnis() []models.Biodata {
	var biodata []models.Biodata
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &biodata)

	var arr []models.Biodata
	var search string = "annis"
	for _, key := range biodata {
		if strings.Contains(strings.ToLower(key.Profile.Fullname), search) {
			arr = append(arr, key)
		}
	}
	return arr
}

func findArticles2000() []models.Biodata {
	var biodata []models.Biodata
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &biodata)

	var arr []models.Biodata
	var layoutFormat string = "2006-01-02T15:04:05"
	for _, key := range biodata {
		if len(key.Articles) != 0 {
			for j := range key.Articles {
				year, _ := time.Parse(layoutFormat, key.Articles[j].PublishedAt)
				//fmt.Println(year)
				if year.Year() == 2000 {
					fmt.Println(key)
					arr = append(arr, key)
				}
			}

		}

	}
	return arr
}

func born1986() []models.Biodata {
	var biodata []models.Biodata
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &biodata)

	var arr []models.Biodata
	var layoutFormat string = "2006-01-02"
	for _, key := range biodata {
		year, _ := time.Parse(layoutFormat, key.Profile.Birthday)
		if year.Year() == 1986 {
			arr = append(arr, key)
		}

	}
	return arr
}

func tipsArticle() []models.Article {
	var biodata []models.Biodata
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &biodata)

	var arr []models.Article
	var search string = "tips"
	for _, key := range biodata {
		if len(key.Articles) != 0 {
			for j := range key.Articles {
				if strings.Contains(strings.ToLower(key.Articles[j].Title), search) {
					arr = append(arr, key.Articles[j])
				}
			}
		}
	}
	return arr
}

func publishedBeforeAugust2019() []models.Article {
	var biodata []models.Biodata
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &biodata)

	var arr []models.Article
	var layoutFormat string = "2006-01-02T15:04:05"
	beforeTime, _ := time.Parse(layoutFormat, "2019-08-01")
	for _, key := range biodata {
		if len(key.Articles) != 0 {
			for j := range key.Articles {
				year, _ := time.Parse(layoutFormat, key.Articles[j].PublishedAt)
				// fmt.Println(year)
				if beforeTime.Before(year) {
					arr = append(arr, key.Articles[j])
				}
			}

		}

	}
	return arr
}
