package models

import (
	"github.com/fallenstedt/gin-example/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"context"
	"time"
)

type User struct {
	ID        string `json:"user_id,omitempty"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
	Gender    string `json:"gender"`
	PhotoURL  string `json:"photo_url"`
	Time      int64  `json:"current_time"`
	Active    bool   `json:"active,omitempty"`
	UpdatedAt int64  `json:"updated_at,omitempty"`
}

func (h User) GetByID() (*User, error) {
	db.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"name": "Alex"}
	collection := db.GetCollection("my_collection")

	var result User
	err := collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return &result, nil
}


