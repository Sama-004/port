package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

func LeftClickLogger() {
	leftClickCount := 0
	event := hook.Start()
	defer hook.End()
	leftClickChan := make(chan int)

	go func() {
		for ev := range event {
			if ev.Kind == hook.MouseDown {
				if ev.Button == hook.MouseMap["left"] {
					leftClickChan <- 1
				}
			}
		}
	}()

	for {
		select {
		case <-leftClickChan:
			leftClickCount++
			fmt.Printf("Left Click count: %d\n", leftClickCount)
		}
	}
}

func rightClickLogger() {
	rightClickCount := 0

	event := hook.Start()
	defer hook.End()

	rightClickChan := make(chan int)

	go func() {
		for ev := range event {
			if ev.Kind == hook.MouseDown {
				if ev.Button == hook.MouseMap["center"] {
					rightClickChan <- 1
				}
			}
		}
		if hook.AddEvent("mright") {
			rightClickChan <- 1
		}
	}()

	for {
		select {
		case <-rightClickChan:
			rightClickCount++
			fmt.Printf("Right Click count: %d\n", rightClickCount)
		}
	}
}

func keypress() {
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
			fmt.Printf("Key press count: %d\n", keyPress)
		}
	}
}

func main() {
	LeftClickLogger()
	rightClickLogger()
	keypress()
}
