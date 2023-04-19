package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"database/sql"
	_ "github.com/lib/pq"
)

type Fortune struct {
	Result string `json:"result"`
}

var db *sql.DB

func init() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/fortune", GetFortune).Methods("GET")

	corsWrapper := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	handler := corsWrapper.Handler(r)

	log.Fatal(http.ListenAndServe(":8000", handler))
}

func GetFortune(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var fortuneResult string
	err := db.QueryRow("SELECT result FROM fortunes ORDER BY RANDOM() LIMIT 1").Scan(&fortuneResult)
	if err != nil {
		http.Error(w, "Error fetching fortune from the database", http.StatusInternalServerError)
		log.Printf("Error fetching fortune from database: %v", err)
		return
	}

	fortune := Fortune{Result: fortuneResult}
	json.NewEncoder(w).Encode(fortune)
}
