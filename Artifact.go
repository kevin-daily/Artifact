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

	// TODO: Add loop that checks the given character anme and prompt if not found.
	fmt.Println()
	fmt.Println("Enter your character name or \"CharacterList\" to get list of available characters")
	fmt.Scanln(&characterName)
	if characterName == "CharacterList" {
		getCharacterList()
		fmt.Println()
		fmt.Println("Enter your character name: ")
		fmt.Scanln(&characterName)
	}

	for {
		fmt.Println("What would you like to do? Input \"help\" to get list of commands.")
		fmt.Scanln(&userInput)

		switch userInput {
		case "help":
			fmt.Println()
			fmt.Println("rest - Recover HP by resting" +
				"\nfight - Fight a single monster" +
				"\nautofight - Fight monster a specified number of times. Auto rests after each fight" +
				"\nmove - Move to a specified location (Uses X Y coordinates)" +
				"\ngather - Gather a single resource" +
				"\nautogather - Gather a specified number of resources" +
				"\nequip - Equip specified item" +
				"\nunequip - Unequip specified item" +
				"\ncraft - Craft an item" +
				"\ninventory - List the items in your inventory")
			fmt.Println()

		case "rest":
			TakeRest(characterName, token)

		case "fight":
			startFight(characterName, token)

		case "move":
			fmt.Println("Where would you like to move to? (Either workshop / resource name or X Y coordinates)")
			fmt.Scanln(&x, &y)

			switch x {
			case "mining":
				x = "1"
				y = "5"
			case "woodcutting":
				x = "-2"
				y = "-3"
			case "cooking":
				x = "1"
				y = "1"
			case "weaponcrafting":
				x = "2"
				y = "1"
			case "gearcrafting":
				x = "3"
				y = "1"
			case "jewelrycrafting":
				x = "1"
				y = "3"
			case "alchemy":
				x = "2"
				y = "3"
			case "copper_rocks":
				x = "2"
				y = "0"
			case "ash_tree":
				x = "-1"
				y = "0"
			case "salmon":
				x = "-3"
				y = "-4"
			}

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

		case "deposit":
			fmt.Println("Which item would you like to deposit?")
			fmt.Scanln(&itemCode)

			fmt.Println("How many would you like to deposit?")
			fmt.Scanln(&quantity)

			quantityNum, _ := strconv.Atoi(quantity)

			DepositBank(characterName, itemCode, quantityNum, token)

		case "withdraw":
			fmt.Println("Which item would you like to withdraw?")
			fmt.Scanln(&itemCode)

			fmt.Println("How many would you like to withdraw?")
			fmt.Scanln(&quantity)

			quantityNum, _ := strconv.Atoi(quantity)

			WithdrawBank(characterName, itemCode, quantityNum, token)

		case "exit":
			os.Exit(1)
		}
	}
}
