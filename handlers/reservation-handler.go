package handlers

import (
	"context"
	"errors"
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"log"
	"time"
)

type MyReservationServer struct {
	protos.UnimplementedReservationServer
	logger *log.Logger
	repo   *ReservationRepo
}

func NewServer(l *log.Logger, r *ReservationRepo) *MyReservationServer {
	return &MyReservationServer{*new(protos.UnimplementedReservationServer), l, r}
}
func isValidDateFormat(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
func (s MyReservationServer) DeleteByAccomnendation(xtx context.Context,in *protos.DeleteRequest) (*protos.Emptyaa, error){
	out, err := s.repo.DeleteByAccomandation(xtx,in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	return out, nil
}
func (s MyReservationServer) GetReservation(_ context.Context, in *protos.ReservationRequest) (*protos.DummyLista, error) {
	if in.GetId() <= 0 {
		return nil, errors.New("Invalid input. Ensure a valid ID is provided.")
	}

	out, err := s.repo.GetById(in.GetId())
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	ss := new(protos.DummyLista)
	ss.Dummy = out
	return ss, nil
}
func (s MyReservationServer) GetAllReservation(_ context.Context, in *protos.Emptyaa) (*protos.DummyLista, error) {
	out, err := s.repo.GetAll()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	ss := new(protos.DummyLista)
	ss.Dummy = out
	return ss, nil
}
func (s MyReservationServer) SetReservation(_ context.Context, in *protos.ReservationResponse) (*protos.Emptyaa, error) {
	// Validate required fields
	if in.GetId() <= 0 || in.GetEmail() == "" || in.GetDateFrom() == "" || in.GetDateTo() == "" {
		return nil, errors.New("Invalid input. Ensure all required fields are provided.")
	}

	// Convert date strings to time.Time
	dateFrom, err := time.Parse("2006-01-02", in.GetDateFrom())
	if err != nil {
		return nil, errors.New("Invalid date format for DateFrom. Use yyyy-mm-dd.")
	}

	dateTo, err := time.Parse("2006-01-02", in.GetDateTo())
	if err != nil {
		return nil, errors.New("Invalid date format for DateTo. Use yyyy-mm-dd.")
	}

	// Validate date range
	if dateFrom.After(dateTo) {
		return nil, errors.New("Invalid date range. DateFrom must be before DateTo.")
	}

	// Additional validation for other fields if needed

	err = s.repo.Create(in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptyaa), nil
}

func (s MyReservationServer) UpdateReservation(_ context.Context, in *protos.ReservationResponse) (*protos.Emptyaa, error) {
	// Validate required fields
	if in.GetId() <= 0 || in.GetEmail() == "" || in.GetDateFrom() == "" || in.GetDateTo() == "" {
		return nil, errors.New("Invalid input. Ensure all required fields are provided.")
	}

	// Validate date format
	if !isValidDateFormat(in.GetDateFrom()) || !isValidDateFormat(in.GetDateTo()) {
		return nil, errors.New("Invalid date format. Use yyyy-mm-dd.")
	}

	// Validate date range
	dateFrom, err := time.Parse("2006-01-02", in.GetDateFrom())
	if err != nil {
		return nil, errors.New("Invalid date format for DateFrom. Use yyyy-mm-dd.")
	}

	dateTo, err := time.Parse("2006-01-02", in.GetDateTo())
	if err != nil {
		return nil, errors.New("Invalid date format for DateTo. Use yyyy-mm-dd.")
	}

	if dateFrom.After(dateTo) {
		return nil, errors.New("Invalid date range. DateFrom must be before DateTo.")
	}

	// Additional validation for other fields if needed

	err = s.repo.Update(in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptyaa), nil
}
