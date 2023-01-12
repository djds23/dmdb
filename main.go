package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	apiKey := os.Getenv("API_KEY")
	movies := sendDiscover(apiKey)
	sendIndividualMovie(movies.Results[0], apiKey)
}

func sendDiscover(key string) Movies {
	// Discover (GET https://api.themoviedb.org/3/discover/movie?api_key=KEY&sort=popularity.desc&page=1)

	// Create client
	client := &http.Client{}

	// Create request
	requestURL := fmt.Sprintf("https://api.themoviedb.org/3/discover/movie?api_key=%s&sort=popularity.desc&page=1", key)
	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		fmt.Println("Failure constructing request: ", err)
	}

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure fetching request: ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)

	var m Movies
	// from json format
	marshallErr := json.Unmarshal(respBody, &m)

	if marshallErr != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Number of Movies: ", len(m.Results))
		fmt.Println("Page num: ", m.Page)

		first := m.Results[0]
		// Display Results
		fmt.Println("Best Movie : ", first.OriginalTitle)
		fmt.Println("Best MovieID ", first.Id)
	}
	return m
}

func sendIndividualMovie(movie Movie, key string) {
	// Individual Movie (GET https://api.themoviedb.org/3/movie/76600/credits?api_key=KEY)

	// Create client
	client := &http.Client{}

	// Create request
	requestURL := fmt.Sprintf("https://api.themoviedb.org/3/movie/%f/credits?api_key=%s", movie.Id, key)
	req, err := http.NewRequest("GET", requestURL, nil)

	if err != nil {
		fmt.Println("Failure constructing request: ", err)
	}

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := io.ReadAll(resp.Body)
	var c Credits
	marshallErr := json.Unmarshal(respBody, &c)

	if marshallErr != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Number of Actors: ", len(c.Cast))
		fmt.Println("ID: ", c.Id)

		first := c.Cast[0]
		fmt.Println("Star: ", first.Name)
	}
}
