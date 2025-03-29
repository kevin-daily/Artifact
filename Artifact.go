package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	getStatus()

	var characterName, itemCode, userInput, quantity, slot, turns, x, y string

	token, err := os.ReadFile("api.txt")

	if err != nil {
		fmt.Println("Failed to get token with following error: " + err.Error())
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Enter your character name: ")
	fmt.Scanln(&characterName)

	for {
		fmt.Println("What would you like to do?")
		fmt.Scanln(&userInput)

		switch userInput {
		case "rest":
			TakeRest(characterName, token)

		case "fight":
			startFight(characterName, token)

		case "move":
			fmt.Println("Where would you like to move to? (x,y)")
			fmt.Scanln(&x, &y)

			MoveTo(characterName, x, y, token)

		case "gather":
			gatherResources(characterName, token)

		case "unequip":
			fmt.Println("What slot would you like to unequip?" +
				"\n(weapon, shield, helmet, body_armor, leg_armor, boots, ring1, ring2, amulet," +
				"artifact1, artifact2, artifact3, utility1, utility2, bag, rune)")
			fmt.Scanln(&slot)

			if slot == "utility1" || slot == "utility2" {
				fmt.Println("How many would you like to unequip?")
				fmt.Scanln(&quantity)
			} else {
				quantity = "1"
			}

			quantityNum, _ := strconv.Atoi(quantity)

			unequipItem(characterName, slot, quantityNum, token)

		case "equip":
			fmt.Println("What slot would you like to equip?" +
				"\n(weapon, shield, helmet, body_armor, leg_armor, boots, ring1, ring2, amulet," +
				"artifact1, artifact2, artifact3, utility1, utility2, bag, rune)")
			fmt.Scanln(&slot)

			fmt.Println("What item would you like to equip?")
			fmt.Scanln(&itemCode)

			if slot == "utility1" || slot == "utility2" {
				fmt.Println("How many would you like to unequip?")
				fmt.Scanln(&quantity)
			} else {
				quantity = "1"
			}

			quantityNum, _ := strconv.Atoi(quantity)

			equipItem(characterName, itemCode, slot, quantityNum, token)

		case "craft":
			fmt.Println("What item would you like to craft?")
			fmt.Scanln(&itemCode)

			fmt.Println("How many would you like to craft?")
			fmt.Scanln(&quantity)

			quantityNum, _ := strconv.Atoi(quantity)

			craftItem(characterName, itemCode, quantityNum, token)

		case "inventory":
			getCharInventory(characterName, token)

		case "autogather":
			fmt.Println("How many times would you like to auto gather?")
			fmt.Scanln(&turns)

			howMany, _ := strconv.Atoi(turns)
			autoGather(characterName, howMany, token)

		case "autoFight":
			fmt.Println("How many fights would you like to complete?")
			fmt.Scanln(&turns)

			howMany, _ := strconv.Atoi(turns)
			autoFight(characterName, howMany, token)

		case "bank":
			getBankDetails(token)

		case "exit":
			os.Exit(1)
		}
	}
}
