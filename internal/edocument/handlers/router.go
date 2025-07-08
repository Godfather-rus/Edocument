package handlers

import "context"

type Repository interface {
	CreateEdoc(ctx context.Context, file any) error
	GetEdoc()
	GetEdocsList()
	//CreateBooking(ctx context.Context, booking *domain.Booking) (*domain.Booking, error)
	//ListBookings(ctx context.Context, workshopID int64) ([]*domain.Booking, error)
}

type Handlers struct {
	repo Repository
}

func NewHandlers(repo Repository) *Handlers {
	return &Handlers{
		repo: repo,
	}
}
