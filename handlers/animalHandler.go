package handlers

import (
	"adoptionAPI/model"
	"encoding/json"
	"log"
	"net/http"
)

func HandleAnimalById(writer http.ResponseWriter, r *http.Request) {
	animalId := r.URL.Query().Get("id")
	log.Println("AnimalId = ", animalId)
	response := model.Animal{Id: "1", CategoryId: "1", Name: "Nero"}
	setupResponse(writer)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		return
	}
}
