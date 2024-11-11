package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stonoy/get_social/internal"
)

type apiConfig struct {
	db *internal.Queries
}

func main() {
	// load env files
	err := godotenv.Load()
	if err != nil {
		log.Panicf("error inloading env variable -> %v", err)
		return
	}

	apiCfg := apiConfig{}

	port := os.Getenv("PORT")
	if port == "" {
		log.Panicln("no port in env file")
		return
	}

	conn_url := os.Getenv("CONN_URL")
	if conn_url != "" {
		// open the driver
		db, err := sql.Open("postgres", conn_url)
		if err != nil {
			log.Panicf("can not connect to the server via go -> %v", err)
		}

		db_quries := internal.New(db)

		apiCfg.db = db_quries
	} else {
		log.Println("server started without database connection")
	}

	// create a main router
	mainRouter := chi.NewRouter()

	// create a new server from http.Server type
	the_server := &http.Server{
		Addr:    ":" + port,
		Handler: mainRouter,
	}

	log.Printf("Server is listenning on port %v", port)

	// starts the server
	log.Panic(the_server.ListenAndServe())
}
