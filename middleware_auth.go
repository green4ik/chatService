package main

import (
	"net/http"

	"github.com/green4ik/chatservice/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {

// }
