package garage

import (
	"gopkg.in/mgo.v2"
)

type CarRepository struct {
	Session *mgo.Session
}

func (repo *CarRepository) collection() *mgo.Collection {
	return repo.Session.DB("dingo_car_api").C("cars")
}

func (repo *CarRepository) FindAll() ([]*Car, error) {
	var cars []*Car
	err := repo.collection().Find(nil).All(&cars)
	return cars, err
}

func (repo *CarRepository) FindByID(id string) (*Car, error) {
	var car *Car
	err := repo.collection().FindId(id).One(&car)
	return car, err
}

func (repo *CarRepository) Insert(car *Car) error {
	return repo.collection().Insert(&car)
}

func (repo *CarRepository) Update(car *Car) error {
	return repo.collection().UpdateId(car.ID, car)
}

func (repo *CarRepository) Delete(id string) error {
	return repo.collection().RemoveId(id)
}

func (repo *CarRepository) IsNotFoundErr(err error) bool {
	return err == mgo.ErrNotFound
}

func (repo *CarRepository) IsAlreadyExistErr(err error) bool {
	return mgo.IsDup(err)
}
