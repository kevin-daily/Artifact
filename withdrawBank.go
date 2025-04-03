package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type BankWithdrawal struct {
	DWBank DWBank `json:"data"`
}

func WithdrawBank(name string, code string, quantity int, token []byte) {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/bank/withdraw"

	payload := BankPayload{
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

	var bank BankDeposit
	json.Unmarshal([]byte(body), &bank)

	fmt.Println()
	fmt.Println("Items in bank:")

	var index int
	for range bank.DWBank.Bank {
		fmt.Println("Item: " + bank.DWBank.Bank[index].Code + " Quantity: " + strconv.Itoa(bank.DWBank.Bank[index].Quantity))
		index++
	}
	fmt.Println()

	fmt.Println("Cooldown will reset in " + strconv.Itoa(bank.DWBank.Cooldown.Remaining_Seconds))

}
