package dal

import (
	"adoptionAPI/model"
	"net/http"
)

func GetCategories() ([]model.Category, error) {
	rows, _ := db.Query("SELECT * FROM categories")
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		rows.Scan(&category.Id, &category.Value, &category.Image)
		categories = append(categories, category)
	}
	return categories, nil
}

func GetCategoryById(request *http.Request) (model.Category, error) {
	selectStmt := `SELECT * FROM categories WHERE id=$1`
	id := request.URL.Query().Get("id")
	rows, _ := db.Query(selectStmt, id)
	defer rows.Close()

	for rows.Next() {
		var category model.Category
		rows.Scan(&category.Id, &category.Value, &category.Image)
		return category, nil
	}

	return model.Category{}, nil
}
