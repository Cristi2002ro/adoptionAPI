package handlers

import (
	"adoptionAPI/model"
	"adoptionAPI/util"
	"encoding/json"
	"log"
	"net/http"
)

func HandleAnimalById(writer http.ResponseWriter, r *http.Request) {
	animalId := r.URL.Query().Get("id")
	log.Println("AnimalId = ", animalId)
	response := model.Animal{Id: "1", CategoryId: "1", Name: "Nero"}
	util.Setup200Response(writer, r)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		return
	}
}
