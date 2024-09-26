package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	hook "github.com/robotn/gohook"
)

var leftClickCount, rightClickCount, keyPress int
var failedWrites []chartinfo
var mu sync.Mutex

type chartinfo struct {
	id         int
	leftclick  int
	rightlcick int
	keypress   int
	time       string
}

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
			// fmt.Printf("Left Click count: %d\n", leftClickCount)
		case <-rightClickChan:
			rightClickCount++
			// fmt.Printf("Right click count:%d\n", rightClickCount)
		case <-keyPressChan:
			keyPress++
			// fmt.Printf("Key press count: %d\n", keyPress)
		}
	}
}

func updateDb() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading the env file")
		return
	}
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Printf("Error opening the database connection %v", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	currentTime := time.Now().Format("2006-01-02 15:04")
	log.Printf("Updating database with LeftClick: %d Rightclick: %d Keypress: %d Time: %s\n", leftClickCount, rightClickCount, keyPress, currentTime)
	query := `
	insert into chartinfo (leftclick,rightclick,keypress,time)
	values($1,$2,$3,$4)
	`
	_, err = db.Exec(query, leftClickCount, rightClickCount, keyPress, currentTime)
	if err != nil {
		log.Printf("Error inserting data into the database: %v", err)
		mu.Lock()
		failedWrites = append(failedWrites, chartinfo{
			leftclick:  leftClickCount,
			rightlcick: rightClickCount,
			keypress:   keyPress,
			time:       currentTime,
		})
		// reset the count to 0 to start logging events for another time interval
		// can be a better way to do this idk
		leftClickCount = 0
		rightClickCount = 0
		keyPress = 0
		mu.Unlock()
	} else {
		log.Printf("Successfully updated the database with LeftClick: %d Rightclick: %d Keypress: %d Time: %s", leftClickCount, rightClickCount, keyPress, currentTime)
		leftClickCount = 0
		rightClickCount = 0
		keyPress = 0
		fmt.Printf("Reset counts Leftclick: %d Rightclick: %d, keypress: %d\n", leftClickCount, rightClickCount, keyPress)
	}
	//retry when there is a successfull connection
	retryfailedWrites(db)
}

func retryfailedWrites(db *sql.DB) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(failedWrites)
	for i := 0; i < len(failedWrites); i++ {
		failedData := failedWrites[i]
		fmt.Println(failedData)
		query := `
	insert into chartinfo (leftclick,rightclick,keypress,time)
	values($1,$2,$3,$4)
	`
		_, err := db.Exec(query, failedData.leftclick, failedData.rightlcick, failedData.keypress, failedData.time)
		if err != nil {
			log.Printf("Write retry failed for data %v\n", failedData, err)
		} else {
			log.Printf("Write retry successful for data %v \n", failedData)
			failedWrites = append(failedWrites[:i], failedWrites[i+1:]...)
			i--
			fmt.Println(failedWrites)
		}
	}
}

func main() {
	log.Print("Starting logger")
	eventChan := hook.Start()
	defer hook.End()
	go logger(eventChan)
	//1*time.Hour here
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			updateDb()
		}
	}
}
