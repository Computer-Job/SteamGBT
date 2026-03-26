package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Game struct {
	AppID int `json:"appid"`
	Name string `json:"name"`
	PlaytimeForever int `json:"playtime_forever"`
}

type OwnedGamesData struct {
	GameCount int `json:"game_count"`
	Games []Game `json:"games"`
}

type OwnedGamesResponse struct {
	Response OwnedGamesData `json:"response"`
}

func main() {
	apiKey := os.Getenv("STEAM_API_KEY")
	steamID := os.Getenv("STEAM_ID")

	baseURL := "https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/"

	params := url.Values{}
	params.Set("key", apiKey)
	params.Set("steamid", steamID)
	params.Set("include_appinfo", "true")
	params.Set("include_played_free_games", "true")
	params.Set("format", "json")
	
	fullURL := baseURL + "?" + params.Encode()

	client := &http.Client {
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(fullURL)	
	if (err != nil) {
		panic(err)
	}
	defer resp.Body.Close()

	if (resp.StatusCode != http.StatusOK) {
		panic(fmt.Sprintf("steam returned status %d", resp.StatusCode))
	}

	var result OwnedGamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		panic(err)
	}

	// fmt.Printf("Owned Games: %d\n", result.Response.GameCount)
	// for _, game := range result.Response.Games {
	// 	fmt.Printf("%d - %s (%d mins)\n", game.AppID, game.Name, game.PlaytimeForever)
	// }

	// On Front-Facing Website
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Owned Games: %d\n", result.Response.GameCount)
		for _, game := range result.Response.Games {
			fmt.Fprintf(w, "%d - %s (%d mins)\n", game.AppID, game.Name, game.PlaytimeForever)
		}
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	// Start server on port specified
	log.Fatal(http.ListenAndServe(port, nil))
 }