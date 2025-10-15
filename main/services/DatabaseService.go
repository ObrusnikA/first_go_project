package services

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectToDb(host string, port int, user string, password string, dbname string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("ping error: %w", err)
	}

	fmt.Println("Connected to database successfully!")

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

func FindUserById(db *sql.DB, id int) User {

	if db == nil {
		fmt.Println("FindUserById: DB is nil (not initialized)")
	}

	user, err := GetUserById(db, id)

	if err != nil {
		return User{}
	}

	return *user
}
