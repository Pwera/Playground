package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	user     string = "example"
	password string = "root"
	host     string = "localhost"
	port     string = "27017"
	name     string = "sampledb"
)

var (
	db *mongo.Database
)
func GetDB() (db *mongo.Database, err error) {
	if db != nil {
		return
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", user, password, host, port, name)
	fmt.Println("uri: {}", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return
	}

	fmt.Println("Connected to MongoDB")
	db = client.Database(name)
	return
}
