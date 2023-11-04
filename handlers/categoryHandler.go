package handlers

import (
	"adoptionAPI/dal"
	"encoding/json"
	"net/http"
)

func setupResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CategoryHandler(writer http.ResponseWriter, r *http.Request) {
	response, _ := dal.GetCategories()
	setupResponse(writer)
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		return
	}
}
