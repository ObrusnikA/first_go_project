package main

import (
	"database/sql"
	"first_pet/services/services"
	"fmt"
	"os"
	"strconv"
)

func main() {
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid port: %v", portStr))
	}

	fmt.Printf("Connecting to DB at %s:%d\n", host, port)

	db, err := services.ConnectToDb(host, port, user, password, dbname)
	if err != nil {
		fmt.Println("Failed", err)
		return
	}

	fmt.Println("Database is ready to use!")
	GetUser(db, 1)

	defer db.Close()
}

func GetUser(db *sql.DB, id int) {

	if db == nil {
		fmt.Println("GetUser: DB is nil (not initialized)")
	}

	services.AddUsersToTable(db)

	user, err := services.FindUserById(db, 1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.Username)
}
