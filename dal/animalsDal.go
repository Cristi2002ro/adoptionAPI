package dal

import (
	"adoptionAPI/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

func GetAnimals(params map[string][]string, whereClause bool) ([]model.Animal, error) {
	query := buildQuery(params, whereClause)
	fmt.Println("Query:", query)
	rows, _ := db.Query(query)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var animals []model.Animal
	for rows.Next() {
		var animal model.Animal
		err := rows.Scan(&animal.Id, &animal.CategoryId, &animal.Name, &animal.Age, &animal.Species, &animal.Gender,
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
	_, err := db.Exec("INSERT INTO animals (id, categoryId, name, age, species, gender, weight, reservationDate, isReserved, isAdopted, image, location) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		animalId, animal.CategoryId, animal.Name, animal.Age, animal.Species, animal.Gender, animal.Weight, animal.ReservationDate, animal.IsReserved, animal.IsAdopted, animal.Image, animal.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(`Successfully inserted`)
}

func buildQuery(params map[string][]string, whereClause bool) string {
	baseQuery := "SELECT * FROM animals"
	if whereClause {
		baseQuery += " where "
		for key, value := range params {
			baseQuery += key + " = '" + value[0] + "' and "
		}
		baseQuery = baseQuery[0 : len(baseQuery)-5]
	}
	return baseQuery
}
