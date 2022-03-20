package main

import (
    "context"
    "fmt"
    "net/http"
    "regexp"
//    "strconv"
//    "log"
    "strings"
)


type Router struct{
    name string
    routes []route
    middlewares Middlewares

}

type route struct {
    method  string
    regex   *regexp.Regexp
    handler http.HandlerFunc
    middlewares []Middleware
    test Controller
}


func (r *Router) SetRoutes(){
    r.AddRoute("GET", "/",                      home,            NewController("home"),   nil)
    r.AddRoute("GET", "/api/widgets/([^/]+)",   apiUpdateWidget, NewController("about"),  nil)
    r.AddRoute("GET", "/contact",               contact,         NewController("home"),   nil)
    r.AddRoute("GET", "/api/widgets",           apiGetWidgets,   NewController("home"),   nil)
    r.AddRoute("GET", "/admin",                 apiGetWidgets,   NewController("home"),   []Middleware{r.middlewares.auth})
    /* 
    r.AddRoute("POST", "/api/widgets", apiCreateWidget, NewController("home"))
    r.AddRoute("POST", "/api/widgets/([^/]+)/parts", apiCreateWidgetPart, NewController("home"))
    r.AddRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/update", apiUpdateWidgetPart, NewController("home"))
    r.AddRoute("POST", "/api/widgets/([^/]+)/parts/([0-9]+)/delete", apiDeleteWidgetPart, NewController("home"))
    r.AddRoute("GET", "/([^/]+)", widget, NewController("home"))
    r.AddRoute("GET", "/([^/]+)/admin", widgetAdmin, NewController("home"))
    r.AddRoute("POST", "/([^/]+)/image", widgetImage, NewController("home")) */
}

func NewRouter(name string, midwares Middlewares)(Router){
    router := Router{name, nil, midwares}
    router.SetRoutes()
    return router
}

func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc, controller Controller, mware []Middleware) {
    r.routes = append(r.routes,route {method, regexp.MustCompile("^" + pattern + "$"), handler, mware, controller })
}

func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var allow []string
    for _, route := range rt.routes {
        matches := route.regex.FindStringSubmatch(r.URL.Path)
        if len(matches) > 0 {
            if r.Method != route.method {
                allow = append(allow, route.method)
                continue
            }
            ctx := context.WithValue(r.Context(), ctxKey{}, struct{matches []string; controller Controller}{matches, route.test})
            if(route.middlewares != nil){
                temp_handler := http.HandlerFunc(route.handler)
                var res http.Handler
                for _,m := range(route.middlewares){
                    fmt.Println(m)
                    res = m.UseMiddleware(temp_handler)
                }
                res.ServeHTTP(w, r.WithContext(ctx))
            }else{
                route.handler(w, r.WithContext(ctx))
            }
            return
        }
    }
    if len(allow) > 0 {
        w.Header().Set("Allow", strings.Join(allow, ", "))
        http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        return
    }
    http.NotFound(w, r)
}

type ctxKey struct{}

func getField(r *http.Request, index int) (string, Controller) {
    fields := r.Context().Value(ctxKey{}).(struct{matches []string; controller Controller})
    return fields.matches[index], fields.controller
}

func home(w http.ResponseWriter, r *http.Request) {
    _, controller := getField(r, 0)
    controller.Execute(w)
}

func contact(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "contact\n")
}

func apiGetWidgets(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "apiGetWidgets\n")
}

func apiCreateWidget(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "apiCreateWidget\n")
}

func apiUpdateWidget(w http.ResponseWriter, r *http.Request) {
    _, controller := getField(r, 0)
    controller.Execute(w)
}

/*
func apiCreateWidgetPart(w http.ResponseWriter, r *http.Request) {
    slug := getField(r, 0)
    fmt.Fprintf(w, "apiCreateWidgetPart %s\n", slug)
}

func apiUpdateWidgetPart(w http.ResponseWriter, r *http.Request) {
    slug := getField(r, 0)
    id, _ := strconv.Atoi(getField(r, 1))
    fmt.Fprintf(w, "apiUpdateWidgetPart %s %d\n", slug, id)
}

func apiDeleteWidgetPart(w http.ResponseWriter, r *http.Request) {
    slug := getField(r, 0)
    id, _ := strconv.Atoi(getField(r, 1))
    fmt.Fprintf(w, "apiDeleteWidgetPart %s %d\n", slug, id)
}

func widget(w http.ResponseWriter, r *http.Request) {
    slug := getField(r, 0)
    fmt.Fprintf(w, "widget %s\n", slug)
}

func widgetAdmin(w http.ResponseWriter, r *http.Request) {
    slug := getField(r, 0)
    fmt.Fprintf(w, "widgetAdmin %s\n", slug)
}

func widgetImage(w http.ResponseWriter, r *http.Request) {
    slug := getField(r, 0)
    fmt.Fprintf(w, "widgetImage %s\n", slug)
}

*/