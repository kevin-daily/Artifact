package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ActionRest struct {
	Rest Rest `json:"data"`
}

func TakeRest(name string, token []byte) int {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/rest"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	if res.StatusCode > 299 {
		fmt.Println("StatusCode: " + strconv.Itoa(res.StatusCode))
		fmt.Println("Status: " + res.Status)
	}

	if res.StatusCode == 499 {
		fmt.Println("Character is in cooldown. Try again later")
		return 0
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var actionRest ActionRest

	json.Unmarshal([]byte(body), &actionRest)

	fmt.Println("Result:")
	fmt.Println("You restored " + strconv.Itoa(actionRest.Rest.HP_Restored) + " HP")
	fmt.Println()
	fmt.Println("Cooldown resets in " + strconv.Itoa(actionRest.Rest.Cooldown.Remaining_Seconds) + " seconds")
	fmt.Println()

	return actionRest.Rest.Cooldown.Remaining_Seconds
}
