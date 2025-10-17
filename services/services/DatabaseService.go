package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func ConnectToDb(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var db *sql.DB
	var err error

	for i := 1; i <= 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Attempt %d: failed to open connection: %v", i, err)
		} else {
			err = db.Ping()
			if err == nil {
				log.Printf("Connected to DB on attempt %d", i)
			}
			log.Printf("Attempt %d: waiting for DB to be ready: %v", i, err)
		}

		time.Sleep(2 * time.Second)
	}

	MigrateTable(db)
	return db, nil
}

func MigrateTable(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL,
		created_at TIMESTAMPTZ NOT NULL DEFAULT now()
	);
	`
	if _, err := db.Exec(schema); err != nil {
		return err
	}

	return nil
}

func AddUsersToTable(db *sql.DB) {
	if db == nil {
		fmt.Println("AddUsersToTable: DB is nil (not initialized)")
	}

	AddUsers(db)
}

func FindUserById(db *sql.DB, id int) (*User, error) {

	if db == nil {
		fmt.Println("FindUserById: DB is nil (not initialized)")
		return nil, errors.New("Failed find user")
	}

	user, err := GetUserById(db, id)

	if err != nil {
		return nil, err
	}

	return user, err
}
