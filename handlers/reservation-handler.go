package handlers

import (
	"context"
	"errors"
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"log"
	"strings"
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

func trimReservationFields(reservation *protos.ReservationResponse) *protos.ReservationResponse {
	return &protos.ReservationResponse{
		Id:       reservation.GetId(),
		Email:    strings.TrimSpace(reservation.GetEmail()),
		DateFrom: strings.TrimSpace(reservation.GetDateFrom()),
		DateTo:   strings.TrimSpace(reservation.GetDateTo()),
		Accid:    strings.TrimSpace(reservation.GetAccid()),
		// Add trimming for other fields here
	}
}
func (s MyReservationServer) DeleteReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	out, err := s.repo.DeleteReservationByEmail(ctx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}

func (s MyReservationServer) CheckActiveReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	_, err := s.repo.CheckActiveReservationByEmail(ctx, in)
	if err == nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptyaa), nil
}
func (s MyReservationServer) CheckActiveReservation(ctx context.Context, in *protos.DateFromDateTo) (*protos.Emptyaa, error) {
	_, err := s.repo.CheckActiveReservation(ctx, in)
	if err == nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptyaa), nil
}
func (s MyReservationServer) DeleteByAccomnendation(xtx context.Context, in *protos.DeleteRequestaa) (*protos.Emptyaa, error) {
	out, err := s.repo.DeleteByAccomandation(xtx, in)
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
	// Trim whitespaces from all fields
	trimmedReservation := trimReservationFields(in)

	// Validate required fields
	if trimmedReservation.GetEmail() == "" || trimmedReservation.GetDateFrom() == "" || trimmedReservation.GetDateTo() == "" {
		return nil, errors.New("Invalid input. Ensure all required fields are provided.")
	}

	// Convert date strings to time.Time
	dateFrom, err := time.Parse("2006-01-02", trimmedReservation.GetDateFrom())
	if err != nil {
		return nil, errors.New("Invalid date format for DateFrom. Use yyyy-mm-dd.")
	}

	dateTo, err := time.Parse("2006-01-02", trimmedReservation.GetDateTo())
	if err != nil {
		return nil, errors.New("Invalid date format for DateTo. Use yyyy-mm-dd.")
	}

	// Validate date range
	if dateFrom.After(dateTo) {
		return nil, errors.New("Invalid date range. DateFrom must be before DateTo.")
	}

	// Additional validation for other fields if needed

	err = s.repo.Create(trimmedReservation)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptyaa), nil
}

func (s MyReservationServer) UpdateReservation(_ context.Context, in *protos.ReservationResponse) (*protos.Emptyaa, error) {
	// Trim whitespaces from all fields
	trimmedReservation := trimReservationFields(in)

	// Validate required fields
	if trimmedReservation.GetId() <= 0 || trimmedReservation.GetEmail() == "" || trimmedReservation.GetDateFrom() == "" || trimmedReservation.GetDateTo() == "" {
		return nil, errors.New("Invalid input. Ensure all required fields are provided.")
	}

	// Validate date format
	if !isValidDateFormat(trimmedReservation.GetDateFrom()) || !isValidDateFormat(trimmedReservation.GetDateTo()) {
		return nil, errors.New("Invalid date format. Use yyyy-mm-dd.")
	}

	// Validate date range
	dateFrom, err := time.Parse("2006-01-02", trimmedReservation.GetDateFrom())
	if err != nil {
		return nil, errors.New("Invalid date format for DateFrom. Use yyyy-mm-dd.")
	}

	dateTo, err := time.Parse("2006-01-02", trimmedReservation.GetDateTo())
	if err != nil {
		return nil, errors.New("Invalid date format for DateTo. Use yyyy-mm-dd.")
	}

	if dateFrom.After(dateTo) {
		return nil, errors.New("Invalid date range. DateFrom must be before DateTo.")
	}

	// Additional validation for other fields if needed

	err = s.repo.Update(trimmedReservation)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptyaa), nil
}
