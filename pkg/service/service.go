package service

import (
	"github.com/ShawaDev/Dealer/pkg/model"
	"github.com/ShawaDev/Dealer/pkg/repository"
)

type Cars interface {
	AddCar(car model.Car) error
	GetCarById(id int) (model.Car, error)
	EditCar(car model.Car, id int) error
	DeleteCar(id int) error
	GetAllCars() ([]model.Car, error)
}

type Service struct {
	Cars
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Cars: NewCarsService(repo.Cars),
	}
}
