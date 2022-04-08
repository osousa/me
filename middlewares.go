package main

import (
	"log"
	"net/http"
)

type Middleware interface {
	UseMiddleware(http.Handler) http.Handler
}

type Middlewares struct {
	name   string
	common []Middleware
	auth   Middleware
}

// Each struct will represent a type of middleware . We can create many of each
// On this app we'll have only one Middleware_Auth  Middleware instantiated for
// example. If one needs more middlewares, start by creating a struct type that
// implements methods from the Middleware interface above
type Middleware_SecureHeaders struct {
	name   string
	header map[string]string
}

type Middleware_Auth struct {
	name string
}

func InitMiddlewares(name string) Middlewares {
	return Middlewares{name: name, common: []Middleware{}}
}

func NewMiddlewares(name string) Middlewares {
	m := InitMiddlewares(name)
	m.SetCommonMiddlewares()
	m.SetAuthMiddleware()
	return m
}

func NewAuth(name string) Middleware_Auth {
	return Middleware_Auth{name}
}

func NewSecureHeaders(name string) Middleware_SecureHeaders {
	return Middleware_SecureHeaders{name, nil}
}

func (m *Middlewares) SetAuthMiddleware() {
	auth := NewAuth("auth")
	m.auth = auth
}

//
func (m *Middlewares) SetCommonMiddlewares() {
	headers := NewSecureHeaders("default")
	headers.SetSecureHeaders()
	m.common = append(m.common, headers)
}

// Add at will security headers, these are added at each outgoing http.Response
// and should provide enhanced security to common flaws, like basic XSS and not
// allowing CORS violations, for example
func (s *Middleware_SecureHeaders) SetSecureHeaders() {
	s.header = map[string]string{
		"Content-Security-Policy":      "default-src 'self' osousa.me; connect-src 'self';",
		"Content-Type":                 "text/html;",
		"X-XSS-Protection":             "1; mode=block",
		"Cross-Origin-Resource-Policy": "same-site",
		"X-Content-Type-Options":       "nosniff",
		"X-Frame-Options":              "deny"}
}

// Each middleware implement Middleware interface methods, and UseMiddleware is
// used to provide changes to requests and responses alike.The following method
// is called upon every instance that implement Middleware and every logic that
// is required should be encapsulated here.With respect to Secure headers we do
// have the following:
func (s Middleware_SecureHeaders) UseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for header, content := range s.header {
			w.Header().Add(header, content)
		}
		next.ServeHTTP(w, r)
	})
}

// Each middleware implement Middleware interface methods, and UseMiddleware is
// used to provide changes to requests and responses alike.The following method
// is called upon every instance that implement Middleware and every logic that
// is required should be encapsulated here.With respect to Authentication we do
// have the following:
func (a Middleware_Auth) UseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
			return
		}
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		log.Printf("Username: %s - Password: %s\n", username, password)
		log.Println(r)

		log.Println("This will authenticate... nice.")
		next.ServeHTTP(w, r)
	})
}

// Method used to load common middlewares from  Middlewares variable defined at
// m.common. Provided with the next parameter of type http.Handler, it produces
// the same pointer at return, if slice m.common is empty , no changes are made
func (m Middlewares) UseCommonMiddlewares(next http.Handler) http.Handler {
	for _, middleware := range m.common {
		next = middleware.UseMiddleware(next)
	}
	return next
}
