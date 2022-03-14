package main

import (
    "net/http"
    "flag"
    "time"
    "log"
    "fmt"
    "db"
)


type timeHandler struct {
    format string
}

func version() string{
    version := "0.1.1";
    return version
}


func main() {
    versionPtr  := flag.Bool("version", false, "versioning")

    flag.Parse()

    if *versionPtr != false {
        fmt.Println(version())
        return
    }

    db, _     := database.ConnectDB("osousa")
    state     := db.GetConnState()
    router, _ := StartRouter("mainRouter")

    router.AddRoute("/",        "home",     "Welcome, this is home",            nil)
    router.AddRoute("/contacts","contacts", "You can contact me through here",  nil)

    fmt.Println(state)

    s := &http.Server{
        Addr:           ":8080",
        Handler:        router.GetMux(),
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
}

