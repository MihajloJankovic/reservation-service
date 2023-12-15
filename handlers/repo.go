package handlers

import (
	"context"
	"errors"
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
	err := rr.cli.Disconnect(ctx)
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
	err := rr.cli.Ping(ctx, readpref.Primary())
	if err != nil {
		rr.logger.Println(err)
	}
	// Print available databases
	databases, err := rr.cli.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		rr.logger.Println(err)
	}
	fmt.Println(databases)
}
func (rr *ReservationRepo) GetAll() ([]*protos.ReservationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationCollection := rr.getCollection()
	var reservationsSlice []*protos.ReservationResponse

	reservationCursor, err := reservationCollection.Find(ctx, bson.M{})
	if err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	if err = reservationCursor.All(ctx, &reservationsSlice); err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	return reservationsSlice, nil
}
func (pr *ReservationRepo) DeleteByAccomandation(xtx context.Context, in *protos.DeleteRequestaa) (*protos.Emptyaa, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resCollection := pr.getCollection()

	filter := bson.M{"accid": in.GetUid()}
	_, err := resCollection.DeleteMany(ctx, filter)
	if err != nil {
		return nil, errors.New("Coundnt delete")
	}

	return new(protos.Emptyaa), nil
}
func (rr *ReservationRepo) GetById(id int32) ([]*protos.ReservationResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationCollection := rr.getCollection()
	var reservationsSlice []*protos.ReservationResponse

	filter := bson.M{"id": id}

	reservationCursor, err := reservationCollection.Find(ctx, filter)
	if err != nil {
		rr.logger.Println(err)
		return nil, err
	}
	defer func(reservationCursor *mongo.Cursor, ctx context.Context) {
		err := reservationCursor.Close(ctx)
		if err != nil {
			rr.logger.Println(err)
		}
	}(reservationCursor, ctx)

	for reservationCursor.Next(ctx) {
		var reservation protos.ReservationResponse
		if err := reservationCursor.Decode(&reservation); err != nil {
			rr.logger.Println(err)
			return nil, err
		}
		reservationsSlice = append(reservationsSlice, &reservation)
	}

	if err := reservationCursor.Err(); err != nil {
		rr.logger.Println(err)
		return nil, err
	}

	return reservationsSlice, nil
}
func (rr *ReservationRepo) Create(profile *protos.ReservationResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	accommodationCollection := rr.getCollection()

	result, err := accommodationCollection.InsertOne(ctx, &profile)
	if err != nil {
		rr.logger.Println(err)
		return err
	}
	rr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}
func (rr *ReservationRepo) CheckIfAvaible(profile *protos.ReservationResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationCollection := rr.getCollection()

	filter := bson.D{
		{"datefrom", bson.D{
			{"$gte", profile.DateFrom},
			{"$lte", profile.DateTo},
		}},
		{"dateto", bson.D{
			{"$gte", profile.DateFrom},
			{"$lte", profile.DateTo},
		}},
	}

	reservationCursor, err := reservationCollection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	defer reservationCursor.Close(ctx)

	for reservationCursor.Next(ctx) {
		var reservation protos.ReservationResponse
		if err := reservationCursor.Decode(&reservation); err != nil {
			rr.logger.Println(err)
			return err
		}
		return errors.New("there is active reservation for this date range")

	}

	return nil
}
func (rr *ReservationRepo) GetReservationsByEmail(ctx context.Context, in *protos.Emaill) (*protos.DummyLista, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationCollection := rr.getCollection()

	filter := bson.M{"email": in.GetEmail()}

	reservationCursor, err := reservationCollection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer reservationCursor.Close(ctx)

	var lista protos.DummyLista
	for reservationCursor.Next(ctx) {
		var reservation protos.ReservationResponse
		if err := reservationCursor.Decode(&reservation); err != nil {
			rr.logger.Println(err)
			return nil, err
		}
		lista.Dummy = append(lista.Dummy, &reservation)
	}

	if err := reservationCursor.Err(); err != nil {
		rr.logger.Println(err)
		return nil, err
	}

	return &lista, nil
}

func (rr *ReservationRepo) DeleteReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resCollection := rr.getCollection()

	filter := bson.M{"email": in.GetEmail()}
	_, err := resCollection.DeleteMany(ctx, filter)
	if err != nil {
		return nil, errors.New("Couldn't delete")
	}

	return new(protos.Emptyaa), nil
}
func (rr *ReservationRepo) DeleteReservationById(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resCollection := rr.getCollection()

	filter := bson.M{"id": in.GetEmail()}
	_, err := resCollection.DeleteMany(ctx, filter)
	if err != nil {
		return nil, errors.New("Couldn't delete")
	}

	return new(protos.Emptyaa), nil
}

func (rr *ReservationRepo) CheckActiveReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationCollection := rr.getCollection()

	layout := "2006-01-02"
	filter := bson.D{
		{"datefrom", bson.D{
			{"$lte", time.Now().Format(layout)},
		}},
		{"email", in.GetEmail()},
		{"dateto", bson.D{
			{"$gte", time.Now().Format(layout)},
		}},
	}

	reservationCursor, err := reservationCollection.Find(ctx, filter)
	if err != nil {
		rr.logger.Println("Error querying MongoDB:", err)
		return nil, err
	}
	defer reservationCursor.Close(ctx)

	// Check if any active reservation was found
	if reservationCursor.Next(ctx) {
		var reservation protos.ReservationResponse
		if err := reservationCursor.Decode(&reservation); err != nil {
			rr.logger.Println("Error decoding result:", err)
			return nil, err
		}
		return new(protos.Emptyaa), nil
	}

	// No active reservation found
	return nil, errors.New("No active reservation found")
}
func (rr *ReservationRepo) CheckActiveReservation(ctx context.Context, in *protos.DateFromDateTo) (*protos.Emptyaa, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	reservationCollection := rr.getCollection()

	layout := "2006-01-02"
	filter := bson.D{
		{"datefrom", bson.D{
			{"$lte", time.Now().Format(layout)},
		}},
		{"accid", in.GetAccid()},
		{"dateto", bson.D{
			{"$gte", time.Now().Format(layout)},
		}},
	}

	reservationCursor, err := reservationCollection.Find(ctx, filter)
	if err != nil {
		rr.logger.Println("Error querying MongoDB:", err)
		return nil, err
	}
	defer reservationCursor.Close(ctx)

	// Check if any active reservation was found
	if reservationCursor.Next(ctx) {
		var reservation protos.ReservationResponse
		if err := reservationCursor.Decode(&reservation); err != nil {
			rr.logger.Println("Error decoding result:", err)
			return nil, err
		}
		return new(protos.Emptyaa), nil
	}

	// No active reservation found
	return nil, errors.New("No active reservation found")
}

func (rr *ReservationRepo) Update(reservation *protos.ReservationResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	reservationCollection := rr.getCollection()

	filter := bson.M{"id": reservation.GetId()}
	update := bson.M{"$set": bson.M{
		"dateFrom": reservation.GetDateFrom(),
		"dateTo":   reservation.GetDateTo(),
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
	accommodationDatabase := rr.cli.Database("mongoReservation")
	accommodationCollection := accommodationDatabase.Collection("reservations")
	return accommodationCollection
}
