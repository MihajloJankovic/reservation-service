package handlers

import (
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"google.golang.org/grpc"
	"log"
)

type MyReservationServer struct{
	protos.UnimplementedAccommodationServer
	logger *log.logger
	repo *ReservationRepo
}

func NewServer(l *log.Logger, r *ReservationRepo) *MyReservationServer {
	return &MyReservationServer{*new(protos.UnimplementedAccommodationServer), l, r}
}

func (s MyReservationServer) GetReservation(ctx context.Context, in *ReservationRequest) (*DummyLista, error){
	out, err := s.repo.GetById(in.GetEmail())
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	ss := new(protos.DummyList)
	ss.Dummy = out
	return ss, nil
}
GetAllReservations(ctx context.Context, in *Emptyaa, opts ...grpc.CallOption) (*DummyLista, error)
SetReservation(ctx context.Context, in *ReservationResponse, opts ...grpc.CallOption) (*Emptyaa, error)
UpdateReservation(ctx context.Context, in *ReservationResponse, opts ...grpc.CallOption) (*Emptyaa, error)