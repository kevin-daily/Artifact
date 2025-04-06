package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetResource struct {
	ResourceRetrieve ResourceRetrieve `json:"data"`
}

func retrieveResource(code string, token []byte) {
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

	var resource GetResource
	json.Unmarshal([]byte(body), &resource)

	fmt.Println("The name is: " + resource.ResourceRetrieve.Name)
	fmt.Println("The code is: " + resource.ResourceRetrieve.Code)
	fmt.Println("The skill associated is: " + resource.ResourceRetrieve.Skill)
	fmt.Println("The level required is: " + strconv.Itoa(resource.ResourceRetrieve.Level))
	fmt.Println("The associated random drops are: ")

	var index int
	for range resource.ResourceRetrieve.Drops {
		fmt.Println("Code: " + resource.ResourceRetrieve.Drops[index].Code)
		fmt.Println("Minimum quantity: " + strconv.Itoa(resource.ResourceRetrieve.Drops[index].Min_Quantity))
		fmt.Println("Maximum quantity: " + strconv.Itoa(resource.ResourceRetrieve.Drops[index].Max_Quantity))
		fmt.Println("Rate of drop: " + strconv.Itoa(resource.ResourceRetrieve.Drops[index].Rate))

		index++
	}

}
