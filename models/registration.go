package models

import "github.com/sambasivareddy-ch/rest_api_go/db"

type Register struct {
	ID      int64
	EventId int64 `binding:"required"`
}

func (r Register) Save() error {
	saveCommand := `INSERT INTO REGISTRATION VALUES (?, ?)`

	stmt, err := db.AppDatabase.Prepare(saveCommand)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(r.ID, r.EventId)
	if err != nil {
		return err
	}

	// Successfully Registered
	return nil
}

func UnregisterToEvent(userId int64, eventId int64) error {
	deleteCommand := `DELETE FROM REGISTRATION WHERE userId = ? AND eventId = ?`

	stmt, err := db.AppDatabase.Prepare(deleteCommand)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userId, eventId)
	if err != nil {
		return err
	}

	return nil
}
