package dbconnector

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type dbConnectionInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLmode  string `json:"sslmode"`
}

// NewDBConnection parses a JSON file located at auctionsjson/database.json and uses
// the information contained in the file to create a database connection
func NewDBConnection() *sql.DB {
	// Open the JSON file that should contain the information to the database
	file, err := os.Open("../auctionjson/database.json")
	if err != nil {
		fmt.Println("Opening JSON file in NewDBConnection failed")
		log.Fatal(err)
	}

	// Read the file into a reader
	reader, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Reading the JSON file failed in NewDBConnection")
		log.Fatal(err)
	}
	dbInfo := dbConnectionInfo{}
	err = json.Unmarshal(reader, &dbInfo)
	if err != nil {
		fmt.Println("Parsing values from JSON file into dbInfo failed in NewDBConnection")
		log.Fatal(err)
	}
	psqlConnInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.Host, dbInfo.Port, dbInfo.User, dbInfo.Password, dbInfo.DBName)
	db, err := sql.Open("postgres", psqlConnInfo)
	if err != nil {
		fmt.Println("sql.Open failed to create a connection in NewDBConnection")
		log.Fatal(err)
	}
	return db
}
