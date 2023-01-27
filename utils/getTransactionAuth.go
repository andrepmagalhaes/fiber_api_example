package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TransactionAuth struct {

	Auth bool `json: "authorization"`

}

func GetTransactionAuth() (bool, error){
	resp, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")

	if err != nil {
		log.Fatalln(err)
		return false, fmt.Errorf("Error getting transaction authorization")
	}

	var transactionAuth TransactionAuth

	err = json.NewDecoder(resp.Body).Decode(&transactionAuth)

	if err != nil {
		log.Fatalln(err)
		return false, fmt.Errorf("Error getting transaction authorization")
	}

	if (transactionAuth.Auth == true) {
		return true, nil
	}

	return false, nil
}