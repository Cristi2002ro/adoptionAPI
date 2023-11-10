package main

import (
	"adoptionAPI/dal"
	"adoptionAPI/handlers"
	"net/http"
)

const BaseURL = "/api"

func main() {
	http.HandleFunc(BaseURL+"/categories", handlers.CategoriesHandler)
	http.HandleFunc(BaseURL+"/animals", handlers.HandleAllAnimals)
	http.HandleFunc(BaseURL+"/animals/add", handlers.HandleAddAnimal)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	defer dal.GetDB().Close()
}
