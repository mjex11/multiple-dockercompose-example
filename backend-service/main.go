package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Fortune struct {
	Result string `json:"result"`
}

var fortunes = []string{
	"大吉",
	"吉",
	"中吉",
	"小吉",
	"末吉",
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

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(fortunes))
	fortuneResult := fortunes[randomIndex]

	fortune := Fortune{Result: fortuneResult}
	json.NewEncoder(w).Encode(fortune)
}
