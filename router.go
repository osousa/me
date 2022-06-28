package main

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	//"strconv"
)

type Router struct {
	name        string
	routes      []route
	middlewares Middlewares
}

type route struct {
	method      string
	regex       *regexp.Regexp
	handler     http.HandlerFunc
	middlewares []Middleware
	ctrl        Controller
}

type Extend interface{}

// Routes are set on Init . Each route added must have a method associated with
// it, as well as a View followed by a controller with the name of said View as
// parameter. An optional slice of type Middleware should be passed as the last
// parameter in order to manipulate the *Request accordingly.See middlewares.go
func (r *Router) SetRoutes() {
	r.AddRoute("GET", "/", home, NewController(Home{}, "home").(Home), nil)
	r.AddRoute("GET", "/static/(.+)", static, NewController(Home{}, "home").(Home), nil)
	r.AddRoute("GET", "/blog/([0-9]+)/([^/]+)", post, NewController(PostControl{}, "post").(PostControl), nil)
	r.AddRoute("GET", "/blog/([0-9]+)", blog, NewController(Blog{}, "blog").(Blog), nil)
	r.AddRoute("GET", "/blog", blog, NewController(Blog{}, "blog").(Blog), nil)
	r.AddRoute("GET", "/contact", contact, NewController(Contact{}, "contact").(Contact), nil)
	r.AddRoute("GET", "/about", about, NewController(About{}, "about").(About), nil)
	r.AddRoute("POST", "/admin", admin, NewController(Admin{}, "admin").(Admin), []Middleware{r.middlewares.auth})
}

// Instantiating Router named name with Middlewares slice.  These are geneneric
// Middlewares hereafter applied to every inbound or outbound  *Request that is
// made to the server and SetRoutes() is called in order to add routes to route
func NewRouter(name string, midwares Middlewares) Router {
	router := Router{name, nil, midwares}
	router.SetRoutes()
	return router
}

// Appends  new Route to Routes from Router. Pay attention to what patterns are
// added to avoid vulnerabilites.They should be fairly strict. Later adding reg
func (r *Router) AddRoute(method, pattern string, handler http.HandlerFunc, controller Controller, mware []Middleware) {
	r.routes = append(r.routes, route{method, regexp.MustCompile("^" + pattern + "$"), handler, mware, controller})
}

// The Router should implement Handler interface,thus the signature must be the
// same . The first  loop will select the corresponding route through the regex
// specified on creation, confirming if the method is allowed , for two methods
// to be allowed on the same route , just add another route and set a different
// Method (i.e.: "GET" , "UPDATE", etc) ; In order to be able to read variables
// extracted using the regex engine, we use context.WithValue method and add it
// to the  *Request.WithContext() . Check logic if middlewares are applied here
func (rt Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range rt.routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, struct {
				matches    []string
				controller Controller
			}{matches, route.ctrl})
			if route.middlewares != nil {
				temp_handler := http.HandlerFunc(route.handler)
				var res http.Handler
				for _, m := range route.middlewares {
					fmt.Println(m)
					res = m.UseMiddleware(temp_handler)
				}
				res.ServeHTTP(w, r.WithContext(ctx))
			} else {
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

// Extract URL parameter that matched against the regex and was passed through
// Context on http.Request,the index will dictate the position to be extracted
func getFields(r *http.Request, keys []string) (map[string]string, Controller) {
	fields := r.Context().Value(ctxKey{}).(struct {
		matches    []string
		controller Controller
	})
	fields_map := make(map[string]string)
	for i, value := range fields.matches[1:] {
		fields_map[keys[i]] = value
	}
	return fields_map, fields.controller
}

func static(w http.ResponseWriter, r *http.Request) {
	Warning.Println("reached static")
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static")))
	fs.ServeHTTP(w, r)
}

func admin(w http.ResponseWriter, r *http.Request) {

}

func home(w http.ResponseWriter, r *http.Request) {
	fields_map, controller := getFields(r, []string{"id"})
	controller.Execute(w, fields_map)
}

func blog(w http.ResponseWriter, r *http.Request) {
	fields_map, controler := getFields(r, []string{"page"})
	controler.Execute(w, fields_map)
}

func post(w http.ResponseWriter, r *http.Request) {
	fields_map, controler := getFields(r, []string{"id", "title"})
	controler.Execute(w, fields_map)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fields_map, controler := getFields(r, []string{})
	controler.Execute(w, fields_map)
}

func about(w http.ResponseWriter, r *http.Request) {
	fields_map, controler := getFields(r, []string{})
	controler.Execute(w, fields_map)
}

func apiCreateWidget(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "apiCreateWidget\n")
}
