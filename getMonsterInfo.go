package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetMonster struct {
	MonsterRetrieve MonsterRetrieve `json:"data"`
}

func retrieveMonster(code string, token []byte) {
	url := "https://api.artifactsmmo.com/monsters/" + code

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	switch res.StatusCode {
	case 404:
		fmt.Println("Monster not found")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var monster GetMonster
	json.Unmarshal([]byte(body), &monster)

	fmt.Println("Name: " + monster.MonsterRetrieve.Name)
	fmt.Println("Code: " + monster.MonsterRetrieve.Code)
	fmt.Println("Level: " + strconv.Itoa(monster.MonsterRetrieve.Level))
	fmt.Println("HP: " + strconv.Itoa(monster.MonsterRetrieve.HP))
	fmt.Println("Attack Fire: " + strconv.Itoa(monster.MonsterRetrieve.Attack_fire))
	fmt.Println("Attack Earth: " + strconv.Itoa(monster.MonsterRetrieve.Attack_earth))
	fmt.Println("Attack Water: " + strconv.Itoa(monster.MonsterRetrieve.Attack_water))
	fmt.Println("Attack Air: " + strconv.Itoa(monster.MonsterRetrieve.Attack_air))
	fmt.Println("Attack Fire: " + strconv.Itoa(monster.MonsterRetrieve.Attack_fire))
	fmt.Println("Resistance Earth : " + strconv.Itoa(monster.MonsterRetrieve.Res_earth))
	fmt.Println("Resistance Water: " + strconv.Itoa(monster.MonsterRetrieve.Res_water))
	fmt.Println("Resistance Air: " + strconv.Itoa(monster.MonsterRetrieve.Res_air))
	fmt.Println("Critical Strike: " + strconv.Itoa(monster.MonsterRetrieve.Critical_strike))
	fmt.Println("Minimum Gold: " + strconv.Itoa(monster.MonsterRetrieve.Min_Gold))
	fmt.Println("Maximum Gold: " + strconv.Itoa(monster.MonsterRetrieve.Max_Gold))

	fmt.Println("The associated effects are: ")
	var index int
	for range monster.MonsterRetrieve.Effects {
		fmt.Println("Code: " + monster.MonsterRetrieve.Effects[index].Code)
		fmt.Println("Value: " + strconv.Itoa(monster.MonsterRetrieve.Effects[index].Value))
		index++
	}

	fmt.Println("The associated random drops are: ")
	index = 0
	for range monster.MonsterRetrieve.Drops {
		fmt.Println("Code: " + monster.MonsterRetrieve.Drops[index].Code)
		fmt.Println("Minimum quantity: " + strconv.Itoa(monster.MonsterRetrieve.Drops[index].Min_Quantity))
		fmt.Println("Maximum quantity: " + strconv.Itoa(monster.MonsterRetrieve.Drops[index].Max_Quantity))
		fmt.Println("Rate of drop: " + strconv.Itoa(monster.MonsterRetrieve.Drops[index].Rate))

		index++
	}
}
