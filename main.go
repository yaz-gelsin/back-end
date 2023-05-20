package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yaz-gelsin/pkg"
)

func main() {
	// Load the configuration
	config, err := pkg.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open a database connection
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	// Check the connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping db:", err)
	}

	fmt.Println("Connected to the database!")

}
