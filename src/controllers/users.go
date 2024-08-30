package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("register"); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryForUser(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func SearchUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryForUser(db)
	users, erro := repository.Search(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryForUser(db)
	user, erro := repository.SearchById(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIdFromToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdFromToken {
		responses.Erro(w, http.StatusForbidden, errors.New("the user being modified is not the logged in one"))
		return
	}

	bodyRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("edit"); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryForUser(db)
	erro = repository.Update(userId, user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userId, erro := strconv.ParseUint(parameters["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userIdFromToken, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userId != userIdFromToken {
		responses.Erro(w, http.StatusForbidden, errors.New("the user being modified is not the logged in one"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryForUser(db)
	erro = repository.Delete(userId)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userId {
		responses.Erro(w, http.StatusForbidden, errors.New("not possible follow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryForUser(db)
	if erro = repository.Follow(userId, followerId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, erro := authentication.ExtractUserId(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	userId, erro := strconv.ParseUint(params["userId"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followerId == userId {
		responses.Erro(w, http.StatusForbidden, errors.New("not possible unfollow yourself"))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryForUser(db)
	if erro = repository.Unfollow(userId, followerId); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
