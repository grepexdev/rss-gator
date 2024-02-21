package main

import (
	"fmt"
	"net/http"

	"github.com/grepexdev/rss-gator/internal/auth"
	"github.com/grepexdev/rss-gator/internal/database"
)

// custom type for handlers that require authentication
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// middleware that creates a request, gets a user, and calls the next authed
// handler
func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "couldn't find api key")
			return
		}
		fmt.Println("api key retrieved")

		user, err := cfg.DB.RetrieveUser(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusNotFound, "couldn't get user")
		}
		fmt.Println("user retrieved")

		handler(w, r, user)
	}
}
