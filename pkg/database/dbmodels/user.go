package dbmodels

import (
	"github.com/danyouknowme/ecommerce/pkg/database"
	"github.com/danyouknowme/ecommerce/pkg/util"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func Register(user User) error {
	db := database.DB

	statementInsert, err := db.Prepare("INSERT INTO Users (Username, Password, FullName, Email) VALUES ( ?, ?, ?, ? )")
	if err != nil {
		return err
	}

	defer statementInsert.Close()

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = statementInsert.Exec(user.Username, hashedPassword, user.FullName, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func GetUser(username string) (User, error) {
	db := database.DB
	var user User

	result := db.QueryRow("SELECT * FROM Users WHERE Username = ?", username)
	err := result.Scan(&user.Username, &user.FullName, &user.Email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
