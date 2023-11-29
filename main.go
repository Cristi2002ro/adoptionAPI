package main

import (
	"adoptionAPI/dal"
	"adoptionAPI/handlers"
	"net/http"
)

const BaseURL = "/api"

func main() {
	http.HandleFunc(BaseURL+"/categories", handlers.CategoriesHandler)

	http.HandleFunc(BaseURL+"/animals", handlers.HandleGetAnimals)
	http.HandleFunc(BaseURL+"/animals/add", handlers.HandleAddAnimal)
	http.HandleFunc(BaseURL+"/reserve/", handlers.HandleReserveAnimal)

	http.HandleFunc(BaseURL+"/breeds/", handlers.HandleGetBreeds)

	http.HandleFunc(BaseURL+"/favorites/add", handlers.HandlePostFavorite)
	http.HandleFunc(BaseURL+"/favorites", handlers.HandleGetFavorites)
	http.HandleFunc(BaseURL+"/favorites/delete", handlers.HandleDeleteFavorite)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	defer dal.GetDB().Close()
}
