package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type FightMonster struct {
	ActionFight ActionFight `json:"data"`
}

func startFight(name string, token []byte) int {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/fight"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	switch res.StatusCode {
	case 200:
	case 499:
		fmt.Println("Character is in cooldown. Try again later")
		return 0
	case 598:
		fmt.Println("Bank is not at this location. Cannot perform this action here.")
		return 0
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var fightMonster FightMonster

	json.Unmarshal([]byte(body), &fightMonster)

	fmt.Println("Result:")
	fmt.Println("You " + fightMonster.ActionFight.Fight.Result + "!")

	fmt.Println()

	fmt.Println("You received " + strconv.Itoa(fightMonster.ActionFight.Fight.XP) + " XP and " + strconv.Itoa(fightMonster.ActionFight.Fight.Gold) + " gold")

	fmt.Print("The monster dropped ")
	if len(fightMonster.ActionFight.Fight.Drops) == 0 {
		fmt.Println("nothing!")
	} else {
		fmt.Println(fightMonster.ActionFight.Fight.Drops)
	}
	fmt.Println()
	fmt.Println("Fight logs:")
	fmt.Println(fightMonster.ActionFight.Fight.Logs)
	fmt.Println()
	fmt.Println("Cooldown will reset in " + strconv.Itoa(fightMonster.ActionFight.Cooldown.Remaining_Seconds) + " seconds")
	fmt.Println()

	return fightMonster.ActionFight.Cooldown.Remaining_Seconds
}
