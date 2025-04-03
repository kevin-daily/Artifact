package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Bank struct {
	Bank BankDetails `json:"data"`
}

func getBankDetails(token []byte) {

	url := "https://api.artifactsmmo.com/my/bank"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println("Body:")
	fmt.Println(string(body))

	var bank Bank
	err := json.Unmarshal([]byte(body), &bank)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Unmarshaled data:")
	fmt.Println(bank)

}
