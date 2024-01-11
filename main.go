package main

import (
	"context"
	protosava "github.com/MihajloJankovic/Aviability-Service/protos/main"
	"github.com/MihajloJankovic/reservation-service/handlers"
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		reservationRepo.Disconnect(ctx)

	}(reservationRepo, timeoutContext)
	reservationRepo.CreateTables(context.Background())

	//Initialize the handler and inject said logger
	connAva, err := grpc.Dial("avaibility-service:9095", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(connAva)

	ccava := protosava.NewAccommodationAviabilityClient(connAva)
	service := handlers.NewServer(logger, reservationRepo, ccava)

	protos.RegisterReservationServer(serverRegister, service)
	err = serverRegister.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
