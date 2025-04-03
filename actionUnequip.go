package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Unequip struct {
	ActionUnequip ActionUnequip `json:"data"`
}

func unequipItem(name string, slot string, quantity int, token []byte) {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/unequip"

	type UnequipPayload struct {
		Slot     string `json:"slot"`
		Quantity int    `json:"quantity"`
	}

	payload := UnequipPayload{
		Slot:     slot,
		Quantity: quantity,
	}

	marshaledJSON, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", url, bytes.NewReader(marshaledJSON))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	switch res.StatusCode {
	case 478:
		fmt.Println("Item mising from inventory or insufficient quantity")
		return
	case 483:
		fmt.Println("Character does not have enough HP to remove this item")
		return
	case 491:
		fmt.Println("Slot is empty")
		return
	case 497:
		fmt.Println("Character inventory is full.")
		return
	case 499:
		fmt.Println("Character is in cooldown. Try again later")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var unequip Unequip
	json.Unmarshal([]byte(body), &unequip)

	fmt.Println()
	fmt.Println("You have unequipped " + unequip.ActionUnequip.Item.Name)
	fmt.Println()
	fmt.Println("Cooldown resets in " + strconv.Itoa(unequip.ActionUnequip.Cooldown.Remaining_Seconds))
}
