package dbmodels

import (
	"database/sql"
	"fmt"

	"github.com/danyouknowme/ecommhuay/pkg/database"
	"github.com/danyouknowme/ecommhuay/pkg/token"
	"github.com/danyouknowme/ecommhuay/pkg/util"
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

type UserResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

type GetUserRequest struct {
	Username string `json:"username"`
}

func Register(user User) error {
	db := database.DB
	var userId int64

	statementInsertUser, err := db.Prepare(`INSERT INTO Users (
		Username,
		Password,
		FullName,
		Email
	) VALUES( ?, ?, ?, ? )`)
	if err != nil {
		return err
	}

	defer statementInsertUser.Close()

	statementInsertCart, err := db.Prepare(`INSERT INTO Carts (UserId) VALUES ( ? )`)
	if err != nil {
		return err
	}

	defer statementInsertCart.Close()

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	res, err := statementInsertUser.Exec(user.Username, hashedPassword, user.FullName, user.Email)
	if err != nil {
		return err
	}

	userId, err = res.LastInsertId()
	if err != nil {
		return err
	}

	_, err = statementInsertCart.Exec(userId)
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

	var accessToken string
	if accessToken, err = token.CreateToken(user.Username); err != nil {
		return LoginResponse{}, fmt.Errorf("cannot create a new token")
	}

	response := LoginResponse{
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Token:    accessToken,
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
