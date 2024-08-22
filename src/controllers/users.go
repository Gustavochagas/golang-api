package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Search users
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Search user")))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		log.Fatal(erro)
	}

	repository := repositories.NewRepositoryForUser(db)
	userID, erro := repository.Create(user)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("Insert ID: %d", userID)))
}

// Search user by
func SearchUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Search user")))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Update user")))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write(([]byte("Delete user")))
}
