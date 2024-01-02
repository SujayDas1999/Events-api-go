package helpers

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const secretKey = "secretKey"

func GenerateToken(email string, userid uuid.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userid,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid Token")
		}

		return []byte(secretKey), nil
	})
	if err != nil {
		return "", errors.New("invalid")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return "", errors.New("invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("invalid Token")
	}

	//email := claims["email"].(string)
	userid := claims["userId"].(string)

	return userid, nil
}
