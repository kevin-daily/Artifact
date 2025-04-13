package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetAllResources struct {
	Resource []ResourceRetrieve `json:"data"`
}

func retrieveAllResources(token []byte) {
	url := "https://api.artifactsmmo.com/resources"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var items GetAllResources
	json.Unmarshal([]byte(body), &items)

	var dataIndex int
	for range items.Resource {
		//fmt.Println()
		//fmt.Println()
		fmt.Println(strconv.Itoa(dataIndex) + ". The name is: " + items.Resource[dataIndex].Name)
		//fmt.Println("The code is: " + items.Resource[dataIndex].Code)
		//fmt.Println("The skill associated is: " + items.Resource[dataIndex].Skill)
		//fmt.Println("The level required is: " + strconv.Itoa(items.Resource[dataIndex].Level))
		//fmt.Println()

		//if len(items.Resource[dataIndex].Drops) > 0 {
		//fmt.Println("The associated random drops are: ")

		//var index int
		//for range items.Resource[dataIndex].Drops {
		//fmt.Println("Code: " + items.Resource[dataIndex].Drops[index].Code)
		//fmt.Println("Minimum quantity: " + strconv.Itoa(items.Resource[dataIndex].Drops[index].Min_Quantity))
		//fmt.Println("Maximum quantity: " + strconv.Itoa(items.Resource[dataIndex].Drops[index].Max_Quantity))
		//fmt.Println("Rate of drop: " + strconv.Itoa(items.Resource[dataIndex].Drops[index].Rate))

		//index++
		//}
		//}
		dataIndex++
	}
}
