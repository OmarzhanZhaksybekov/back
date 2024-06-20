package service

import (
	"github.com/ShawaDev/Dealer/pkg/model"
	"github.com/ShawaDev/Dealer/pkg/repository"
)

type CarsService struct {
	repo repository.Cars
}

func NewCarsService(repo repository.Cars) *CarsService {
	return &CarsService{
		repo: repo,
	}
}

func (s *CarsService) AddCar(car model.Car) error {
	return s.repo.AddCar(car)
}

func (s *CarsService) GetCarById(id int) (model.Car, error) {
	return s.repo.GetCarById(id)
}

func (s *CarsService) EditCar(car model.Car, id int) error {
	return s.repo.EditCar(car, id)
}

func (s *CarsService) DeleteCar(id int) error {
	return s.repo.DeleteCar(id)
}

func (s *CarsService) GetAllCars() ([]model.Car, error) {
	cars, err := s.repo.GetAllCars()
	if err != nil {
		return []model.Car{}, err
	}

	return cars, nil
}
