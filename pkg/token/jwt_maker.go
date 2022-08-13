package token

import (
	"github.com/danyouknowme/ecommerce/pkg/util"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(username string) (string, error) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		return "", err
	}

	payload, err := NewPayload(username)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(config.TokenSymmetricKey))
}
