package garage

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type CarRepository struct {
	Client *mongo.Client
}

func (repo *CarRepository) collection() *mongo.Collection {
	return repo.Client.Database("dingo_car_api").Collection("cars")
}

func (repo *CarRepository) FindAll() ([]*Car, error) {
	var cars []*Car
	_, err := repo.collection().Find(nil, &cars)

	return cars, err
}

func (repo *CarRepository) FindByID(id string) (*Car, error) {
	var car *Car
	_, err := repo.collection().Find(nil, &car, id)
	return car, err
}

func (repo *CarRepository) Insert(car *Car) error {
	_, err := repo.collection().InsertOne(nil, &car)
	return err
}

func (repo *CarRepository) Update(car *Car) error {
	return repo.collection().UpdateByID(nil, car, car.ID)
}

func (repo *CarRepository) Delete(id string) error {
	return nil
}

func (repo *CarRepository) IsNotFoundErr(err error) bool {
	//return err == mgo.ErrNotFound
	return false
}

func (repo *CarRepository) IsAlreadyExistErr(err error) bool {
	//return mgo.IsDup(err)
	return false
}
