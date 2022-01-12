package client

type SearchParams struct {
	Affid          string `json:"affid"`
	User_ip        string `json:"user_ip"`
	User_agent     string `json:"user_agent"`
	Locale_code    string `json:"locale_code"`
	Location       string `json:"location"`
	Keywords       string `json:"keywords"`
	Sort           string `json:"sort"`
	Start_num      string `json:"start_num"`
	Pagesize       string `json:"pagesize"`
	Page           string `json:"page"`
	Contracttype   string `json:"contracttype"`
	Contractperiod string `json:"contractperiod"`
}
