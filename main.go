package main

import (
	"database/sql"
	"first_pet/main/services"
	"fmt"
)

func main() {
	db, err := services.ConnectToDb("localhost", 5432, "postgres", "admin", "postgres")
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
