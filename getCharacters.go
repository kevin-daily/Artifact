package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Characters struct {
	Data Data
}

func main() {

	url := "https://api.artifactsmmo.com/my/characters"

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

	var characters Characters
	err = json.Unmarshal(body, &characters)

	if err != nil {
		fmt.Println("Failed to parse JSON: " + err.Error())
	}

	fmt.Println(characters)

}
