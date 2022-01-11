package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const searchURL = "http://public.api.careerjet.net/search"

func fetch(searchParams *SearchParams, wg *sync.WaitGroup, c chan string) {
	resp, err := http.Get(searchURL)
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the body to type string
	sb := string(body)

	wg.Done()
	c <- sb
}

func Search(searchParams *SearchParams) (map[string]interface{}, error) {
	var wg sync.WaitGroup
	results := map[string]interface{}{}
	searchResultsChannel := make(chan string, 1)

	wg.Add(1)
	go fetch(searchParams, &wg, searchResultsChannel)
	wg.Wait()

	body := <-searchResultsChannel
	json.Unmarshal([]byte(body), &results)

	return results, nil
}
