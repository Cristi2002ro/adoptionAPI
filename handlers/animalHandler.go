package handlers

import (
	"adoptionAPI/dal"
	"adoptionAPI/model"
	"adoptionAPI/util"
	"encoding/json"
	"net/http"
)

func HandleAllAnimals(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet && request.Method != http.MethodOptions {
		http.Error(writer, "GET method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	response, _ := dal.GetAnimals()
	util.Setup200Response(writer, request)
	json.NewEncoder(writer).Encode(response)
}

func HandleAddAnimal(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost && request.Method != http.MethodOptions {
		http.Error(writer, "POST method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	var newAnimal model.Animal
	err := json.NewDecoder(request.Body).Decode(&newAnimal)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	dal.AddAnimal(newAnimal, writer)
	util.Setup200Response(writer, request)
}
