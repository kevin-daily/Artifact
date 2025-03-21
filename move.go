package artifact

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type ActionMove struct {
	Move Move `json:"data"`
}

func MoveTo() {

	args := os.Args

	name := args[1]

	url := "https://api.artifactsmmo.com/my/" + name + "/action/move"

	payload := strings.NewReader("{\n  \"x\": " + args[2] + ",\n  \"y\": " + args[3] + "\n}")

	req, _ := http.NewRequest("POST", url, payload)

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

	var actionMove ActionMove

	json.Unmarshal([]byte(body), &actionMove)

	fmt.Println("Destination:")
	fmt.Println(actionMove.Move.Destination)

	fmt.Println("Cooldown:")
	fmt.Println(actionMove.Move.Cooldown)

	fmt.Println("Character:")
	fmt.Println(actionMove.Move.Character)

}
