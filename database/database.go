package database

import (
	"context"
	"log"
	"time"

	"github.com/aleksandarmilanovic/gqlgen-todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB ...
type DB struct {
	client *mongo.Client
}
// Connect ...
func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://alexandar12:Acoaco77@cluster0.68jgw.gcp.mongodb.net/<dbname>?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

// Save ...
func (db *DB) Save(input *model.NewLion) *model.Lion {
	collection := db.client.Database("test").Collection("lions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Lion{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
		Name: input.Name,
		IsKing: input.IsKing,
	}
}

// GetByID ...
func(db *DB) GetByID(ID string) *model.Lion {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection := db.client.Database("test").Collection("lions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"_id": ObjectID})
	lion := model.Lion{}
	res.Decode(&lion)
	return &lion
}

// GetAllLions ...
func (db *DB) GetAllLions() []*model.Lion {
	collection := db.client.Database("test").Collection("lions")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var lions []*model.Lion
	for cur.Next(ctx) {
		var lion *model.Lion
		err := cur.Decode(&lion)
		if err != nil {
			log.Fatal(err)
		}
		lions = append(lions, lion)
	}
	return lions
}