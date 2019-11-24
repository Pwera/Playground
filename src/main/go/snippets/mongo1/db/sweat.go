package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const SWEAT_TABLE string = "sweat"

type Sweat struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"userid,omitempty" json:"userid"`
	CreatedAt time.Time          `bson:created_at`
}

func (s *Sweat) Create() (err error) {
	db, err := GetDB()
	if err != nil {
		return
	}
	s.CreatedAt = time.Now()
	collection := db.Collection(SWEAT_TABLE)
	_ , err = collection.InsertOne(context.TODO(), s)
	if err != nil{
		fmt.Println("Error inserting sweat: %v \n%s", s, err)
		return
	}
	fmt.Println("Inserted sweat into collection")
	return
}

func (s *Sweat) Delete() (err error) {
	return
}
