package artifact

import "fmt"

func main() {

	getStatus()

	var characterName, userInput string

	fmt.Println()
	fmt.Println("Enter your character name: ")
	fmt.Scanln(&characterName)

	for {
		fmt.Println("What would you like to do?")
		fmt.Scanln(&userInput)

		switch userInput {
		case "rest":
			TakeRest(characterName)

		case "fight":
			startFight(characterName)
		}
	}
}
