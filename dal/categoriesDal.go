package dal

import (
	"adoptionAPI/model"
	"log"
)

func GetCategories() ([]model.Category, error) {
	rows, err := db.Query("SELECT * FROM Categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []model.Category

	for rows.Next() {
		var category model.Category
		if err := rows.Scan(&category.Id, &category.Value); err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)
	}
	return categories, nil
}
