package repository

import (
	"github.com/ShawaDev/Dealer/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Cars interface {
	AddCar(car model.Car) error
	GetCarById(id int) (model.Car, error)
	EditCar(car model.Car, id int) error
	DeleteCar(id int) error
	GetAllCars() ([]model.Car, error)
}

type Repository struct {
	Cars
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Cars: NewCarsPostgres(db),
	}
}
