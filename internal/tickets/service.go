package tickets

import (
	"context"
	"desafio-goweb-nadiaMartinMontesi/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error)
	GetByCountry(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error) {
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (s *service) GetTicketByDestination(ctx context.Context, destination string) ([]domain.Ticket, error) {
	ticketsDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return nil, err
	}
	return ticketsDestination, nil
}

func (s *service) GetByCountry(ctx context.Context, destination string) (int, error) {
	ticketsDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	countTicket := len(ticketsDestination)
	return countTicket, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (int, error) {
	totalTickets, err := s.repository.GetAll(ctx)
	ticketsDestination, err := s.GetByCountry(ctx, destination)
	if err != nil {
		return 0, err
	}

	avgDestination := ticketsDestination * 100 / len(totalTickets)
	return avgDestination, nil
}
