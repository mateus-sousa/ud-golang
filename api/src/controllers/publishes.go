package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreatePublish(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var publish models.Publish
	if err = json.Unmarshal(bodyRequest, &publish); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = publish.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	publish.AuthorID = userID

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repo := repository.NewPublishRepository(db)
	publish.ID, err = repo.Create(publish)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, publish)
}

func GetPublishes(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	repo := repository.NewPublishRepository(db)
	publishes, err := repo.GetPublishes(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publishes)
}

func GetPublish(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	publishId, err := strconv.ParseUint(params["publishId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repo := repository.NewPublishRepository(db)
	publish, err := repo.GetPublish(publishId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publish)
}

func UpdatePublish(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publishId, err := strconv.ParseUint(params["publishId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	repo := repository.NewPublishRepository(db)
	storedPublish, err := repo.GetPublish(publishId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if storedPublish.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errors.New("Não é possivel alterar uma publicação que não seja a tua."))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	var publish models.Publish
	err = json.Unmarshal(requestBody, &publish)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = publish.Prepare()
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	err = repo.Update(publish, userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletePublish(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.ExtractUserIDFromToken(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	publishID, err := strconv.ParseUint(params["publishId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	repo := repository.NewPublishRepository(db)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	storedPublish, err := repo.GetPublish(publishID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if storedPublish.AuthorID != userID {
		responses.Error(w, http.StatusInternalServerError, errors.New("Não é possível deletar uma publicação que não seja a sua"))
		return
	}

	if err = repo.Delete(publishID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, nil)
}

func GetPublishesByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewPublishRepository(db)
	publishes, err := repo.GetPublishesByUser(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, publishes)
	return
}
