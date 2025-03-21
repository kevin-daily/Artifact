package artifact

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Status struct {
	StatusData StatusData `json:"data"`
}

func getStatus() {
	url := "https://api.artifactsmmo.com/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println("Body:")
	fmt.Println(string(body))

	var status Status

	err := json.Unmarshal(body, &status)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Unmarshaled data:")
	fmt.Println(status)

}
