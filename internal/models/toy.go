package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Toy struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price"`
	Quantity    int                `bson:"quantity" json:"quantity"`
	Thumbnail   string             `bson:"thumbnail" json:"thumbnail"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

func GetAllToys(db *mongo.Database) ([]Toy, error) {
	var toys []Toy
	cursor, err := db.Collection("toys").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &toys); err != nil {
		return nil, err
	}

	return toys, nil
}

func GetToyByID(db *mongo.Database, id primitive.ObjectID) (*Toy, error) {
	var toy Toy
	err := db.Collection("toys").FindOne(context.Background(), bson.M{"_id": id}).Decode(&toy)
	if err != nil {
		return nil, err
	}
	return &toy, nil
}

func (t *Toy) Create(db *mongo.Database) error {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	result, err := db.Collection("toys").InsertOne(context.Background(), t)
	if err != nil {
		return err
	}
	t.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (t *Toy) Update(db *mongo.Database) error {
	t.UpdatedAt = time.Now()
	_, err := db.Collection("toys").UpdateOne(
		context.Background(),
		bson.M{"_id": t.ID},
		bson.M{"$set": t},
	)
	return err
}

func DeleteToy(db *mongo.Database, id primitive.ObjectID) error {
	_, err := db.Collection("toys").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
