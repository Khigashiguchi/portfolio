package app

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/handlers"
    "github.com/Khigashiguchi/portfolio/app/infrastructure"
)

// confDir is configuration file directory.
const confDir = "./config"

// conf is global configuration.
var conf infrastructure.Config

func init() {
    var err error
    conf, err = infrastructure.NewConfig(os.Getenv("APP_ENV"), confDir)
    if err != nil {
        msg := fmt.Sprintf("failed to get configuration: %s", err)
        fmt.Fprintf(os.Stderr, msg)
        panic(msg)
    }
}

func main() {
    // Create global infrastructure handler.
    db, err := infrastructure.NewDB(conf)
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to create database global connection: %s", err)
    }
	router := infrastructure.Router(db)

	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, router))

	// Start server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
