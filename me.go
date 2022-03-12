package main

import (
    "net/http"
    "flag"
    "fmt"
)

func version() string{
    version := "0.1";
    return version
}


//1--little comment
func HelloWorld(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Hello World")
}

func main() {
    versionPtr := flag.Bool("version", false, "versioning")
    flag.Parse()

    if *versionPtr != false {
        fmt.Println(version())
        return
    }

    http.HandleFunc("/", HelloWorld)
    http.ListenAndServe(":3000", nil)
}
