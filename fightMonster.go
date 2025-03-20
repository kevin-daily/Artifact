package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type FightData struct {
	Data Data
}

func main() {
	args := os.Args

	name := args[1]

	url := "https://api.artifactsmmo.com/my/" + name + "/action/fight"

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

	var fightData FightData

	json.Unmarshal([]byte(body), &fightData)

	fmt.Println("Result:")
	fmt.Println("You " + fightData.Data.Fight.Result + "!")

	fmt.Println()

	fmt.Println("You received " + strconv.Itoa(fightData.Data.Fight.XP) + " XP and " + strconv.Itoa(fightData.Data.Fight.Gold) + " gold")

	fmt.Println("The monster dropped ")
	if len(fightData.Data.Fight.Drops) == 0 {
		fmt.Print("nothing!")
	} else {
		fmt.Println(fightData.Data.Fight.Drops)
	}
	fmt.Println()
	fmt.Println("Fight logs:")
	fmt.Println(fightData.Data.Fight.Logs)

}
