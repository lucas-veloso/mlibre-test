package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
)

// API response model

func getApiResponse() (MercadoLibreUser, error)  {

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

	return user, err
}

func main() {

}