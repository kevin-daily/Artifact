package artifact

import (
	"fmt"
	"os"
)

func main() {

	getStatus()

	var characterName, userInput string

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
		}
	}
}
