package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type chartinfo struct {
	ID         int    `json:"id"`
	LeftClick  int    `json:"leftclick"`
	RightClick int    `json:"rightclick"`
	KeyPress   int    `json:"keypress"`
	Time       string `json:"time"`
}

func getChartinfo(db *sql.DB) ([]chartinfo, error) {
	rows, err := db.Query("SELECT * FROM chartinfo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chartData []chartinfo
	for rows.Next() {
		var ci chartinfo
		err := rows.Scan(&ci.ID, &ci.LeftClick, &ci.RightClick, &ci.KeyPress, &ci.Time)
		if err != nil {
			return nil, err
		}
		chartData = append(chartData, ci)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return chartData, nil
}

func chartInfoHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chartData, err := getChartinfo(db)
		if err != nil {
			http.Error(w, "Failed to fetch chart info", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(chartData); err != nil {
			http.Error(w, "Failed to encode chart info to JSON", http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect to database", err)
		//TODO: Send error message in json
	}
	fmt.Println("Database connected successfully")
	fmt.Println("--------------------------------")
	http.HandleFunc("/chartinfo", chartInfoHandler(db))
	log.Println("Starting server at port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Failed to start the server", err)
	}
}
