package user

import (
	"database/sql"
	"fmt"
)

func RegisterUser(db *sql.DB, username, email, password string) error {
	query := `INSERT INTO users(username, email, password) VALUES ($1,$2,$3)`

	_, err := db.Exec(query, username, email, password)
	if err != nil {
		return fmt.Errorf("Error inserting user into database: %s", err)
	}
	fmt.Println("User registered")
	return nil

}

func LoginUser(db *sql.DB, username, password string) error {
	query := `SELECT id FROM users WHERE EXISTS (SELECT 1 FROM users WHERE username = $1 AND password = $2)`

	var exist bool
	err := db.QueryRow(query, username, password).Scan(&exist)

	if err != nil {
		fmt.Println("User does not found")
		return err
	}

	fmt.Printf("Welcome back %s!", username)
	return nil
}
