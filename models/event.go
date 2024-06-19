package models

import (
	"time"

	"github.com/sambasivareddy-ch/rest_api_go/db"
)

// Base event struct
type Event struct {
	ID          int64
	EventName   string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}

/*
Saves the event in the Events tables.
*/
func (ev Event) Save() error {
	// Insert query to insert the event data into EVENTS table
	insertQuery := `INSERT INTO EVENTS(EVENTNAME, DESCRIPTION, LOCATION, DATETIME, USERID) VALUES(?, ?, ?, ?, ?)`

	insertPreparedStmt, err := db.AppDatabase.Prepare(insertQuery)
	if err != nil {
		return err
	}

	defer insertPreparedStmt.Close() // Close insert query on function exists

	// Execute the insert prepared query by passing user sent data, return err if any occurs
	_, err1 := insertPreparedStmt.Exec(
		ev.EventName,
		ev.Description,
		ev.Location,
		ev.DateTime,
		ev.UserID,
	)

	if err1 != nil {
		return err
	}

	return nil // returns null successfully inserted else err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"

	rows, err := db.AppDatabase.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Event{}

	for rows.Next() {
		event := Event{}

		if err := rows.Scan(&event.ID,
			&event.EventName,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(eventID int64) (Event, error) {
	query := `SELECT * FROM EVENTS WHERE ID = ?`
	targetEvent := Event{}

	// Execute the query by passing eventID as placeholder
	rows := db.AppDatabase.QueryRow(query, eventID)

	// Fetch the row data
	if err := rows.Scan(
		&targetEvent.ID,
		&targetEvent.EventName,
		&targetEvent.Description,
		&targetEvent.Location,
		&targetEvent.DateTime,
		&targetEvent.UserID,
	); err != nil {
		return Event{}, err
	}

	return targetEvent, nil
}

func (ev Event) UpdateEvent(eventId int64) error {
	updateQuery := `UPDATE EVENTS
		SET EventName = ?, Description = ?, Location = ?, DateTime = ?
		WHERE ID = ?
	`

	preparedStmt, err := db.AppDatabase.Prepare(updateQuery)
	if err != nil {
		return err
	}

	defer preparedStmt.Close()

	_, err = preparedStmt.Exec(
		ev.EventName,
		ev.Description,
		ev.Location,
		ev.DateTime,
		eventId,
	)
	if err != nil {
		return err
	}

	return nil //Successfully Updated
}

func DeleteEventByID(eventId int64) error {
	deleteQuery := "DELETE FROM EVENTS WHERE ID = ?"

	preparedStmt, err := db.AppDatabase.Prepare(deleteQuery)

	if err != nil {
		return err
	}

	defer preparedStmt.Close()

	_, err = preparedStmt.Exec(eventId)
	if err != nil {
		return err
	}

	return nil // Successfully Deleted
}
