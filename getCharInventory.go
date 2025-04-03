package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetCharInv struct {
	Character Character `json:"data"`
}

func getCharInventory(name string, token []byte) {

	url := "https://api.artifactsmmo.com/characters/" + name

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	switch res.StatusCode {
	case 200:
	case 499:
		fmt.Println("Character is in cooldown. Try again later")
		return
	case 598:
		fmt.Println("Bank is not at this location. Cannot perform this action here.")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var inventory GetCharInv
	err := json.Unmarshal(body, &inventory)

	if err != nil {
		fmt.Println("Failed to parse JSON: " + err.Error())
	}

	fmt.Println("Inventory:")

	fmt.Println("Max storage: " + strconv.Itoa(inventory.Character.Inventory_Max_Items))
	fmt.Println()

	var index int
	for range inventory.Character.Inventory {
		fmt.Println(inventory.Character.Inventory[index])
		index++
	}

	fmt.Println()
}
