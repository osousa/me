package main

import (
    "net/http"
    "testing"
    //"net/http/httptest"
)


func Test_NewRouter(t *testing.T) {
    router := NewRouter("Test", Middlewares{})
    if(router.name!="Test"){
        t.Fatalf(`Router name not attributed`)
    }
}


func Test_AddRoute(t *testing.T){
    router := NewRouter("Test", Middlewares{})
    router.AddRoute("GET", "/",                      test_home,     NewController("home"),   nil)
    router.AddRoute("GET", "/api/widgets/([^/]+)",   test_home,     NewController("about"),  nil)
}

func test_home(w http.ResponseWriter, r *http.Request){
    //assert.Equal(t, b, code >= 300)
}