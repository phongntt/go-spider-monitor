package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	//connStr := "host=192.168.72.122 port=5444 user=recruit1 password=recruit dbname=helijobs sslmode=disable"
	connStr := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Args[1], os.Args[2], os.Args[3], os.Args[4], os.Args[5])

	println(connStr)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print("1- WARNING| Cannot connect to DB.")
		os.Exit(1)
	}

	err = db.Ping()
	if err != nil {
		fmt.Print("1- WARNING| Cannot connect to DB.")
		os.Exit(1)
	}

	rows, err := db.Query("SELECT 1 FROM DUAL")
	if err != nil {
		fmt.Print("1- WARNING| Cannot SELECT on DB.")
		os.Exit(1)
	}
	defer rows.Close()

	fmt.Print("0- OK| Success running select on DB.")
}
