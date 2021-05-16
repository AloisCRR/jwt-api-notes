package models

import (
	"fmt"
	"github.com/AloisCRR/jwt-api-notes/helpers"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type AccessDetails struct {
	UserEmail string
}

func CreateJWT(userEmail string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_email"] = userEmail
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(helpers.GetEnvVariable("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ExtractJWT(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	bearerToken := strings.Split(authHeader, " ")
	if len(bearerToken) == 2 {
		return bearerToken[1]
	}
	return ""
}

func CheckToken(r *http.Request) (*jwt.Token, error) {
	tokenStr := ExtractJWT(r)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(helpers.GetEnvVariable("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValidation(r *http.Request) error {
	token, err := CheckToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMeta(r *http.Request) (*AccessDetails, error) {
	token, err := CheckToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userEmail, ok := claims["user_email"].(string)
		if !ok {
			return nil, err
		}
		return &AccessDetails{
			UserEmail: userEmail,
		}, nil
	}
	return nil, err
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValidation(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token error",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
