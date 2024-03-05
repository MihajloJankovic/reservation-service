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

	// Return repository with logger and DB session
	return &ReservationRepo{
		session: session,
		logger:  logger,
	}, nil
}

func (rr *ReservationRepo) CreateTables(ctx context.Context) {
	err := rr.session.Query(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s
			(id int,
			accid text,
			email text,
			datefrom DATE,
			dateto DATE , PRIMARY KEY(id,accid))`, "reservations_by_id_and_accid")).Exec()
	if err != nil {
		rr.logger.Println(err)
	}
	err = rr.session.Query(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s
			(id int,
			accid text,
			email text,
			datefrom DATE,
			dateto DATE , PRIMARY KEY(id,email))`, "reservations_by_id_and_email")).Exec()
	if err != nil {
		rr.logger.Println(err)
	}
	err = rr.session.Query(fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s
			(id int ,
			accid text,
			email text,
			datefrom DATE,
			dateto DATE , PRIMARY KEY(accid,datefrom,dateto))`, "reservations_by_accid_and_dates")).Exec()
	if err != nil {
		rr.logger.Println(err)
	}
}

// Disconnect from database
func (rr *ReservationRepo) Disconnect(ctx context.Context) {
	rr.session.Close()
}

func (rr *ReservationRepo) DeleteReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	temp, _ := rr.GetReservationsByEmail(ctx, in)
	for _, element := range temp.Dummy {
		query := "DELETE FROM reservations_by_id_and_accid WHERE id = ? AND accid = ?"
		if err := rr.session.Query(query, element.GetId(), element.GetAccid()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
		query1 := "DELETE FROM reservations_by_id_and_email WHERE email = ? AND id = ?"
		if err := rr.session.Query(query1, in.GetEmail(), element.GetId()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
		query2 := "DELETE FROM reservations_by_accid_and_dates WHERE accid = ? AND datefrom = ? AND dateto = ?"
		if err := rr.session.Query(query2, element.GetAccid(), element.GetDateFrom(), element.GetDateTo()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
	}

	return new(protos.Emptyaa), nil
}
func (rr *ReservationRepo) DeleteByAccommodation(ctx context.Context, in *protos.DeleteRequestaa) (*protos.Emptyaa, error) {
	temp, _ := rr.GetReservationsByAcc(ctx, in)
	for _, element := range temp.Dummy {
		query := "DELETE FROM reservations_by_id_and_accid WHERE id = ? AND accid = ?"
		if err := rr.session.Query(query, element.GetId(), element.GetAccid()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
		query1 := "DELETE FROM reservations_by_id_and_email WHERE  email = ? AND  id = ?"
		if err := rr.session.Query(query1, element.GetEmail(), element.GetId()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
		query2 := "DELETE FROM reservations_by_accid_and_dates WHERE accid = ? AND datefrom = ? AND dateto = ?"
		if err := rr.session.Query(query2, element.GetAccid(), element.GetDateFrom(), element.GetDateTo()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
	}

	return new(protos.Emptyaa), nil
}
func (rr *ReservationRepo) DeleteReservationById(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {

	temp, _ := rr.GetByIdIntern(in.GetEmail())
	for _, element := range temp {
		query := "DELETE FROM reservations_by_id_and_accid WHERE id = ? AND accid = ?"
		if err := rr.session.Query(query, element.GetId(), element.GetAccid()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
		query1 := "DELETE FROM reservations_by_id_and_email WHERE  email = ? AND id = ?"
		if err := rr.session.Query(query1, element.GetEmail(), element.GetId()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
		query2 := "DELETE FROM reservations_by_accid_and_dates WHERE accid = ? AND datefrom = ? AND dateto = ?"
		if err := rr.session.Query(query2, element.GetAccid(), element.GetDateFrom(), element.GetDateTo()).Exec(); err != nil {
			rr.logger.Println(err)
			return nil, errors.New("Couldn't delete")
		}
	}

	return new(protos.Emptyaa), nil
}
func (rr *ReservationRepo) CheckActiveReservation(ctx context.Context, in *protos.DateFromDateTo) (*protos.Emptyaa, error) {
	layout := "2006-01-02"
	query := "SELECT id FROM reservations_by_accid_and_dates WHERE accid = ? AND datefrom <= ? AND dateto >= ? "
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

func (rr *ReservationRepo) CheckIfAvailable(profile *protos.ReservationResponse) error {
	//ssss
	query := "SELECT id, accid, email, datefrom, dateto FROM reservations_by_accid_and_dates WHERE accid = ? AND datefrom <= ? AND dateto >= ? "
	iter := rr.session.Query(query,
		profile.Accid,
		profile.DateFrom,
		profile.DateFrom,
	).Iter()

	if iter.NumRows() > 0 {
		rr.logger.Println("Reservation not available for the specified date range")
		return errors.New("Reservation not available for the specified date range")
	}
	query1 := "SELECT id, accid, email, datefrom, dateto FROM reservations WHERE accid = ? AND datefrom <= ? AND dateto >= ?"

	iter1 := rr.session.Query(query1,
		profile.Accid,
		profile.DateTo,
		profile.DateTo,
	).Iter()

	if iter1.NumRows() > 0 {
		rr.logger.Println("Reservation not available for the specified date range")
		return errors.New("Reservation not available for the specified date range")
	}
	return nil
}
func (rr *ReservationRepo) GetReservationsByAcc(ctx context.Context, in *protos.DeleteRequestaa) (*protos.DummyLista, error) {
	query := "SELECT id, accid, email, datefrom, dateto FROM reservations_by_accid_and_dates WHERE accid = ? "
	iter := rr.session.Query(query, in.GetUid()).Iter()

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
func (rr *ReservationRepo) GetReservationsByEmail(ctx context.Context, in *protos.Emaill) (*protos.DummyLista, error) {
	query := "SELECT id, accid, email, datefrom, dateto FROM reservations_by_id_and_email WHERE email = ? "
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
func (rr *ReservationRepo) GetById(id int32) ([]*protos.ReservationResponse, error) {
	query := "SELECT id, accid, email, datefrom, dateto FROM reservations_by_id_and_accid WHERE id = ?"
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
func (rr *ReservationRepo) GetByIdIntern(id string) ([]*protos.ReservationResponse, error) {
	query := "SELECT id, accid, email, datefrom, dateto FROM reservations_by_id_and_accid WHERE id = ?"
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
	query := "INSERT INTO reservations_by_id_and_accid (id, accid, email, datefrom, dateto) VALUES (?, ?, ?, ?, ?)"

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
	query2 := "INSERT INTO reservations_by_id_and_email (id, accid, email, datefrom, dateto) VALUES (?, ?, ?, ?, ?)"

	err = rr.session.Query(query2,
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
	query1 := "INSERT INTO reservations_by_accid_and_dates (id, accid, email, datefrom, dateto) VALUES (?, ?, ?, ?, ?)"

	err = rr.session.Query(query1,
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

func (rr *ReservationRepo) Update(reservation *protos.ReservationResponse) error {
	return nil
}

func (rr *ReservationRepo) CheckActiveReservationByEmail(ctx context.Context, in *protos.Emaill) (*protos.Emptyaa, error) {
	return nil, nil
}
func (rr *ReservationRepo) GetAll() ([]*protos.ReservationResponse, error) {
	return nil, nil
}
