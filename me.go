package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var DB = &db{}

func version() string {
	version := "0.1.1"
	return version
}

func main() {
	versionPtr := flag.Bool("version", false, "versioning")

	flag.Parse()

	if *versionPtr != false {
		fmt.Println(version())
		return
	}

	// Fire up the database, no need to disconnect.
	// Just make sure all connections are deferred/closed.
	//
	DB, _ = ConnectDB("osousa")
	_ = DB.GetConnState()

	// Start a router and activate preconfigured routes.
	// Middleware association should probably be done here.
	//
	middlewares := NewMiddlewares("default")
	router := NewRouter("default", middlewares)

	user := new(User)
	err_user := GetById(user, 9)
	//user := new(User)
	//err_user := user.GetById(9)
	fmt.Println(err_user)
	fmt.Println("user: ", user)
	fmt.Println("user.Pass: ", *user.Pass)
	fmt.Println("user.Pass: ", user.Pass)
	user.Save()

	exp := new(Experience)
	err_exp := GetById(exp, 9)
	//exp := new(Experience)
	//err_exp := exp.GetById(9)
	fmt.Println(err_exp)
	fmt.Println("Experience: ", exp)
	fmt.Println("Experience.Company: ", *exp.Company)
	fmt.Println("Experience.Company: ", exp.Company)

	// Create a Config struct for server later on
	// Do not use middlewares here.
	//
	s := &http.Server{
		Addr:           ":8080",
		Handler:        middlewares.UseCommonMiddlewares(router),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
