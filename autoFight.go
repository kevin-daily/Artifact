package main

import (
	"time"
)

func autoFight(name string, turns int, token []byte) {
	for i := 0; i < turns; i++ {
		cooldown := startFight(name, token)

		time.Sleep(time.Duration(cooldown) * time.Second)

		cooldown = TakeRest(name, token)

		time.Sleep(time.Duration(cooldown) * time.Second)
	}
}
