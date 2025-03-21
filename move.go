package artifact

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ActionMove struct {
	Move Move `json:"data"`
}

func MoveTo(name string, x string, y string, token []byte) {
	url := "https://api.artifactsmmo.com/my/" + name + "/action/move"

	payload := strings.NewReader("{\n  \"x\": " + x + ",\n  \"y\": " + y + "\n}")

	req, _ := http.NewRequest("POST", url, payload)

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
