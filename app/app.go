package app

import (
	"banking/domain"
	"banking/service"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
)

// SERVER_ADDRESS=localhost SERVER_PORT=8000 DB_USER=root DB_PASSWD=I240959ko DB_ADDR=localhost DB_PORT=3306 DB_NAME=banking go run main.go

// функция, позволяющая проверить наличие всех переменных окружения
// можно экспортировать переменные, чтобы их не вводить каждый раз export SERVER_ADDRESS=localhost
func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not defined")
	}
}

func Start() {

	sanityCheck()
	router := mux.NewRouter()

	dbClient := getDbClient()

	// связать все воедино // wiring
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	ch := CustomerHandlers{service: service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	th := TransactionHandler{service: service.NewTransactionService(transactionRepositoryDb)}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet).Name("GetAllCustomers")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet).Name("GetCustomer")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost).Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost).Name("NewTransaction")

	// adding middleware to the routes
	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	//внеднение переменных окружения для изменения адреса и порта из командной строки SERVER_ADDRESS=localhost SERVER_PORT=8262 go run main.go
	//injecting environment variables
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	// start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
