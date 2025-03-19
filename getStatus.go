package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type returned struct {
	data data
}

type data struct {
	status            string
	version           string
	max_level         int64
	characters_online int64
	//server_time       string
	//last_wipe         string
	//next_wipe         string
}

func main() {

	url := "https://api.artifactsmmo.com/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println("Body:")
	fmt.Println(string(body))

	var returned returned

	err := json.Unmarshal(body, &returned)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Unmarshaled data:")
	fmt.Println(returned)

}
