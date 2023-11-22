package dal

import (
	"adoptionAPI/model"
	"fmt"
)

func GetBreedsForCategory(categoryId string) ([]model.Breed, error) {
	query := fmt.Sprintf("select * from breeds where categoryId='%s'", categoryId)
	fmt.Println(query)
	rows, _ := db.Query(query)
	defer rows.Close()

	var breeds []model.Breed
	for rows.Next() {
		var breed model.Breed
		rows.Scan(&breed.Id, &breed.CategoryId, &breed.Name)
		breeds = append(breeds, breed)
	}
	return breeds, nil
}
