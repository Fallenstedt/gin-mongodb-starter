package models

import (
	"github.com/fallenstedt/gin-example/src/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"context"
	"time"
	"errors"
)

type UserRepository interface {
	AddUser() (*mongo.InsertOneResult, error)
	GetById() error
}

type User struct {
	ID 		  primitive.ObjectID `bson:"_id" json:"id"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday"`
}

func NewUser(c *gin.Context, id string) (User, error) {
	var us User
	err := c.BindJSON(&us)
	if err != nil {
		log.Printf("Unable to parse body %f", err)
		return User{}, err
	}

	if len(id) > 0 {
		docID, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			log.Printf("Failed to create object id from hex %v", id)
			return User{}, err
		}
		us.ID = docID
	}

	if len(us.Name) == 0 {
		return User{}, errors.New("name is required")
	}

	if len(us.Birthday) == 0 {
		return User{}, errors.New("birthday is required")
	}

	return us, nil
}

func (h User) AddUser() ( *mongo.InsertOneResult, error) {
	db.Init()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	collection := db.GetCollection("my_collection")

	res, err := collection.InsertOne(ctx, bson.M{"name": h.Name, "birthday": h.Birthday})

	if err != nil {
		log.Printf("Failed to insert user into DB %f", err)
		return nil, err
	}

	log.Printf("Successfully inserted user %f", res.InsertedID)
	return res, nil
}


// TODO create error types. This can create a 500 or a 404!. Maybe we can create a 404 error type for the controller to check?
func (h User) GetByID(id string) (*User, error) {
	db.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection("my_collection")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Failed to create object id from hex %v", id)
		return nil, err
	}

	var result User
	err = collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&result)

	if err != nil {
		log.Printf("Unable find user by id %f", err)
		return nil, err
	}

	return &result, nil
}


func (h User) DeleteByID(id string) (int64, error) {
	db.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection("my_collection")

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Failed to create object id from hex %v", id)
		return 0, err
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": docID})

	if err != nil {
		log.Printf("Failed to delete id %v %v", id, err)
		return 0, err
	}


	return res.DeletedCount, nil
}

func (h User) UpdateUser() (*mongo.UpdateResult, error) {
	db.Init()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.GetCollection("my_collection")

	res, err := collection.UpdateOne(ctx, bson.M{"_id": h.ID}, bson.M{
		"$set": bson.M{"name": h.Name, "birthday": h.Birthday},
	})

	if err != nil {
		log.Printf("Failed to update id %v %v", h.ID, err)
		return nil, err
	}

	return res, err
}

