package main

import(
    "net/http"
    //"log"
)


type Middleware interface {
    UseMiddleware(http.Handler)http.Handler
}

type Middlewares struct{
    name string
    active []Middleware
}

type SecureHeaders struct{
    name string
    header map[string]string
}

func InitMiddlewares(name string) Middlewares {
    return Middlewares{name:name, active:[]Middleware{}}
}

func (m *Middlewares) SetDefaultMiddlewares() {
    headers := NewSecureHeaders("default")
    headers.SetDefaultSecureHeaders()
    m.active = append(m.active, headers)
}

func NewSecureHeaders(name string) SecureHeaders {
    return SecureHeaders{name, nil}
}

func (s *SecureHeaders) SetDefaultSecureHeaders(){
    s.header = map[string]string{   "Content-Security-Policy"       : "default-src 'self' osousa.me; connect-src 'self';",
                                    "Content-Type"                  : "application/json",
                                    "X-XSS-Protection"              : "1; mode=block",
                                    "Cross-Origin-Resource-Policy"  : "same-site",
                                    "X-Content-Type-Options"        : "nosniff",
                                    "X-Frame-Options"               : "deny" }
}

func (s SecureHeaders) UseMiddleware(next http.Handler) http.Handler{
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        for header,content := range s.header{
            w.Header().Add(header,content)
        }
        next.ServeHTTP(w, r)
    })
}

func NewMiddlewares(name string) Middlewares{
    m := InitMiddlewares(name)
    m.SetDefaultMiddlewares()
    return m
}

func (m Middlewares) UseDefaultMiddlewares(next http.Handler) http.Handler{
    for _, middleware := range m.active{
        next = middleware.UseMiddleware(next)
    }
    return next
}