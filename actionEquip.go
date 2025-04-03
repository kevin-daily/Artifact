package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Equip struct {
	ActionEquip ActionUnequip `json:"data"`
}

func equipItem(name string, code string, slot string, quantity int, token []byte) {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/equip"

	type EquipPayload struct {
		Code     string `json:"code"`
		Slot     string `json:"slot"`
		Quantity int    `json:"quantity"`
	}

	payload := EquipPayload{
		Code:     code,
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
	case 484:
		fmt.Println("Cannot equip more than 100 utilities in same slot")
		return
	case 485:
		fmt.Println("Item already equipped")
		return
	case 491:
		fmt.Println("Another item already equipped in slot")
		return
	case 496:
		fmt.Println("Character level is insufficient to equip this item. Git gud")
	case 499:
		fmt.Println("Character is in cooldown. Try again later")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var equip Equip
	json.Unmarshal([]byte(body), &equip)

	fmt.Println()
	fmt.Println("You have equipped " + equip.ActionEquip.Item.Name)
	fmt.Println()
	fmt.Println("Cooldown resets in " + strconv.Itoa(equip.ActionEquip.Cooldown.Remaining_Seconds))
}
