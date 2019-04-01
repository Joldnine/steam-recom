package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Game is the return game type
type Game struct {
	Name string
	ID   string
}

type Recom struct {
	UserID string   `json:"user_id"`
	Games  []string `json:"recommend_games"`
}

type UserGame struct {
	UserID string   `json:"user_id"`
	Games  []string `json:"games"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func favoriteHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid input.")
		return
	}
	var games []Game
	// -------------logic here
	ids := userGameMapping[userID]
	if ids == nil {
		games = []Game{}
	} else {
		for _, id := range ids {
			games = append(games, Game{ID: id})
		}
	}
	// -------------logic above
	jsonResult, err := json.Marshal(games)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

func recommendHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		w.WriteHeader(422)
		fmt.Fprintf(w, "Invalid input.")
		return
	}
	var games []Game
	// -------------logic here
	ids := recomMapping[userID]
	if ids == nil {
		ids = []string{"730", "304930", "105600", "4000", "218620", "550", "230410", "72850", "620", "240", "301520", "252490", "49520", "236390", "219640", "220", "8190", "218230", "322330", "227940"}
	}
	for _, id := range ids {
		games = append(games, Game{ID: id})
	}

	// -------------logic above

	jsonResult, err := json.Marshal(games)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)
}

var recomMapping map[string][]string
var userGameMapping map[string][]string

func main() {
	// recommend data
	file, _ := ioutil.ReadFile("result.json")
	var data []Recom
	json.Unmarshal([]byte(file), &data)
	recomMapping = make(map[string][]string)
	for _, user := range data {
		recomMapping[user.UserID] = user.Games
	}
	// users' data
	file, _ = ioutil.ReadFile("user_game_result.json")
	var userGame []UserGame
	json.Unmarshal([]byte(file), &userGame)
	userGameMapping = make(map[string][]string)
	for _, user := range userGame {
		userGameMapping[user.UserID] = user.Games
	}

	// handlers
	http.HandleFunc("/favorite", favoriteHandler)
	http.HandleFunc("/recommend", recommendHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
