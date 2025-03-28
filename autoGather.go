package main

import (
	"time"
)

func autoGather(name string, turns int, token []byte) {
	for i := 0; i < turns; i++ {
		cooldown := gatherResources(name, token)

		time.Sleep(time.Duration(cooldown) * time.Second)
	}
}
