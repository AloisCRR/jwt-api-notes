package routes

import (
	"github.com/AloisCRR/jwt-api-notes/controllers"
	"github.com/AloisCRR/jwt-api-notes/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Notes API made with MongoDB & Golang",
	})
	return
}

func noRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  http.StatusOK,
		"message": "No route defined",
	})
}

func Routes(router *gin.Engine) {
	router.GET("/", home)
	router.GET("/notes", models.TokenAuthMiddleware(), controllers.AllNotes)
	router.GET("/notes/:id", models.TokenAuthMiddleware(), controllers.GetNote)
	router.POST("/notes", models.TokenAuthMiddleware(), controllers.CreateNote)
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.PUT("/notes/:id", models.TokenAuthMiddleware(), controllers.EditNote)
	router.DELETE("/notes/:id", models.TokenAuthMiddleware(), controllers.DeleteNote)
	router.NoRoute(noRoute)
}
