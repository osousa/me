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

// Each struct will represent a type of middleware
// We can create as many of each. On this app we'll
// have only one Auth Middleware instantiated
//
type SecureHeaders struct {
	name   string
	header map[string]string
}

type Auth struct {
	name string
}

func NewAuth(name string) Auth {
	return Auth{name}
}

func (a Auth) UseMiddleware(next http.Handler) http.Handler {
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

func NewSecureHeaders(name string) SecureHeaders {
	return SecureHeaders{name, nil}
}

func (s *SecureHeaders) SetSecureHeaders() {
	s.header = map[string]string{"Content-Security-Policy": "default-src 'self' osousa.me; connect-src 'self';",
		"Content-Type":                 "text/html;",
		"X-XSS-Protection":             "1; mode=block",
		"Cross-Origin-Resource-Policy": "same-site",
		"X-Content-Type-Options":       "nosniff",
		"X-Frame-Options":              "deny"}
}

func (s SecureHeaders) UseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for header, content := range s.header {
			w.Header().Add(header, content)
		}
		next.ServeHTTP(w, r)
	})
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

func (m *Middlewares) SetAuthMiddleware() {
	auth := NewAuth("auth")
	m.auth = auth
}

func (m *Middlewares) SetCommonMiddlewares() {
	headers := NewSecureHeaders("default")
	headers.SetSecureHeaders()
	m.common = append(m.common, headers)
}

func (m Middlewares) UseCommonMiddlewares(next http.Handler) http.Handler {
	for _, middleware := range m.common {
		next = middleware.UseMiddleware(next)
	}
	return next
}
