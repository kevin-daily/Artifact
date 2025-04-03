package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Status struct {
	StatusData StatusData `json:"data"`
}

func getStatus() {
	url := "https://api.artifactsmmo.com/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	switch res.StatusCode {
	case 200:
	case 499:
		fmt.Println("Character is in cooldown. Try again later")
		return
	case 598:
		fmt.Println("Bank is not at this location. Cannot perform this action here.")
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var status Status
	err := json.Unmarshal(body, &status)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(status.StatusData.Announcements[0])
	fmt.Println()
	fmt.Println("The server is " + status.StatusData.Status)
	fmt.Println("There are currently " + strconv.Itoa(status.StatusData.Characters_Online) + " players online")

}
