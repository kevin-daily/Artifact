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

	var unequip Unequip
	json.Unmarshal([]byte(body), &unequip)

	fmt.Println()
	fmt.Println("You have unequipped " + unequip.ActionUnequip.Item.Name)
	fmt.Println()
	fmt.Println("Cooldown resets in " + strconv.Itoa(unequip.ActionUnequip.Cooldown.Remaining_Seconds))
}
