package handlers

import (
	"adoptionAPI/dal"
	"adoptionAPI/model"
	"adoptionAPI/util"
	"encoding/json"
	"net/http"
)

func HandleGetFavorites(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet && request.Method != http.MethodOptions {
		http.Error(writer, "GET method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	userID := request.URL.Query().Get("userId")
	animalId := request.URL.Query().Get("animalId")
	if userID == "" {
		http.Error(writer, "No userId provided", http.StatusBadRequest)
		return
	}
	favoritesList := dal.GetFavoritesByUid(userID, animalId)
	if favoritesList == nil {
		favoritesList = make([]model.Animal, 0)
	}
	util.Setup200Response(writer, request)
	json.NewEncoder(writer).Encode(favoritesList)
}

func HandlePostFavorite(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost && request.Method != http.MethodOptions {
		http.Error(writer, "POST method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	var favorite model.Favorite
	err := json.NewDecoder(request.Body).Decode(&favorite)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	dal.PostFavorite(favorite, writer)
	util.Setup200Response(writer, request)
}

func HandleDeleteFavorite(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodDelete && request.Method != http.MethodOptions {
		http.Error(writer, "POST method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	var favorite model.Favorite
	err := json.NewDecoder(request.Body).Decode(&favorite)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	dal.DeleteFavorite(favorite, writer)
	util.Setup200Response(writer, request)
}
