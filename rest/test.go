package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"log"
)

// API response model

type buyer_rep_t struct {
	Tags				[]string				`json:"tags"`
}

type status_t struct {
	SiteStatus			string					`json:"site_status"`
}

type ratings_t struct {
	Negative			int					`json:"negative"`
	Neutral				int					`json:"neutral"`
	Positive			int					`json:"positive"`
}

type transactions_t struct {
	Cancelled			int					`json:"cancelled"`
	Completed			int					`json:"completed"`
	Period				string					`json:"period"`
	Ratings				ratings_t				`json:"ratings"`
	Total				int					`json:"total"`
}

type seller_reputation_t struct {
	LevelId				string					`json:"level_id"`
	PowerSellerStatus		string					`json:"power_seller_status"`
	Transactions   			transactions_t				`json:"transactions"`
}

type adress_t struct {
	City				string					`json:"city"`
	State				string					`json:"state"`
}

type MercadoLibreUser struct { 
	Id 				int					`json:"id"`
	Nickname			string 					`json:"nickname"`
	RegistrationDate		string 				`	json:"registration_date"`
	CountryId			string					`json:"country_id"`
	Adress				adress_t				`json:"adress"`
	UserType			string					`json:"user_type"`
	Logo				string					`json:"logo"`
	Points				int					`json:"points"`
	SiteId				string					`json:"site_id"`
	Permalink			string					`json:"permalink"`
	Tags				[]string 				`json:"tags"`
	SellerReputation  		seller_reputation_t			`json:"seller_reputation"`
	BuyerReputation			buyer_rep_t 	    			`json:"buyer_reputation"`
	Status 				status_t				`json:"status"`
}


func main() {

	url := "https://api.mercadolibre.com/users/1"

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err:= client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	user := MercadoLibreUser{}
	jsonresp := json.Unmarshal(body,&user)
	if jsonresp != nil {
		log.Fatal(jsonresp)
	}

	fmt.Println(user.Id)
}


