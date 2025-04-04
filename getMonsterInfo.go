package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GetMonster struct {
	MonsterRetrieve MonsterRetrieve `json:"data"`
}

func retrieveMonster(code string, token []byte) {
	url := "https://api.artifactsmmo.com/resources/" + code

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	switch res.StatusCode {
	case 404:
		fmt.Println("Resource not found")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var resource ResourceRetrieve
	json.Unmarshal([]byte(body), &resource)

}
