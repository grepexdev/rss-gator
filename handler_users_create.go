package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/grepexdev/rss-gator/internal/database"
)

func (cfg *apiConfig) handlerUsersCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating user...")
	type parameters struct {
		Name string `json:"name"`
	}
	type response struct {
		User
	}

	decoder := json.NewDecoder(r.Body)
	fmt.Println("json decoded")
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	fmt.Println("parameters decoded")

	user, err := cfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}
	fmt.Println("user created")
	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
