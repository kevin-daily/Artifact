package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Bank struct {
	Data Data
}

func main() {

	url := "https://api.artifactsmmo.com/my/bank"

	req, _ := http.NewRequest("GET", url, nil)

	token, err := os.ReadFile("api.txt")

	if err != nil {
		fmt.Println("Failed to get token with following error: " + err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println("Body:")
	fmt.Println(string(body))

	var bank Bank

	err = json.Unmarshal([]byte(body), &bank)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Unmarshaled data:")
	fmt.Println(bank)

}
