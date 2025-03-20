package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type restData struct {
	Data Data
}

func main() {
	args := os.Args

	name := args[1]

	url := "https://api.artifactsmmo.com/my/" + name + "/action/rest"

	req, _ := http.NewRequest("POST", url, nil)

	token, err := os.ReadFile("api.txt")

	if err != nil {
		fmt.Println("Failed to get token with following error: " + err.Error())
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var restData restData

	json.Unmarshal([]byte(body), &restData)

	fmt.Println("Result:")
	fmt.Println("You restored " + strconv.Itoa(restData.Data.HP_Restored) + " HP")

	fmt.Println("Cooldown is now " + strconv.Itoa(restData.Data.Cooldown.Remaining_Seconds))
}
