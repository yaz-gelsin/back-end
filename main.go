package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yaz-gelsin/internal/handler"
	"github.com/yaz-gelsin/pkg"
)

func InitDB() (*sql.DB, error) {
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

	return conn, nil
}

func main() {
	// SQL bağlantısını oluşturun
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Router'ı başlatın
	router := handler.InitRouter(db)

	// Sunucuyu belirli bir portta dinlemeye başlayın
	log.Fatal(http.ListenAndServe(":8080", router))
}
