package main

import (
	"log"

	client "github.com/leonardyeoxl/careerjet-go/client"
)

func main() {
	searchParams := &client.SearchParams{
		Affid:          "",
		User_ip:        "111.222.234.123",
		User_agent:     "careerjet-api-client-v3.0.1",
		Locale_code:    "en_SG",
		Location:       "singapore",
		Keywords:       "python,java",
		Sort:           "relevance",
		Start_num:      "",
		Pagesize:       "",
		Page:           "",
		Contracttype:   "p",
		Contractperiod: "",
	}
	results, err := client.Search(searchParams)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(results)
}
