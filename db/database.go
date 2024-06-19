package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Global Instance to access the Database created
var AppDatabase *sql.DB

/*
Initialize the Database on App start/restart and creates the required tables
if those tables doesn't exists in the database.
*/
func InitDB() error {
	var err error
	AppDatabase, err = sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal("failed to connect to Database")
	}

	createTables()

	return AppDatabase.Ping()
}

/*
Create the tables required like Events & Users
*/
func createTables() {
	createUsersTableCommand := `CREATE TABLE IF NOT EXISTS USERS (
		id 			integer primary key autoincrement,
		email 		text unique not null,
		password 	text not null
	)`

	createEventsTableCommand := `CREATE TABLE IF NOT EXISTS EVENTS (
		id          integer primary key autoincrement,
		eventname   varchar(30) not null,    
		description text not null,
		location    varchar(30) not null,
		datetime    timestamp not null,
		userid      integer,
		FOREIGN KEY (userid) REFERENCES USERS(id)
	)`

	createRegistrationTableCommand := `CREATE TABLE IF NOT EXISTS REGISTRATION (
		userId	integer not null,
		eventId integer not null,
		FOREIGN KEY (userId) REFERENCES USERS(id)
		FOREIGN KEY (eventId) REFERENCES EVENTS(id)
	)`

	stmt, err := AppDatabase.Prepare(createUsersTableCommand)
	if err != nil {
		panic("unable to prepare create users table command")
	}

	if _, err = stmt.Exec(); err != nil {
		panic("unable to create users table")
	}

	stmt, err = AppDatabase.Prepare(createEventsTableCommand)
	if err != nil {
		panic("unable to prepare create events table command")
	}

	if _, err = stmt.Exec(); err != nil {
		panic("unable to create events table")
	}

	stmt, err = AppDatabase.Prepare(createRegistrationTableCommand)
	if err != nil {
		panic("unable to prepare create registration table command")
	}

	if _, err = stmt.Exec(); err != nil {
		panic("unable to create registration table")
	}
}
