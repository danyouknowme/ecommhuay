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
	Amount      int     `json:"amount"`
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
		result.Scan(&product.Id, &product.Title, &product.Description, &product.ImagePath, &product.Category, &product.Price, &product.Amount)
		products = append(products, product)
	}

	return products, nil
}

func GetProductById(id int) (Product, error) {
	db := database.DB
	var product Product

	result := db.QueryRow("SELECT * FROM Products WHERE id = ?", id)
	err := result.Scan(&product.Id, &product.Title, &product.Description, &product.ImagePath, &product.Category, &product.Price, &product.Amount)
	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func AddNewProduct(newProduct Product) error {
	db := database.DB

	statementInsert, err := db.Prepare("INSERT INTO Products (Title, Description, ImagePath, Category, Price, Amount) VALUES ( ?, ?, ?, ?, ?, ? )")
	if err != nil {
		return err
	}

	defer statementInsert.Close()

	_, err = statementInsert.Exec(newProduct.Title, newProduct.Description, newProduct.ImagePath, newProduct.Category, newProduct.Price, newProduct.Amount)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProductAmount(id int) error {
	db := database.DB

	statementUpdate, err := db.Prepare("UPDATE Products SET Amount = Amount - 1 WHERE Id = ?")
	if err != nil {
		return err
	}

	defer statementUpdate.Close()

	_, err = statementUpdate.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
