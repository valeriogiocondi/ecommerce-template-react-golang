package jwtUtils

import (
	env "customers/utils/env"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	Payload interface{} `json:"payload"`
	jwt.StandardClaims
}

func Tokerize(obj interface{}) (string, error) {

	expTime, _ := strconv.Atoi(env.ReadFile("JWT_EXP_TIME_HOURS"))
	timeNow := time.Now().Local()
	hashStr := fmt.Sprintf("%d", timeNow.Unix())
	claims := CustomClaims{
		obj,
		jwt.StandardClaims{
			Id:        hashStr,
			IssuedAt:  timeNow.Unix(),
			ExpiresAt: timeNow.Add(time.Hour * time.Duration(expTime)).Unix(),
			Issuer:    env.ReadFile("JWT_ISSUER"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(env.ReadFile("JWT_SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
