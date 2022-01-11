package main

type SearchParams struct {
	affid          string `json:"affid"`
	user_ip        string `json:"user_ip"`
	user_agent     string `json:"user_agent"`
	locale_code    string `json:"locale_code"`
	location       string `json:"location"`
	keywords       string `json:"keywords"`
	sort           string `json:"sort"`
	start_num      string `json:"start_num"`
	pagesize       string `json:"pagesize"`
	page           string `json:"page"`
	contracttype   string `json:"contracttype"`
	contractperiod string `json:"contractperiod"`
}
