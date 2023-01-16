package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

type Tour struct{
	Name, Price string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func contentFromWeb( url string) string {
	response, err := http.Get(url)
	checkError(err)
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	return string(bytes)
}


func toursFromJson(content string) [] Tour{
	tours := make([]Tour,0, 100)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)
	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}


func main() {
	url := "http://services.explorecalifornia.org/json/tours.php"
	content := contentFromWeb(url)
	//fmt.Println(content)
	tours := toursFromJson(content)
	//fmt.Println(tours)

	for _, tour := range tours {
		price, _, _ := big.ParseFloat(tour.Price, 10, 2, big.ToZero)
		fmt.Printf("%v ($%.2f)\n", tour.Name, price)
	}
}