package dal

import (
	"adoptionAPI/model"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func GetAnimals() ([]model.Animal, error) {
	rows, _ := db.Query("SELECT * FROM animals")
	defer rows.Close()

	var animals []model.Animal
	for rows.Next() {
		var animal model.Animal
		err := rows.Scan(&animal.Id, &animal.CategoryId, &animal.Name, &animal.ShelterId, &animal.Age, &animal.Species, &animal.Gender,
			&animal.Weight, &animal.ReservationDate, &animal.IsReserved, &animal.IsAdopted, &animal.Image, &animal.Location)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func AddAnimal(animal model.Animal, w http.ResponseWriter) {
	animalId, _ := uuid.NewUUID()
	_, err := db.Exec("INSERT INTO animals (id, categoryId, name, shelterId, age, species, gender, weight, reservationDate, isReserved, isAdopted, image, location) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
		animalId, animal.CategoryId, animal.Name, animal.ShelterId, animal.Age, animal.Species, animal.Gender, animal.Weight, animal.ReservationDate, animal.IsReserved, animal.IsAdopted, animal.Image, animal.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(`Successfully inserted`)
}
