package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func main() {
	keyPress := 0
	event := hook.Start()
	defer hook.End()
	keyPressChan := make(chan int)
	go func() {
		for ev := range event {
			if ev.Kind == hook.KeyDown {
				keyPressChan <- 1
			}
		}
	}()
	for {
		select {
		case <-keyPressChan:
			keyPress++
			fmt.Printf("Total keypresses: %d\n", keyPress)
		}
	}
}
