package garage

import (
	"github.com/Pwera/Playground/src/main/go/snippets/di/helpers"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type CarManager struct {
	Repo   *CarRepository
	Logger *zap.Logger
}

func (m *CarManager) GetAll() ([]*Car, error) {
	cars, err := m.Repo.FindAll()

	if cars == nil {
		cars = []*Car{}
	} else {
		m.Logger.Error(err.Error())
	}
	return cars, err
}

func (m *CarManager) Get(id string) (car *Car, err error) {
	car, err = m.Repo.FindByID(id)

	if m.Repo.IsNotFoundErr(err) {
		return nil, helpers.NewErrNotFound("Car " + id + " does not exist")
	}

	if err != nil {
		m.Logger.Error(err.Error())
	}
	return car, err
}

func (m *CarManager) Create(car *Car) (*Car, error) {
	if err := ValidateCar(car); err != nil {
		return nil, err
	}

	car.ID = bson.NewObjectId().Hex()
	err := m.Repo.Insert(car)

	if m.Repo.IsAlreadyExistErr(err) {
		return m.Create(car)
	}
	if err != nil {
		m.Logger.Error(err.Error())
		return nil, err
	}

	return car, nil
}

func (m *CarManager) Update(id string, car *Car) (*Car, error) {
	if err := ValidateCar(car); err != nil {
		return nil, err
	}

	car.ID = id

	err := m.Repo.Update(car)

	if m.Repo.IsNotFoundErr(err) {
		return nil, helpers.NewErrNotFound("Car " + id + " does not exist")
	}

	if err != nil {
		m.Logger.Error(err.Error())
		return nil, err
	}
	return car, err
}

func (m *CarManager) Delete(id string) error {
	err := m.Repo.Delete(id)

	if m.Repo.IsNotFoundErr(err) {
		return nil
	}
	if err != nil {
		m.Logger.Error(err.Error())
	}
	return err
}
