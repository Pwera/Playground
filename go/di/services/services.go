package services

import (
	"context"
	"fmt"
	"github.com/pwera/di/garage"
	"github.com/pwera/di/logging"
	"github.com/sarulabs/di"
	mongo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var Services = []di.Def{
	{
		Name:  "logger",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return logging.Logger, nil
		},
	}, {
		Name:  "mongo-pool",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:example@localhost:27017/test"))
			//url, err := mgo.ParseURL("mongodb://root:example@localhost:27017/test")
			//fmt.Println(url)
			fmt.Println(err)
			//return mgo.DialWithTimeout(os.Getenv("MONGO_URL"), 5*time.Second)
			//return mgo.DialWithTimeout("mongodb://root:example@localhost:27017/", 5*time.Second)
			//return mgo.DialWithTimeout("mongodb://root:example@localhost:27017/test", 5*time.Second)
			return client, err
		},

		Close: func(obj interface{}) error {
			//obj.(*mongo.Client).Close()
			return nil
		},
	}, {
		Name:  "mongo",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			session, err := ctn.Get("mongo-pool").(*mongo.Client).StartSession()
			return &session, err
		},
		Close: func(obj interface{}) error {
			//obj.(*mgo.Session).Close()
			return nil
		},
	}, {
		Name:  "car-repository",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &garage.CarRepository{
				Session: ctn.Get("mongo").(*mgo.Session),
			}, nil
		},
	},
	{
		Name:  "car-manager",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return &garage.CarManager{
				Repo:   ctn.Get("car-repository").(*garage.CarRepository),
				Logger: ctn.Get("logger").(*zap.Logger),
			}, nil
		},
	},
}
