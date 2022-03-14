package main


import(
	"net/http"
	"log"
)
//var Routes interface{} = map[string]route{"hi":{path:""}}


type GenericHandler struct{
	content string
}


func (genhandlr GenericHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Content: " + genhandlr.content))
}



type routerModel interface {
	AddRoute(string, route)
	GetRoute(route) route
	PrintAllRoutes()
}

type route struct {
	genhandlr GenericHandler
	alias string
	path string
}

type router struct {
	routes map[string]route
	mux *http.ServeMux
	total int
	name string
}

func StartRouter(name string) (router, error){
	rtr := router{ routes:make(map[string]route) , name:"name", total:0, mux: http.NewServeMux()}
	return rtr, nil
}

func (rtr *router) GetMux() (*http.ServeMux){
	if rtr.mux == nil {
		log.Fatal("Router has no mux")
	}
	return rtr.mux
}

func (rtr *router) AddRoute(path, alias, content string, rt *route) {
	genhandlr := GenericHandler{content: content}
	if rt == nil {
		rtr.routes[path] = route{genhandlr:genhandlr, alias:alias, path:path}
	}else{
		rtr.routes[path] = *rt
	}
	if rtr.mux != nil {
		rtr.mux.Handle(path, genhandlr)
	} else {
		log.Fatal("Router has no mux")
	}
	return
}

