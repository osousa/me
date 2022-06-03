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

var DB = &db{}

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

	Info.Println(db_name)
	Info.Println(db_pass)

	flag.Parse()

	versionPtr := flag.Bool("version", false, "versioning")
	if *versionPtr != false {
		fmt.Println(os.Getenv("VERSION"))
		return
	}

	// Fire up the database, no need to disconnect.
	// Just make sure all connections are deferred/closed.
	//
	DB, _ = ConnectDB(db_user, db_pass, db_name)
	_ = DB.GetConnState()

	// Start a router and activate preconfigured routes.
	// Middleware association should probably be done here.
	//
	middlewares := NewMiddlewares("default")
	router := NewRouter("default", middlewares)

	//user := new(User)
	//err_user := GetById(user, 9)
	////user := new(User)
	////err_user := user.GetById(9)
	//fmt.Println(err_user)
	//fmt.Println("user: ", user)
	//fmt.Println("user.Pass: ", *user.Pass)
	//new_pass := "do_tests_cmon_Man"
	//user.Pass = &new_pass
	//err_usr := user.Save()
	//if err_usr != nil {
	//	Warning.Println(err_usr)
	//}

	//post := new(Post)
	//err_post := GetById(post, 9)
	////post := new(User)
	////err_post := user.GetById(9)
	//if post == nil {
	//	Warning.Println("post is nil")
	//}
	//fmt.Println("error Post : ", err_post)
	//fmt.Println("post: ", post)
	//new_title := "do_tests_cmon_Man"
	//post.Title = &new_title
	//err_pst := post.Save()
	//if err_pst != nil {
	//	Warning.Println(err_pst)
	//}

	//fmt.Println(err_pst)
	//fmt.Println("Post: ", post)
	//fmt.Println("Post Title : ", *post.Title)
	//fmt.Println("Post body : ", *post.Body)
	//fmt.Println("Post date: ", *post.Date)

	//exp := new(Experience)
	//err_exp := GetById(exp, 9)
	//exp := new(Experience)
	//err_exp := exp.GetById(9)
	//fmt.Println(err_exp)
	//fmt.Println("Experience: ", exp)
	//fmt.Println("Experience Company: ", *exp.Company)
	//fmt.Println("Experience.Position: ", *exp.Position)
	//exp2 := &Experience{10, new(string), new(string), new(string)}
	//*exp2.Description = "good"
	//*exp2.Position = "nice"
	//*exp2.Company = "sveet"
	//exp2.Id = 11
	//DB.InsertRow(exp2)
	lpost := new(Post)
	err_lst, lst := GetList(lpost, 0)

	if err_lst != nil {
		Error.Println(err_lst)
	}

	for _, val := range lst {
		Info.Println(val)
	}
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
