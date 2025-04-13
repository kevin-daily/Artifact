package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetItems struct {
	ItemRetrieve []ItemRetrieve `json:"data"`
}

func retrieveAllItems(token []byte) {
	url := "https://api.artifactsmmo.com/items"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var items GetItems
	json.Unmarshal([]byte(body), &items)

	var dataIndex int
	for range items.ItemRetrieve {
		//fmt.Println()
		//fmt.Println()
		fmt.Println(strconv.Itoa(dataIndex) + ". Name: " + items.ItemRetrieve[dataIndex].Name)
		//fmt.Println("Code: " + items.ItemRetrieve[dataIndex].Code)
		//fmt.Println("Level: " + strconv.Itoa(items.ItemRetrieve[dataIndex].Level))
		//fmt.Println("Type: " + items.ItemRetrieve[dataIndex].Type)
		//fmt.Println("Subtype: " + items.ItemRetrieve[dataIndex].Subtype)
		//fmt.Println("Description: " + items.ItemRetrieve[dataIndex].Description)
		//fmt.Println("Tradeable: " + strconv.FormatBool(items.ItemRetrieve[dataIndex].Tradeable))

		//if len(items.ItemRetrieve[dataIndex].Effects) > 0 {
		//fmt.Println()
		//fmt.Println("The associated effects are: ")

		//var index int
		//for range items.ItemRetrieve[dataIndex].Effects {
		//fmt.Println("Code: " + items.ItemRetrieve[dataIndex].Effects[index].Code)
		//fmt.Println("Value: " + strconv.Itoa(items.ItemRetrieve[dataIndex].Effects[index].Value))
		//index++
		//}
		//}

		//fmt.Println("Crafting skill: " + items.ItemRetrieve[dataIndex].Craft.Skill)
		//fmt.Println("Crafting level: " + strconv.Itoa(items.ItemRetrieve[dataIndex].Craft.Level))
		//fmt.Println("Quantity crafted: " + strconv.Itoa(items.ItemRetrieve[dataIndex].Craft.Quantity))

		//if len(items.ItemRetrieve[dataIndex].Craft.Items) > 0 {
		//fmt.Println()
		//fmt.Println("Required items: ")

		//var index int
		//for range items.ItemRetrieve[dataIndex].Craft.Items {
		//fmt.Println("Code: " + items.ItemRetrieve[dataIndex].Craft.Items[index].Code)
		//fmt.Println("Quantity: " + strconv.Itoa(items.ItemRetrieve[dataIndex].Craft.Items[index].Quantity))
		//index++
		//}
		//}
		dataIndex++
	}
}
