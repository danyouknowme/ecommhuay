package dbmodels

import (
	"github.com/danyouknowme/ecommerce/pkg/database"
)

type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImagePath   string  `json:"imagePath"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

func GetAllProducts() ([]Product, error) {
	db := database.DB
	result, err := db.Query("SELECT * FROM Products")
	if err != nil {
		return nil, err
	}

	defer result.Close()

	products := []Product{}
	for result.Next() {
		var product Product
		result.Scan(&product.Id, &product.Title, &product.Description, &product.ImagePath, &product.Category, &product.Price)
		products = append(products, product)
	}

	return products, nil
}
