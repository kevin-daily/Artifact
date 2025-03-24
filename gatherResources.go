package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Gather struct {
	ActionGathering ActionGathering `json:"data"`
}

func gatherResources(name string, token []byte) {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/move"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var gather Gather
	json.Unmarshal([]byte(body), &gather)

	index := 0
	for index < len(gather.ActionGathering.Details.Items) {
		fmt.Println()
		fmt.Println("You got " + strconv.Itoa(gather.ActionGathering.Details.Items[index].Quantity) +
			" of " + gather.ActionGathering.Details.Items[index].Code)
	}

	fmt.Println()
	fmt.Println("You got " + strconv.Itoa(gather.ActionGathering.Details.XP) + " XP")
	fmt.Println()
	fmt.Println("Cooldown will reset in " + strconv.Itoa(gather.ActionGathering.Cooldown.Remaining_Seconds) + " seconds")
}
