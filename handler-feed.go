package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/green4ik/chatservice/internal/database"
)

func (apiConfig apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parametrs struct {
		Name string `json:"name"`
		URL  string `json:"url`
	}
	decoder := json.NewDecoder(r.Body)

	params := parametrs{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}
	feed, err := apiConfig.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC().In(time.FixedZone("UTC+3", 3*3600)),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to add a feed %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}
