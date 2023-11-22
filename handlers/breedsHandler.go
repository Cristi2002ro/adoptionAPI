package handlers

import (
	"adoptionAPI/dal"
	"adoptionAPI/model"
	"adoptionAPI/util"
	"encoding/json"
	"net/http"
	"strings"
)

func HandleGetBreeds(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet && request.Method != http.MethodOptions {
		http.Error(writer, "GET method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	categoryId := strings.TrimPrefix(request.URL.Path, "/api/breeds/")
	response, _ := dal.GetBreedsForCategory(categoryId)
	if response == nil {
		response = make([]model.Breed, 0)
	}
	util.Setup200Response(writer, request)
	json.NewEncoder(writer).Encode(response)
}
