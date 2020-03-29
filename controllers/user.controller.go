package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"projects/mongodb-notes-api/models"
)

type Users models.Users

// DB instance
var usersCollection *mongo.Collection

func UsersCollection(c *mongo.Database) {
	usersCollection = c.Collection("users")
}

// Route's controllers
func SignUp(c *gin.Context) {
	var user Users

	/*if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}*/

	c.BindJSON(&user)

	v := validator.New() // TODO fuente: https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835

	if err := v.Struct(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	email := user.Email //strings.Trim(user.Email, "\t "+"\n "+" ")
	pass, _ := models.HashPassword(user.Password)

	newUser := Users{
		Email:    email,
		Password: pass,
	}

	_, err := usersCollection.InsertOne(context.TODO(), newUser)

	if err != nil {
		log.Printf("Error creating new user, %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "An error occurred",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Successfully registered! ",
	})
	return
}

// TODO login fuente: https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
func Login(c *gin.Context) {

	var user Users
	c.BindJSON(&user)

	v := validator.New()

	if err := v.Struct(user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	usrLogin := Users{}
	err := usersCollection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&usrLogin)

	if err != nil || !models.CheckPasswordHash(user.Password, usrLogin.Password) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Email or password error. Try again!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login completed!",
	})
	return
}
