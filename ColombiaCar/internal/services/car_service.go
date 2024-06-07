package services

import (
	"ColombiaCar/internal/models"
	"ColombiaCar/internal/repository"
	"context"
)

type CarService interface {
	SearchCars(ctx context.Context, filters map[string]string) ([]models.Car, error)
}

type carService struct {
	carRepository repository.CarRepository
}

func NewCarService(carRepo repository.CarRepository) CarService {
	return &carService{carRepository: carRepo}
}

func (s *carService) SearchCars(ctx context.Context, filters map[string]string) ([]models.Car, error) {
	return s.carRepository.SearchCars(ctx, filters)
}
