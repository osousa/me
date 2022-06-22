package main

import (
	"log"
	"net/http"
	"time"
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

// Function to initialize Middlewares struct , update if new implementations of
// Middleware are made , otherwise they will not be taken into effect properly.
func InitMiddlewares(name string) Middlewares {
	return Middlewares{name: name, common: []Middleware{}}
}

// Creates a Middleware repository for easy access all in one place. The struct
// Middlewares has a slice of Middleware to accomodate all the middlewares used
// on each response/request . Other types of Middleware should be created as to
// allow for arbitrary use of each
func NewMiddlewares(name string) Middlewares {
	m := InitMiddlewares(name)
	m.SetCommonMiddlewares()
	return m
}

// Define common middlewares to be used at all times for each request/response.
// Currently there is no way to specify to which of the two or even both to see
// them applied. Future revisions should allow to set this option
func (m *Middlewares) SetCommonMiddlewares() {
	secure_headers := Middleware_SecureHeaders{"secureheaders", nil}
	secure_headers.SetSecureHeaders()
	auth_headers := Middleware_Auth{"auth"}
	m.common = append(m.common, auth_headers)
	m.common = append(m.common, secure_headers)
}

// Add at will security headers, these are added at each outgoing http.Response
// and should provide enhanced security to common flaws, like basic XSS and not
// allowing CORS violations, for example
func (s *Middleware_SecureHeaders) SetSecureHeaders() {
	s.header = map[string]string{
		//"Content-Security-Policy": "default-src 'self' osousa.me; connect-src 'self'; style-src 'self'; script-src 'nonce-dd55ea695fb7c34c29a07ce4e56488ba'; ",
		//"Content-Type":                 "text/html;",
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
		log.Printf("Post from website! r.PostFrom = %v\n", r.PostForm)
		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")
		log.Printf("Username: %s - Password: %s\n", username, password)
		log.Println(r)

		log.Println("printing cookies now!")
		log.Println(len(r.Cookies()))
		if len(r.Cookies()) != 0 {
			for _, cookie := range r.Cookies() {
				log.Println(cookie)
			}
		} else {
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Value: "hello", Name: "os_session", Expires: expiration}
			log.Println(cookie)
			http.SetCookie(w, &cookie)
		}

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
