package main

import (
	"log"

	client "github.com/leonardyeoxl/careerjet-go/client"
)

func main() {
	searchParams := &client.SearchParams{
		Affid:          "a43fb0f5c7454495183575ee2d75dcf3",
		User_ip:        "111.222.234.123",
		User_agent:     "careerjet-api-client-v3.0.1",
		Locale_code:    "en_SG",
		Location:       "singapore",
		Keywords:       "python",
		Sort:           "",
		Start_num:      "",
		Pagesize:       "",
		Page:           "",
		Contracttype:   "",
		Contractperiod: "",
	}
	results, err := client.Search(searchParams)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(results)
}
