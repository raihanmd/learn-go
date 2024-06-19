package main

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (l *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before Handler")
	l.Handler.ServeHTTP(w, r)
	fmt.Println("After Handler")
}

type ErrorMiddleware struct {
	Handler http.Handler
}

func (l *ErrorMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, err)
		}
	}()
	l.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/error", func(w http.ResponseWriter, r *http.Request) {
		panic("Ups")
	})
	log := &LogMiddleware{Handler: mux}

	errorHandler := &ErrorMiddleware{Handler: log}

	http.ListenAndServe(":3000", errorHandler)
}
