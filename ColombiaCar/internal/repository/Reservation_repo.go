package repository

import (
	"ColombiaCar/internal/models"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ReservationRepository interface {
	CreateReservation(ctx context.Context, reservation *models.Reservation) error
	GetUserReservations(ctx context.Context, userID int) ([]models.Reservation, error)
}

type reservationRepository struct {
	db *pgxpool.Pool
}

func NewReservationRepository(db *pgxpool.Pool) ReservationRepository {
	return &reservationRepository{db: db}
}

func (r *reservationRepository) CreateReservation(ctx context.Context, reservation *models.Reservation) error {
	_, err := r.db.Exec(ctx, "INSERT INTO reservations (user_id, car_id) VALUES ($1, $2)", reservation.UserID, reservation.CarID)
	return err
}

func (r *reservationRepository) GetUserReservations(ctx context.Context, userID int) ([]models.Reservation, error) {
	rows, err := r.db.Query(ctx, "SELECT id, user_id, car_id FROM reservations WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []models.Reservation
	for rows.Next() {
		var reservation models.Reservation
		err := rows.Scan(&reservation.ID, &reservation.UserID, &reservation.CarID)
		if err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}
