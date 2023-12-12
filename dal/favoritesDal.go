package dal

import (
	"adoptionAPI/model"
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

func PostFavorite(favorite model.Favorite, writer http.ResponseWriter) {
	favoriteId, _ := uuid.NewUUID()
	_, err := db.Exec("INSERT INTO favorites (id, userId, animalId) VALUES ($1, $2, $3)",
		favoriteId, favorite.UserId, favorite.AnimalId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(writer).Encode(`Successfully inserted`)
}

func GetFavoritesByUid(userID string, animalId string) []model.Animal {
	var result *sql.Rows

	if animalId == "" {
		query := "SELECT animals.* FROM animals JOIN favorites ON animals.id = favorites.animalId WHERE favorites.userId = $1;"
		result, _ = db.Query(query, userID)
	} else {
		query := "SELECT animals.* FROM animals JOIN favorites ON animals.id = favorites.animalId WHERE favorites.userId = $1 AND favorites.animalId=$2;"
		result, _ = db.Query(query, userID, animalId)
	}

	var animals []model.Animal
	for result.Next() {
		var animal model.Animal
		result.Scan(&animal.Id, &animal.CategoryId, &animal.BreedId, &animal.Name, &animal.Image, &animal.Age, &animal.Weight, &animal.Gender,
			&animal.Adopted, &animal.Reserved, &animal.UserId, &animal.Description, &animal.Location)
		animals = append(animals, animal)
	}
	return animals
}

func DeleteFavorite(favorite model.Favorite, writer http.ResponseWriter) {
	_, err := db.Exec("DELETE FROM favorites WHERE userId=$1 AND animalId=$2", favorite.UserId, favorite.AnimalId)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(writer).Encode(`Successfully deleted`)
}
