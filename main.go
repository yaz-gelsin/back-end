package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/yaz-gelsin/internal/adapter/mysqlrepo/product"
	"github.com/yaz-gelsin/internal/handler/api"
	"github.com/yaz-gelsin/internal/usecase"
	"github.com/yaz-gelsin/pkg"
)

func InitDB() (*sqlx.DB, error) {
	// Load the configuration
	config, err := pkg.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open a database connection
	conn, err := sqlx.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Check the connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping db:", err)
	}

	fmt.Println("Connected to the database!")

	return conn, nil
}

func main() {
	// SQL connection initialization
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router
	router := mux.NewRouter()

	repo := product.NewProductRepo(db)

	// Create an instance of your use case implementation
	useCase := usecase.NewUseCase(repo)

	// Create an instance of your handler and pass the use case implementation and DB instance
	handler := api.NewHandler(useCase, db)

	// Initialize the router using the InitRouter function
	routerHandler := handler.InitRouter(router, db)

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", routerHandler))
}
