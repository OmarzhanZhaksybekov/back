package repository

import (
	"github.com/ShawaDev/Dealer/pkg/model"
	"github.com/jmoiron/sqlx"
)

type CarsPostgres struct {
	db *sqlx.DB
}

func NewCarsPostgres(db *sqlx.DB) *CarsPostgres {
	return &CarsPostgres{
		db: db,
	}
}

func (r *CarsPostgres) AddCar(car model.Car) error {
	query := "INSERT INTO cars (brand, model, image_url, year, price, color) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := r.db.Exec(query, car.Brand, car.Model, car.Image_url, car.Year, car.Price, car.Color)

	return err
}

func (r *CarsPostgres) GetCarById(id int) (model.Car, error) {
	query := "SELECT * FROM cars WHERE id = $1"

	var car model.Car

	err := r.db.Get(&car, query, id)

	return car, err
}

func (r *CarsPostgres) EditCar(car model.Car, id int) error {
	query := "UPDATE cars SET brand = $1, model = $2, image_url = $3, year = $4, price = $5, color = $6 WHERE id = $7"

	_, err := r.db.Exec(query, car.Brand, car.Model, car.Image_url, car.Year, car.Price, car.Color, id)

	return err
}

func (r *CarsPostgres) DeleteCar(id int) error {
	query := "DELETE FROM cars WHERE id = $1"
	sorting := "UPDATE cars SET id = id -1 WHERE id > $1"

	if _, err := r.db.Exec(query, id); err != nil {
		return err
	}

	_, err := r.db.Exec(sorting, id)

	return err
}

func (r *CarsPostgres) GetAllCars() ([]model.Car, error) {
	query := "SELECT id, brand, model, image_url, year, price, color FROM cars"

	rows, err := r.db.Query(query)
	if err != nil {
		return []model.Car{}, err
	}

	defer rows.Close()

	var cars []model.Car
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(&car.Id, &car.Brand, &car.Model, &car.Image_url, &car.Year, &car.Price, &car.Color); err != nil {
			return []model.Car{}, err
		}

		cars = append(cars, car)
	}

	if err := rows.Err(); err != nil {
		return []model.Car{}, err
	}

	return cars, nil
}
