package main

import (
    "net/http"
    "flag"
    "time"
    "log"
    "fmt"
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

    // Fire up the database, no need to disconnect.
    // Just make sure all connections are deferred/closed.
    //
    db, _     := ConnectDB("osousa")
    _          = db.GetConnState()


    // Start a router and activate preconfigured routes.
    // Middleware association should probably be done here.
    //
    middlewares  := NewMiddlewares("default")
    router       := NewRouter("default", middlewares)
    router.InitRoutes()




    // Create a Config struct for server later on
    // Do not use middlewares here.
    //
    s := &http.Server{
        Addr:           ":8080",
        Handler:        middlewares.UseDefaultMiddlewares(router),
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    log.Fatal(s.ListenAndServe())
}
