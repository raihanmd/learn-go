package main

import (
	"fmt"
	"net/http"
)

func GetRoute() *http.ServeMux {
	route := http.NewServeMux()

	route.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == "" {
			fmt.Fprint(w, "Hello, World!")
		} else {
			fmt.Fprintf(w, "Hello, %s!", name)
		}
	})

	return route
}

func main() {
	route := GetRoute()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: route,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running on port 3000")
}
