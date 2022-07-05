package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var DB Database

var Info = log.New(os.Stdout, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Warning = log.New(os.Stdout, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
var Error = log.New(os.Stdout, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
var Debug = log.New(os.Stdout, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Does it exist?")
	}

	db_pass := os.Getenv("DATABASE_PASSWORD")
	db_name := os.Getenv("DATABASE_NAME")
	db_user := os.Getenv("DATABASE_USER")

	flag.Parse()
	versionPtr := flag.Bool("version", false, "versioning")
	if *versionPtr != false {
		fmt.Println(os.Getenv("VERSION"))
		return
	}

	// Fire up the database, no need to disconnect.
	// Just make sure all connections are deferred/closed.
	DB, _ = ConnectSQL(db_user, db_pass, db_name)
	_ = DB.GetState()

	// Start a router and activate preconfigured routes.
	// Middleware association should probably be done here.
	middlewares := NewMiddlewares("default")
	router := NewRouter("default", middlewares)

	// Create a Config struct for server later on
	// Do not use middlewares here.
	s := &http.Server{
		Addr:           ":8080",
		Handler:        middlewares.UseCommonMiddlewares(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
