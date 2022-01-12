package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const searchURL = "http://public.api.careerjet.net/search"

func fetch(searchParams *SearchParams, wg *sync.WaitGroup, c chan string) {
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("affid", searchParams.Affid)
	q.Add("user_ip", searchParams.User_ip)
	q.Add("user_agent", searchParams.User_agent)
	q.Add("locale_code", searchParams.Locale_code)
	q.Add("location", searchParams.Location)
	q.Add("keywords", searchParams.Keywords)
	q.Add("sort", searchParams.Sort)
	q.Add("start_num", searchParams.Start_num)
	q.Add("pagesize", searchParams.Pagesize)
	q.Add("page", searchParams.Page)
	q.Add("contracttype", searchParams.Contracttype)
	q.Add("contractperiod", searchParams.Contractperiod)
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}

	newReq, _ := http.NewRequest("GET", req.URL.String(), nil)
	newReq.Header.Add("User-Agent", "careerjet-api-client-v3.0.1")
	newReq.Header.Add("Referer", "http://public.api.careerjet.net")
	resp, err := client.Do(newReq)
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
