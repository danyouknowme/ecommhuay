package dbmodels

import (
	"database/sql"
	"fmt"

	"github.com/danyouknowme/ecommhuay/pkg/database"
)

type Cart struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

type ProductInCart struct {
	Id        int `json:"id"`
	CartId    int `json:"cart_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type ProductInCartRequest struct {
	Username        string `json:"username"`
	ProductId       int    `json:"product_id"`
	IsAddedQuantity bool   `json:"is_added_quantity"`
}

func AddOrUpdateProductInCart(username string, productId int, isAddedQuantity bool) error {
	db := database.DB
	var productInCart ProductInCart

	statementInsert, err := db.Prepare("INSERT INTO ProductInCart (CartId, ProductId, Quantity) VALUES ( ?, ?, ? )")
	if err != nil {
		return err
	}

	defer statementInsert.Close()

	statementUpdate, err := db.Prepare("UPDATE ProductInCart SET Quantity = ? WHERE CartId = ? AND ProductId = ?")
	if err != nil {
		return err
	}

	defer statementUpdate.Close()

	statementDelete, err := db.Prepare("DELETE FROM ProductInCart WHERE CartId = ? AND ProductId = ?")
	if err != nil {
		return err
	}

	defer statementDelete.Close()

	user, err := GetUser(username)
	if err != nil {
		return err
	}

	cart, err := GetUserCart(user.Id)
	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT * FROM ProductInCart WHERE CartId = ? AND ProductId = ?", cart.Id, productId).Scan(&productInCart.Id, &productInCart.CartId, &productInCart.ProductId, &productInCart.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = statementInsert.Exec(cart.Id, productId, 1)
			if err != nil {
				return err
			}
			return fmt.Errorf("added product in cart successfully")
		}
		return err
	}

	if isAddedQuantity {
		productInCart.Quantity = productInCart.Quantity + 1
	} else {
		productInCart.Quantity = productInCart.Quantity - 1
	}

	if productInCart.Quantity == 0 {
		_, err = statementDelete.Exec(cart.Id, productId)
		if err != nil {
			return err
		}
		return fmt.Errorf("deleted product in cart successfully")
	}

	_, err = statementUpdate.Exec(productInCart.Quantity, cart.Id, productId)
	if err != nil {
		return err
	}

	return nil
}

func GetUserCart(userId int) (Cart, error) {
	db := database.DB
	var cart Cart

	result := db.QueryRow("SELECT * FROM Carts WHERE UserId = ?", userId)
	err := result.Scan(&cart.Id, &cart.UserId, &cart.CreatedAt)
	if err != nil {
		return Cart{}, err
	}

	return cart, nil
}
