package services

import (
	"ColombiaCar/internal/models"
	"ColombiaCar/internal/repository"
	"context"
)

type ReservationService interface {
	CreateReservation(ctx context.Context, reservation *models.Reservation) error
	GetUserReservations(ctx context.Context, userID int) ([]models.Reservation, error)
}

type reservationService struct {
	reservationRepository repository.ReservationRepository
}

func NewReservationService(reservationRepo repository.ReservationRepository) ReservationService {
	return &reservationService{reservationRepository: reservationRepo}
}

func (s *reservationService) CreateReservation(ctx context.Context, reservation *models.Reservation) error {
	return s.reservationRepository.CreateReservation(ctx, reservation)
}

func (s *reservationService) GetUserReservations(ctx context.Context, userID int) ([]models.Reservation, error) {
	return s.reservationRepository.GetUserReservations(ctx, userID)
}
