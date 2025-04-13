package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetItem struct {
	ItemRetrieve ItemRetrieve `json:"data"`
}

func retrieveItem(code string, token []byte) {
	url := "https://api.artifactsmmo.com/items/" + code

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	switch res.StatusCode {
	case 404:
		fmt.Println("Item not found")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var item GetItem
	json.Unmarshal([]byte(body), &item)

	fmt.Println()
	fmt.Println()
	fmt.Println("Name: " + item.ItemRetrieve.Name)
	fmt.Println("Code: " + item.ItemRetrieve.Code)
	fmt.Println("Level: " + strconv.Itoa(item.ItemRetrieve.Level))
	fmt.Println("Type: " + item.ItemRetrieve.Type)
	fmt.Println("Subtype: " + item.ItemRetrieve.Subtype)
	fmt.Println("Description: " + item.ItemRetrieve.Description)
	fmt.Println("Tradeable: " + strconv.FormatBool(item.ItemRetrieve.Tradeable))

	if len(item.ItemRetrieve.Effects) > 0 {
		fmt.Println()
		fmt.Println("The associated effects are: ")

		var index int
		for range item.ItemRetrieve.Effects {
			fmt.Println("Code: " + item.ItemRetrieve.Effects[index].Code)
			fmt.Println("Value: " + strconv.Itoa(item.ItemRetrieve.Effects[index].Value))
			index++
		}
	}

	fmt.Println("Crafting skill: " + item.ItemRetrieve.Craft.Skill)
	fmt.Println("Crafting level: " + strconv.Itoa(item.ItemRetrieve.Craft.Level))
	fmt.Println("Quantity crafted: " + strconv.Itoa(item.ItemRetrieve.Craft.Quantity))

	if len(item.ItemRetrieve.Craft.Items) > 0 {
		fmt.Println()
		fmt.Println("Required items: ")

		var index int
		for range item.ItemRetrieve.Craft.Items {
			fmt.Println("Code: " + item.ItemRetrieve.Craft.Items[index].Code)
			fmt.Println("Quantity: " + strconv.Itoa(item.ItemRetrieve.Craft.Items[index].Quantity))
			index++
		}
	}

}
