package repository

import (
	"ColombiaCar/internal/models"
	"context"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type CarRepository interface {
	SearchCars(ctx context.Context, filters map[string]string) ([]models.Car, error)
}

type carRepository struct {
	db *pgxpool.Pool
}

func NewCarRepository(db *pgxpool.Pool) CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) SearchCars(ctx context.Context, filters map[string]string) ([]models.Car, error) {
	query := "SELECT id, type, transmission, fuel, model, brand FROM cars WHERE 1=1"
	args := []interface{}{}
	i := 1
	for k, v := range filters {
		query += " AND " + k + "=$" + strconv.Itoa(i)
		args = append(args, v)
		i++
	}
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		err := rows.Scan(&car.ID, &car.Type, &car.Transmission, &car.Fuel, &car.Model, &car.Brand)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}
