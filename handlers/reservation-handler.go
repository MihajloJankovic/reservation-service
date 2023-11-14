package handlers

import (
	"context"
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"log"
)

type MyReservationServer struct {
	protos.UnimplementedReservationServer
	logger *log.Logger
	repo   *ReservationRepo
}

func NewServer(l *log.Logger, r *ReservationRepo) *MyReservationServer {
	return &MyReservationServer{*new(protos.UnimplementedReservationServer), l, r}
}

func (s MyReservationServer) GetReservation(_ context.Context, in *protos.ReservationRequest) (*protos.DummyLista, error) {
	out, err := s.repo.GetById(in.GetId())
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	ss := new(protos.DummyLista)
	ss.Dummy = out
	return ss, nil
}
func (s MyReservationServer) GetAllAccommodation(_ context.Context, in *protos.Emptyaa) (*protos.DummyLista, error) {
	out, err := s.repo.GetAll()
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	ss := new(protos.DummyLista)
	ss.Dummy = out
	return ss, nil
}
func (s MyReservationServer) SetAccommodation(_ context.Context, in *protos.ReservationResponse) (*protos.Emptyaa, error) {
	out := new(protos.ReservationResponse)
	out.Id = in.GetId()
	out.Email = in.GetEmail()
	out.DateFrom = in.GetDateFrom()
	out.DateTo = in.GetDateTo()
	out.PricePerson = in.GetPricePerson()
	out.PriceAcc = in.GetPriceAcc()
	out.NumberOfPeople = in.GetNumberOfPeople()

	err := s.repo.Create(out)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return new(protos.Emptyaa), nil
}
func (s MyReservationServer) UpdateAccommodation(_ context.Context, in *protos.ReservationResponse) (*protos.Emptyaa, error) {
	err := s.repo.Update(in)
	if err != nil {
		return nil, err
	}
	return new(protos.Emptyaa), nil
}
