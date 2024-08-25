package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/green4ik/chatservice/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port is not found!")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Database is not found!")
	}
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed connection to database!")
	}
	queries := database.New(conn)
	apiCfg := apiConfig{
		DB: queries,
	}
	router := chi.NewRouter()

	userRouter := chi.NewRouter()
	feedRouter := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	router.Mount("/users", userRouter)
	userRouter.Post("/add", apiCfg.handlerCreateUser)
	userRouter.Get("/get", apiCfg.middlewareAuth(apiCfg.handlerGetUserByApiKey))

	router.Mount("/feeds", feedRouter)
	feedRouter.Post("/add", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	feedRouter.Get("/get", apiCfg.handlerGetFeeds)

	fmt.Println("Server started and running at port : ", portString,
		"\nCtrl+C to stop manually")
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
