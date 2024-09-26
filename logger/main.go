package main

import (
	"fmt"
	"log"
	"time"

	hook "github.com/robotn/gohook"
)

var leftClickCount, rightClickCount, keyPress int

func logger(event chan hook.Event) {
	leftClickChan := make(chan int)
	rightClickChan := make(chan int)
	keyPressChan := make(chan int)

	go func() {
		for ev := range event {
			if ev.Kind == hook.MouseDown {
				if ev.Button == hook.MouseMap["left"] {
					leftClickChan <- 1
				}
				// something is wrong with this package it registers center as right click and does not work on right click
				// TODO: try to move away from this package later
				if ev.Button == hook.MouseMap["center"] {
					rightClickChan <- 1
				}
			}
			if ev.Kind == hook.KeyDown {
				keyPressChan <- 1
			}
		}
	}()

	for {
		select {
		case <-leftClickChan:
			leftClickCount++
			fmt.Printf("Left Click count: %d\n", leftClickCount)
		case <-rightClickChan:
			rightClickCount++
			fmt.Printf("Right click count:%d\n", rightClickCount)
		case <-keyPressChan:
			keyPress++
			fmt.Printf("Key press count: %d\n", keyPress)
		}
	}
}

func updateDb() {
	//also need time here. Time should be date and time.
	currentTime := time.Now().Format("2006-01-02 15:04")
	log.Printf("Updating database with LeftClick: %d Rightclick: %d Keypress: %d Time: %s\n", leftClickCount, rightClickCount, keyPress, currentTime)
	// db logic goes here
	// what happens when failed to update the db??? store the count with time in array and try to write later??
	//after db write reset the counts
	leftClickCount = 0
	rightClickCount = 0
	keyPress = 0
	fmt.Printf("Reset counts Leftclick: %d Rightclick: %d, keypress: %d\n", leftClickCount, rightClickCount, keyPress)
}

func main() {
	log.Print("Starting logger")
	eventChan := hook.Start()
	defer hook.End()
	go logger(eventChan)
	//1*time.Hour here
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			updateDb()
		}
	}
}
