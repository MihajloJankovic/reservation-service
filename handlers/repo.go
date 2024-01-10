package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	protos "github.com/MihajloJankovic/reservation-service/protos/genfiles"
	"github.com/gocql/gocql"
)

type ReservationRepo struct {
	logger  *log.Logger
	session *gocql.Session
}

// New NoSQL: Constructor which reads db configuration from environment
func New(ctx context.Context, logger *log.Logger) (*ReservationRepo, error) {
	db := os.Getenv("CASS_DB")

	// Connect to default keyspace
	cluster := gocql.NewCluster(db)
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	if err != nil {
		logger.Println(err)
		return nil, err
	}
	defer session.Close()

	// Create 'reservation' keyspace
	err = session.Query(
		fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s
					WITH replication = {
						'class' : 'SimpleStrategy',
						'replication_factor' : %d
					}`, "reservation", 1)).Exec()
	if err != nil {
		logger.Println(err)
	}

	// Connect to reservation keyspace
	cluster.Keyspace = "reservation"
	cluster.Consistency = gocql.One
	session, err = cluster.CreateSession()
	if err != nil {
		logger.Println(err)
		return nil, err
	}

	// Create 'reservations' table
	err = session.Query(
		`CREATE TABLE IF NOT EXISTS reservations (
			id UUID PRIMARY KEY,
			accid UUID,
			email text,
			datefrom date,
			dateto date
		)`).Exec()
	if err != nil {
		logger.Println(err)
	}

	// Return repository with logger and DB session
	return &ReservationRepo{
		session: session,
		logger:  logger,
	}, nil
}

// Disconnect from database
func (rr *ReservationRepo) Disconnect(ctx context.Context) {
	rr.session.Close()
}

func (rr *ReservationRepo) GetAll() ([]*protos.ReservationResponse, error) {
	session := rr.session

	var reservationsSlice []*protos.ReservationResponse

	query := "SELECT id, accid, email, datefrom, dateto FROM reservations ALLOW FILTERING"
	iter := session.Query(query).Iter()

	var reservation protos.ReservationResponse
	for iter.Scan(
		&reservation.Id,
		&reservation.Accid,
		&reservation.Email,
		&reservation.DateFrom,
		&reservation.DateTo,
	) {
		// Create a new instance for each row
		currentReservation := &protos.ReservationResponse{
			Id:       reservation.Id,
			Accid:    reservation.Accid,
			Email:    reservation.Email,
			DateFrom: reservation.DateFrom,
			DateTo:   reservation.DateTo,
		}

		// Append the new instance to the slice
		reservationsSlice = append(reservationsSlice, currentReservation)
	}

	if err := iter.Close(); err != nil {
		rr.logger.Println(err)
		return nil, err
	}

	return reservationsSlice, nil
}

func (rr *ReservationRepo) DeleteByAccommodation(ctx context.Context, in *protos.DeleteRequestaa) (*protos.Emptyaa, error) {
	query := "DELETE FROM reservations WHERE accid = ?"
	if err := rr.session.Query(query, in.GetUid()).Exec(); err != nil {
		rr.logger.Println(err)
		return nil, errors.New("Couldn't delete")
	}

	return new(protos.Emptyaa), nil
}

func (rr *ReservationRepo) GetById(id int32) ([]*protos.ReservationResponse, error) {
	query := "SELECT id, accid, email, datefrom, dateto FROM reservations WHERE id = ? ALLOW FILTERING"
	iter := rr.session.Query(query, id).Iter()

	var reservationsSlice []*protos.ReservationResponse
	var reservation protos.ReservationResponse
	for iter.Scan(
		&reservation.Id,
		&reservation.Accid,
		&reservation.Email,
		&reservation.DateFrom,
		&reservation.DateTo,
	) {
		// Create a new instance for each row
		currentReservation := &protos.ReservationResponse{
			Id:       reservation.Id,
			Accid:    reservation.Accid,
			Email:    reservation.Email,
			DateFrom: reservation.DateFrom,
			DateTo:   reservation.DateTo,
		}

		// Append the new instance to the slice
		reservationsSlice = append(reservationsSlice, currentReservation)
	}

	if err := iter.Close(); err != nil {
		rr.logger.Println(err)
		return nil, err
	}

	return reservationsSlice, nil
}

func (rr *ReservationRepo) Create(profile *protos.ReservationResponse) error {
	query := "INSERT INTO reservations (id, accid, email, datefrom, dateto) VALUES (?, ?, ?, ?, ?) ALLOW FILTERING"

	err := rr.session.Query(query,
		profile.Id,
		profile.Accid,
		profile.Email,
		profile.DateFrom,
		profile.DateTo,
	).Exec()

	if err != nil {
		rr.logger.Println("Error inserting reservation record:", err)
		return err
	}

	return nil
}

func (rr *ReservationRepo) CheckIfAvailable(profile *protos.ReservationResponse) error {

	query := "SELECT id FROM reservations WHERE accid = ? AND ((datefrom <= ? AND dateto >= ?) OR (datefrom <= ? AND dateto >= ?)) ALLOW FILTERING"
	iter := rr.session.Query(query,
		profile.Accid,
		profile.DateFrom,
		profile.DateFrom,
		profile.DateTo,
		profile.DateTo,
	).Iter()

	if iter.NumRows() > 0 {
		rr.logger.Println("Reservation not available for the specified date range")
		return errors.New("Reservation not available for the specified date range")
	}

	return nil
}

func (rr *ReservationRepo) GetReservationsByEmail(ctx context.Context, in *protos.Emaill) (*protos.DummyLista, error) {
	query := "SELECT id, accid, email, datefrom, dateto FROM reservations WHERE email = ? ALLOW FILTERING"
	iter := rr.session.Query(query, in.GetEmail()).Iter()

	var lista protos.DummyLista
	var reservation protos.ReservationResponse
	for iter.Scan(
		&reservation.Id,
		&reservation.Accid,
		&reservation.Email,
		&reservation.DateFrom,
		&reservation.DateTo,
	) {
		// Create a new instance for each row
		currentReservation := &protos.ReservationResponse{
			Id:       reservation.Id,
			Accid:    reservation.Accid,
			Email:    reservation.Email,
			DateFrom: reservation.DateFrom,
			DateTo:   reservation.DateTo,
		}

		// Append the new instance to the slice
		lista.Dummy = append(lista.Dummy, currentReservation)
	}

	if err := iter.Close(); err != nil {
		rr.logger.Println(err)
		return nil, err
	}

	return &lista, nil
}

func (rr *ReservationRepo) DeleteReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	query := "DELETE FROM reservations WHERE email = ?"
	if err := rr.session.Query(query, in.GetEmail()).Exec(); err != nil {
		rr.logger.Println(err)
		return nil, errors.New("Couldn't delete")
	}

	return new(protos.Emptyaa), nil
}

func (rr *ReservationRepo) DeleteReservationById(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	query := "DELETE FROM reservations WHERE id = ?"
	if err := rr.session.Query(query, in.GetEmail()).Exec(); err != nil {
		rr.logger.Println(err)
		return nil, errors.New("Couldn't delete")
	}

	return new(protos.Emptyaa), nil
}

func (rr *ReservationRepo) CheckActiveReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	layout := "2006-01-02"
	query := "SELECT id FROM reservations WHERE email = ? AND datefrom <= ? AND dateto >= ? ALLOW FILTERING"
	iter := rr.session.Query(query,
		in.GetEmail(),
		time.Now().Format(layout),
		time.Now().Format(layout),
	).Iter()

	if iter.NumRows() > 0 {
		rr.logger.Println("Active reservation found for the specified date range")
		return new(protos.Emptyaa), nil
	}

	return nil, errors.New("No active reservation found")
}

func (rr *ReservationRepo) CheckActiveReservation(ctx context.Context, in *protos.DateFromDateTo) (*protos.Emptyaa, error) {
	layout := "2006-01-02"
	query := "SELECT id FROM reservations WHERE accid = ? AND datefrom <= ? AND dateto >= ? ALLOW FILTERING"
	iter := rr.session.Query(query,
		in.GetAccid(),
		time.Now().Format(layout),
		time.Now().Format(layout),
	).Iter()

	if iter.NumRows() > 0 {
		rr.logger.Println("Active reservation found for the specified date range")
		return new(protos.Emptyaa), nil
	}

	return nil, errors.New("No active reservation found")
}

func (rr *ReservationRepo) Update(reservation *protos.ReservationResponse) error {
	query := "UPDATE reservations SET datefrom = ?, dateto = ? WHERE id = ? ALLOW FILTERING"
	if err := rr.session.Query(query,
		reservation.DateFrom,
		reservation.DateTo,
		reservation.Id,
	).Exec(); err != nil {
		rr.logger.Println(err)
		return err
	}

	return nil
}
