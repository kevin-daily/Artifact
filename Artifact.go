package artifact

import (
	"fmt"
	"os"
)

func main() {

	getStatus()

	var characterName, userInput, x, y string

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

		case "exit":
			os.Exit(1)
		}
	}
}
