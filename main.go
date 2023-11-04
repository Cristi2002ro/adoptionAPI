package main

import (
	"adoptionAPI/dal"
	"adoptionAPI/handlers"
	"net/http"
)

const BaseURL = "/api"

func main() {
	http.HandleFunc(BaseURL+"/categories", handlers.CategoryHandler)
	http.HandleFunc(BaseURL+"/animal", handlers.HandleAnimalById)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	defer dal.GetDB().Close()
}
