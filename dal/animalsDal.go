package dal

import (
	"adoptionAPI/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strings"
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
		err := rows.Scan(&animal.Id, &animal.CategoryId, &animal.BreedId, &animal.Name, &animal.Image, &animal.Age, &animal.Weight, &animal.Gender,
			&animal.Adopted, &animal.Reserved, &animal.UserId, &animal.Description, &animal.Location)
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
	_, err := db.Exec("INSERT INTO animals (id, categoryId, breedId, name, image, age, weight, gender,adopted, reserved, userId, description, location) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
		animalId, animal.CategoryId, animal.BreedId, animal.Name, animal.Image, animal.Age, animal.Weight, animal.Gender,
		animal.Adopted, animal.Reserved, animal.UserId, animal.Description, animal.Location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(`Successfully inserted`)
}

func ReserveAnimal(userId, animalId string, w http.ResponseWriter) {
	fmt.Println("user:", userId, "animal:", animalId)
	_, err := db.Exec("UPDATE animals set userId=$1, reserved =true where id=$2", userId, animalId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(`Reservation succesfully created`)
}

func buildQuery(params map[string][]string, whereClause bool) string {
	baseQuery := "SELECT * FROM animals"
	if whereClause {
		baseQuery += " where "
		for key, value := range params {
			if key != "startAge" && key != "endAge" && !strings.Contains(value[0], ",") {
				baseQuery += key + " = '" + value[0] + "' and "
			} else {
				switch key {
				case "startAge":
					baseQuery += "age >= " + value[0] + " and "
				case "endAge":
					baseQuery += "age <= " + value[0] + " and "
				case "categoryId":
					categories := strings.Split(value[0], ",")
					baseQuery += "("
					for i := range categories {
						baseQuery += "categoryId = '" + categories[i] + "' or "
					}
					baseQuery = baseQuery[0:len(baseQuery)-4] + ") and "
				case "breedId":
					breeds := strings.Split(value[0], ",")
					baseQuery += "("
					for i := range breeds {
						baseQuery += "breedId = '" + breeds[i] + "' or "
					}
					baseQuery = baseQuery[0:len(baseQuery)-4] + ") and "
				}
			}
		}
		baseQuery = baseQuery[0 : len(baseQuery)-5]
	}
	return baseQuery
}
