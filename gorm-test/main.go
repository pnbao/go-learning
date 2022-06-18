package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Log struct {
	ID               string
	URL              string
	FUNC             string
	REQUEST_ID       string
	MSG              string
	DATA_REQ         string
	DATA_RES         string
	TID              string
	STATUS           string
	COUNT_TIME       int
	TRANSACTION_TYPE string
	STEP             string
	STATE            string
	DATEM            string
	HOURM            string
	ERROR_STACK      string
	ACCOUNT_NUMBER   string
	CREATED_AT       string
}

func main() {
	db, err := gorm.Open(sqlite.Open("stm.sqlite3"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var log Log
	db.First(&log, "id = ?", "2113")

	fmt.Println("First log", log)
}
