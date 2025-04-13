package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type GetAllMonster struct {
	MonsterRetrieve []MonsterRetrieve `json:"data"`
}

func retrieveAllMonster(token []byte) {
	url := "https://api.artifactsmmo.com/monsters"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(token))

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var monster GetAllMonster
	json.Unmarshal([]byte(body), &monster)

	var dataIndex int
	for range monster.MonsterRetrieve {
		//fmt.Println()
		//fmt.Println()
		fmt.Println(strconv.Itoa(dataIndex) + ". Name: " + monster.MonsterRetrieve[dataIndex].Name)
		//fmt.Println("Code: " + monster.MonsterRetrieve[dataIndex].Code)
		//fmt.Println("Level: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Level))
		//fmt.Println("HP: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].HP))
		//fmt.Println("Attack Fire: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Attack_fire))
		//fmt.Println("Attack Earth: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Attack_earth))
		//fmt.Println("Attack Water: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Attack_water))
		//fmt.Println("Attack Air: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Attack_air))
		//fmt.Println("Attack Fire: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Attack_fire))
		//fmt.Println("Resistance Earth : " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Res_earth))
		//fmt.Println("Resistance Water: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Res_water))
		//fmt.Println("Resistance Air: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Res_air))
		//fmt.Println("Critical Strike: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Critical_strike))
		//fmt.Println("Minimum Gold: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Min_Gold))
		//fmt.Println("Maximum Gold: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Max_Gold))

		//if len(monster.MonsterRetrieve[dataIndex].Effects) > 0 {
		//fmt.Println("The associated effects are: ")

		//var index int
		//for range monster.MonsterRetrieve[dataIndex].Effects {
		//fmt.Println("Code: " + monster.MonsterRetrieve[dataIndex].Effects[index].Code)
		//fmt.Println("Value: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Effects[index].Value))
		//index++
		//}
		//}

		//if len(monster.MonsterRetrieve[dataIndex].Drops) > 0 {
		//fmt.Println("The associated random drops are: ")

		//var index int
		//for range monster.MonsterRetrieve[dataIndex].Drops {
		//fmt.Println("Code: " + monster.MonsterRetrieve[dataIndex].Drops[index].Code)
		//fmt.Println("Minimum quantity: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Drops[index].Min_Quantity))
		//fmt.Println("Maximum quantity: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Drops[index].Max_Quantity))
		//fmt.Println("Rate of drop: " + strconv.Itoa(monster.MonsterRetrieve[dataIndex].Drops[index].Rate))

		//index++
		//}
		//}
		dataIndex++
	}
}
