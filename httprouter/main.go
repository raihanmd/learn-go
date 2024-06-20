package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources
var resources embed.FS

type LogMid struct {
	Handler http.Handler
}

func (l *LogMid) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Log Mid")
	l.Handler.ServeHTTP(w, r)
}

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Hello World")
	})

	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Hello %v", p.ByName("name"))
	})

	router.GET("/panic", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("Weladalah..")
	})

	dir, _ := fs.Sub(resources, "resources")

	router.ServeFiles("/static/*filepath", http.FS(dir))

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found Buddy")
	})

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, rcv any) {
		fmt.Fprintf(w, "Panic: %v", rcv)
	}
	// ? by default if method not allowed will return http.Error but you can change it by using this
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "Method Not Allowed Buddy")
	})

	// ? Middleware not implemented by this pkg
	router.Handler("GET", "/log", &LogMid{Handler: router})

	return router
}

func main() {
	router := GetRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
