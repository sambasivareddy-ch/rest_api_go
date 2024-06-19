package models

import (
	"github.com/sambasivareddy-ch/rest_api_go/db"
	"github.com/sambasivareddy-ch/rest_api_go/utils"
)

type User struct {
	ID       int64
	EMAIL    string `binding:"required"`
	PASSWORD string `binding:"required"`
}

func (usr User) Save() error {
	insertQuery := `INSERT INTO USERS (EMAIL, PASSWORD) VALUES (?, ?)`

	stmt, err := db.AppDatabase.Prepare(insertQuery)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err1 := utils.GenerateHashPassword(usr.PASSWORD)

	if err1 != nil {
		return err1
	}

	_, err = stmt.Exec(usr.EMAIL, hashedPassword)

	if err != nil {
		return err
	}

	return nil
}

// Returns user's password based on email provided else ""
func GetUserIDAndPasswordByEmail(email string) (int64, string, error) {
	queryStmt := "SELECT id, password FROM users WHERE EMAIL = ?"

	preparedStmt, err := db.AppDatabase.Prepare(queryStmt)

	if err != nil {
		return 0, "", err
	}

	var userPassword string
	var userId int64
	row := preparedStmt.QueryRow(email)

	err = row.Scan(&userId, &userPassword) // Scan password from resultant row

	if err != nil {
		return 0, "", err
	}

	return userId, userPassword, nil
}
