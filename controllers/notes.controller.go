package controllers

import (
	"context"
	"github.com/AloisCRR/jwt-api-notes/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

/*func error(err any.Any) {
	log.Printf("Error getting notes, value: %v", err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"status":  http.StatusInternalServerError,
		"message": "An error occurred",
	})
}*/

type Notes models.Notes

// Db instance
var notesCollection *mongo.Collection

func NotesCollection(c *mongo.Database) {
	notesCollection = c.Collection("notes")
}

// Route's controllers
func AllNotes(c *gin.Context) {
	notes := []Notes{}
	cursor, err := notesCollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error getting notes, %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "An error occurred",
		})
		return
	}

	for cursor.Next(context.TODO()) {
		var note Notes
		cursor.Decode(&note)
		notes = append(notes, note)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All notes",
		"data":    notes,
	})
	return
}

func GetNote(c *gin.Context) { //TODO input validation
	noteID := c.Param("id")

	docID, _ := primitive.ObjectIDFromHex(noteID)

	note := Notes{}

	err := notesCollection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&note)

	if err != nil {
		log.Printf("Error getting notes, %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "An error occurred",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Note",
		"data":    note,
	})
	return
}

func CreateNote(c *gin.Context) {
	var note Notes
	c.BindJSON(&note)

	title := note.Title
	content := note.Content
	user, err := models.ExtractTokenMeta(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "An error occurred",
		})
		return
	}

	newNote := Notes{
		Title:     title,
		Content:   content,
		User:      user.UserEmail,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, errIn := notesCollection.InsertOne(context.TODO(), newNote)

	if errIn != nil {
		log.Printf("Error on inserting new note, %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "An error occurred",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Note created successfully!",
	})
	return
}

func EditNote(c *gin.Context) {
	noteID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(noteID)

	var note Notes

	c.BindJSON(&note)

	title := note.Title
	content := note.Content

	newData := bson.M{
		"$set": bson.M{
			"title":     title,
			"content":   content,
			"updatedat": time.Now(),
		},
	}

	_, err := notesCollection.UpdateOne(context.TODO(), bson.M{"_id": docID}, newData)

	if err != nil {
		log.Printf("Error on inserting new note, %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "An error occurred",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Note edited successfully!",
	})
	return
}

func DeleteNote(c *gin.Context) {
	noteID := c.Param("id")
	docID, _ := primitive.ObjectIDFromHex(noteID)

	_, err := notesCollection.DeleteOne(context.TODO(), bson.M{"_id": docID})

	if err != nil {
		log.Printf("Error getting notes, %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "An error occurred",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Note deleted successfully!",
	})
	return
}
