package main

import (
	"context"
	"github.com/MihajloJankovic/reservation-service/handlers"
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	lis, err := net.Listen("tcp", ":9096")
	if err != nil {
		log.Fatal(err)
	}
	serverRegister := grpc.NewServer()

	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	logger := log.New(os.Stdout, "[reservation-genfiles] ", log.LstdFlags)
	reservationLog := log.New(os.Stdout, "[reservation-repo-log] ", log.LstdFlags)

	reservationRepo, err := handlers.New(timeoutContext, reservationLog)
	if err != nil {
		logger.Fatal(err)
	}
	defer func(reservationRepo *handlers.ReservationRepo, ctx context.Context) {
		err := reservationRepo.Disconnect(ctx)
		if err != nil {

		}
	}(reservationRepo, timeoutContext)

	// NoSQL: Checking if the connection was established
	reservationRepo.Ping()

	//Initialize the handler and inject said logger
	service := handlers.NewServer(logger, reservationRepo)

	protos.RegisterReservationServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
