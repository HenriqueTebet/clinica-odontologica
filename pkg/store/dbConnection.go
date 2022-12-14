package store

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

var dbConnection = ""

func LoadEnvVariables() {
	var err error
	if err = godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	dbConnection = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
	os.Getenv("USERNAME"), os.Getenv("PASS"), os.Getenv("DB_URL"), os.Getenv("PORT"), os.Getenv("DB_NAME"))
}

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConnection)
	
	if(err != nil) {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}