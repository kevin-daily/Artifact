package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Crafting struct {
	ActionCraft ActionCraft `json:"data"`
}

func craftItem(name string, code string, quantity int, token []byte) {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/crafting"

	type CraftingPayload struct {
		Code     string `json:"code"`
		Quantity int    `json:"quantity"`
	}

	payload := CraftingPayload{
		Code:     code,
		Quantity: quantity,
	}

	marshaledJSON, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(marshaledJSON))

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

	var crafting Crafting
	json.Unmarshal([]byte(body), &crafting)

	fmt.Println()
	for index := range crafting.ActionCraft.Details.Items {
		fmt.Println()
		fmt.Println("You got " + strconv.Itoa(crafting.ActionCraft.Details.Items[index].Quantity) +
			" of " + crafting.ActionCraft.Details.Items[index].Code)
	}

	fmt.Println("You got " + strconv.Itoa(crafting.ActionCraft.Details.XP) + " XP")
	fmt.Println()
	fmt.Println("Cooldown resets in " + strconv.Itoa(crafting.ActionCraft.Cooldown.Remaining_Seconds))
}
