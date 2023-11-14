package handlers

import (
	"context"
	"fmt"
	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type ReservationRepo struct {
	logger *log.Logger
	cli    *mongo.Client
}

// New NoSQL: Constructor which reads db configuration from environment
func New(ctx context.Context, logger *log.Logger) (*ReservationRepo, error) {
	dburi := os.Getenv("MONGO_DB_URI")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		return nil, err
	}

	return &ReservationRepo{
		cli:    client,
		logger: logger,
	}, nil
}

// Disconnect from database
func (rr *ReservationRepo) Disconnect(ctx context.Context) error {
	err := ar.cli.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

// Ping Check database connection
func (rr *ReservationRepo) Ping() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check connection -> if no error, connection is established
	err := ar.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		ar.logger.Println(err)
	}
	// Print available databases
	databases, err := ar.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		ar.logger.Println(err)
	}
	fmt.Println(databases)
}
func (rr *ReservationRepo) GetAll() ([]*protos.ReservationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationCollection := ar.getCollection()
	var accommodationsSlice []*protos.ReservationResponse

	accommodationCursor, err := accommodationCollection.Find(ctx, bson.M{})
	if err != nil {
		ar.logger.Println(err)
		return nil, err
	}
	if err = accommodationCursor.All(ctx, &accommodationsSlice); err != nil {
		ar.logger.Println(err)
		return nil, err
	}
	return accommodationsSlice, nil
}
func (rr *ReservationRepo) GetById(email string) ([]*protos.AccommodationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	accommodationCollection := ar.getCollection()
	var accommodationsSlice []*protos.AccommodationResponse

	// Assuming you have a filter based on the email, modify the filter as needed
	filter := bson.M{"email": email}

	accommodationCursor, err := accommodationCollection.Find(ctx, filter)
	if err != nil {
		ar.logger.Println(err)
		return nil, err
	}
	defer func(accommodationCursor *mongo.Cursor, ctx context.Context) {
		err := accommodationCursor.Close(ctx)
		if err != nil {
			ar.logger.Println(err)
		}
	}(accommodationCursor, ctx)

	for accommodationCursor.Next(ctx) {
		var accommodation protos.AccommodationResponse
		if err := accommodationCursor.Decode(&accommodation); err != nil {
			ar.logger.Println(err)
			return nil, err
		}
		accommodationsSlice = append(accommodationsSlice, &accommodation)
	}

	if err := accommodationCursor.Err(); err != nil {
		ar.logger.Println(err)
		return nil, err
	}

	return accommodationsSlice, nil
}
func (ar *ReservationRepo) Create(profile *protos.AccommodationResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accommodationCollection := ar.getCollection()

	result, err := accommodationCollection.InsertOne(ctx, &profile)
	if err != nil {
		ar.logger.Println(err)
		return err
	}
	ar.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (rr *ReservationRepo) Update(reservation *protos.ReservationResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationCollection := rr.getCollection()

	filter := bson.M{"email": reservation.GetEmail()}
	update := bson.M{"$set": bson.M{
		"id":       reservation.GetName(),
		"price":    reservation.GetPrice(),
		"location": reservation.GetLocation(),
		"adress":   reservation.GetAdress(),
	}}
	result, err := reservationCollection.UpdateOne(ctx, filter, update)
	rr.logger.Printf("Documents matched: %v\n", result.MatchedCount)
	rr.logger.Printf("Documents updated: %v\n", result.ModifiedCount)

	if err != nil {
		rr.logger.Println(err)
		return err
	}
	return nil
}

func (rr *ReservationRepo) getCollection() *mongo.Collection {
	accommodationDatabase := ar.cli.Database("mongoAccommodation")
	accommodationCollection := accommodationDatabase.Collection("accommodations")
	return accommodationCollection
}
