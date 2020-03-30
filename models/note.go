package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notes struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User      string             `json:"user_id"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
