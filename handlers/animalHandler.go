package handlers

import (
	"adoptionAPI/dal"
	"adoptionAPI/model"
	"adoptionAPI/util"
	"encoding/json"
	"net/http"
	"strings"
)

var acceptedParams = [...]string{"id", "categoryId", "breedId", "age", "gender", "adopted", "reserved", "userId",
	"location", "startAge", "endAge"}

func HandleGetAnimals(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet && request.Method != http.MethodOptions {
		http.Error(writer, "GET method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	params := request.URL.Query()
	whereClause := false
	checkParams(params, &whereClause, writer)

	response, _ := dal.GetAnimals(params, whereClause)
	util.Setup200Response(writer, request)
	if response == nil {
		response = make([]model.Animal, 0)
	}
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

func HandleReserveAnimal(writer http.ResponseWriter, request *http.Request) {
	util.Setup200Response(writer, request)
	if request.Method != http.MethodPut && request.Method != http.MethodOptions {
		http.Error(writer, "POST method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	var userId, animalId string
	pathParams := strings.TrimPrefix(request.URL.Path, "/api/reserve/")
	parameters := strings.Split(pathParams, "/")
	for range parameters {
		userId = parameters[0]
		animalId = parameters[1]
	}

	dal.ReserveAnimal(userId, animalId, writer)
	util.Setup200Response(writer, request)
}

func checkParams(params map[string][]string, whereClause *bool, writer http.ResponseWriter) {
	if len(params) != 0 {
		*whereClause = true
		for key := range params {
			isValid := false
			for i := range acceptedParams {
				if key == acceptedParams[i] {
					isValid = true
				}
			}
			if !isValid {
				http.Error(writer, "Unaccepted query params", http.StatusBadRequest)
				return
			}
		}
	}
}
