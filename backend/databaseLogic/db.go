package databaseLogic

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DbConnection() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error with the .env file")
	}
	host := os.Getenv("HOST")
	portStr := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		fmt.Print("Error with opening the db connection")
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Established a connection!")
	return db
}
