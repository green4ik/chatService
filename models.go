package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/green4ik/chatservice/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
	}
}
