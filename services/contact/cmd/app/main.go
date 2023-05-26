package main

import (
	"aida.net/pkg/store/postgres"
	"fmt"
)

func main() {
	// params to connect db
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "29012004910_Rich"
	dbname := "example"

	// open connection
	db, err := postgres.ConnectDB(host, port, user, password, dbname)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	} else {
		fmt.Println("ok")
	}

	defer db.Close()
}
