package main

import (
	"time"
)

func autoGather(name string, turns int, token []byte) {
	for i := 0; i < turns; i++ {
		cooldown := gatherResources(name, token)

		if cooldown == 0 {
			return
		}

		time.Sleep(time.Duration(cooldown) * time.Second)
	}
}
