package handlers

import (
	"adoptionAPI/dal"
	"adoptionAPI/util"
	"encoding/json"
	"net/http"
)

func CategoriesHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet && request.Method != http.MethodOptions {
		http.Error(writer, "GET method only available for this endpoint", http.StatusMethodNotAllowed)
		return
	}
	if request.URL.Query().Get("id") != "" {
		getCategoryByIdHandler(writer, request)
	} else {
		getAllCategoriesHandler(writer, request)
	}
}

func getAllCategoriesHandler(writer http.ResponseWriter, request *http.Request) {
	response, _ := dal.GetCategories()
	util.Setup200Response(writer, request)
	json.NewEncoder(writer).Encode(response)
}

func getCategoryByIdHandler(writer http.ResponseWriter, request *http.Request) {
	response, _ := dal.GetCategoryById(request)
	util.Setup200Response(writer, request)
	json.NewEncoder(writer).Encode(response)
}
