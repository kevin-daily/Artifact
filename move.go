package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ActionMove struct {
	Move Move `json:"data"`
}

func MoveTo(name string, x string, y string, token []byte) {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/move"

	payload := strings.NewReader("{\n  \"x\": " + x + ",\n  \"y\": " + y + "\n}")

	req, _ := http.NewRequest("POST", url, payload)

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
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var actionMove ActionMove
	json.Unmarshal([]byte(body), &actionMove)

	fmt.Print("You are in a ")
	fmt.Println(actionMove.Move.Destination.Name)
	fmt.Println()

	locationType := actionMove.Move.Destination.Content.Type

	switch locationType {
	case "monster":
		fmt.Println("This location contains a " + actionMove.Move.Destination.Content.Code + " you can fight")
	case "resource":
		fmt.Println("This location contains a resource of type " + actionMove.Move.Destination.Content.Code)
	case "workshop":
		fmt.Println("This location has a workshop where you can learn " + actionMove.Move.Destination.Content.Code)
	case "npc":
		fmt.Println("This location has a " + actionMove.Move.Destination.Content.Code)
	case "bank":
		fmt.Println("This location has a " + actionMove.Move.Destination.Content.Code)
	case "grand_exchange":
		fmt.Println("This location has the Grand Exchange")
	case "tasks_master":
		fmt.Println("This location has the Task Master for " + actionMove.Move.Destination.Content.Code)
	case "":
		fmt.Println("There is nothing here")
	}
	fmt.Println()
	fmt.Println("Cooldown will reset in " + strconv.Itoa(actionMove.Move.Cooldown.Remaining_Seconds) + " seconds")
	fmt.Println()

}
