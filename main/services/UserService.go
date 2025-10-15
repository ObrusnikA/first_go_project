package services

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

var ctx context.Context

type User struct {
	Id        int
	Username  string
	Email     string
	CreatedAt time.Time
}

func AddUsers(db *sql.DB) {

	if db == nil {
		fmt.Println("AddUsers: DB is nil (not initialized)")
	}

	var lastId int

	qResult := db.QueryRowContext(ctx, `select COALESCE(MAX(id), 0) from users`).Scan(&lastId)
	if qResult != nil {
		fmt.Printf(qResult.Error())
	}

	const q = `INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id, created_at`

	insResult := db.QueryRowContext(ctx, q, fmt.Sprintf("User_%d", lastId+1), fmt.Sprintf("Email@%d", lastId+1))
	if insResult != nil {
		fmt.Println(insResult.Err())
	}
}

func GetUserById(db *sql.DB, id int) (*User, error) {

	if db == nil {
		fmt.Println("GetUserById: DB is nil (not initialized)")
	}

	var u User

	qResult := db.QueryRowContext(ctx, `select id, username, email, created_at from users where id = @id`, id).Scan(&u.Id, &u.Username, &u.Email, &u.CreatedAt)
	if qResult != nil {
		return nil, qResult
	}

	return &u, qResult
}
