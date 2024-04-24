package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

type User struct {
	email string
}

var db *sql.DB

func connectToDatabase() {
	cfg := *mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.DBName = "my_website"

    var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")
}

func userByEmail(email string) (User, error) {
	var user User
    var throwaway_value string

	row := db.QueryRow("SELECT * FROM users WHERE unique_email = ?;", email)
    if err := row.Scan(&user.email, &throwaway_value); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("userByEmail %v: no such user", email)
		}
        return user, fmt.Errorf("userByEmail %v: %v", email, err)
	}
	return user, nil
}
