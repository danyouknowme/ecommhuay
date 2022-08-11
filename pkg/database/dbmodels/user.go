package dbmodels

import (
	"database/sql"
	"fmt"

	"github.com/danyouknowme/ecommerce/pkg/database"
	"github.com/danyouknowme/ecommerce/pkg/util"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
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

func Login(username string, password string) (LoginResponse, error) {
	user, err := GetUser(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return LoginResponse{}, fmt.Errorf("user with %s not found", username)
		}
		return LoginResponse{}, err
	}

	err = util.CheckPassword(password, user.Password)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("password is not match")
	}

	response := LoginResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Token:    "abcdefghijklmnopqrstuvwxyz",
	}

	return response, nil
}

func GetUser(username string) (User, error) {
	db := database.DB
	var user User

	result := db.QueryRow("SELECT * FROM Users WHERE Username = ?", username)
	err := result.Scan(&user.Id, &user.Username, &user.Password, &user.FullName, &user.Email, &user.IsAdmin)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
