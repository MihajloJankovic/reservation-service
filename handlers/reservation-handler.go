package handlers

import (
	"context"
	"errors"
	protosava "github.com/MihajloJankovic/Aviability-Service/protos/main"
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"log"
	"strings"
	"time"
)

type MyReservationServer struct {
	protos.UnimplementedReservationServer
	logger *log.Logger
	repo   *ReservationRepo
	cc     protosava.AccommodationAviabilityClient
}

func NewServer(l *log.Logger, r *ReservationRepo, cc protosava.AccommodationAviabilityClient) *MyReservationServer {
	return &MyReservationServer{*new(protos.UnimplementedReservationServer), l, r, cc}
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

func (s MyReservationServer) GetAllReservationsByEmail(ctx context.Context, in *protos.Emaill) (*protos.DummyLista, error) {
	out, err := s.repo.GetReservationsByEmail(ctx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}
func (s MyReservationServer) DeleteReservationById(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	out, err := s.repo.DeleteReservationById(ctx, in)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return out, nil
}

func (s MyReservationServer) CheckActiveReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	_, err := s.repo.CheckActiveReservationByEmail(ctx, in)
	if err == nil {
		return nil, errors.New("there is active reservation")
	}
	return new(protos.Emptyaa), nil
}
func (s MyReservationServer) CheckActiveReservation(ctx context.Context, in *protos.DateFromDateTo) (*protos.Emptyaa, error) {

	_, err := s.repo.CheckActiveReservation(ctx, in)
	if err == nil {
		return nil, errors.New("there is active reservation")
	}
	return new(protos.Emptyaa), nil
}
func (s MyReservationServer) DeleteByAccomnendation(xtx context.Context, in *protos.DeleteRequestaa) (*protos.Emptyaa, error) {
	out, err := s.repo.DeleteByAccommodation(xtx, in)
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
	parsedTime1, err := time.Parse(time.RFC3339, trimmedReservation.GetDateFrom())
	if err != nil {
		return nil, errors.New("Invalid date format for DateFrom. Use yyyy-mm-dd.")
	}
	parsedTime2, err := time.Parse(time.RFC3339, trimmedReservation.GetDateTo())
	if err != nil {
		return nil, errors.New("Invalid date format for DateTo. Use yyyy-mm-dd.")
	}
	formattedDate1 := parsedTime1.Format("2006-01-02")
	formattedDate2 := parsedTime2.Format("2006-01-02")
	// Convert date strings to time.Time
	dateFrom := formattedDate1
	dateTo := formattedDate2
	trimmedReservation.DateFrom = dateFrom
	trimmedReservation.DateTo = dateTo
	parsedTime3, _ := time.Parse(time.RFC3339, formattedDate1)

	parsedTime4, _ := time.Parse(time.RFC3339, formattedDate2)
	// Validate date range
	if parsedTime3.After(parsedTime4) {
		return nil, errors.New("Invalid date range. DateFrom must be before DateTo.")
	}

	temp := new(protosava.CheckRequest)
	temp.Id = in.GetAccid()
	temp.From = dateFrom
	temp.To = dateTo

	_, err = s.cc.GetAccommodationCheck(context.Background(), temp)
	if err != nil {
		log.Printf("RPC failed: %v\n", err)
		return nil, err
	}
	// Additional validation for other fields if needed
	err = s.repo.CheckIfAvailable(trimmedReservation)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
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
