package main

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/fouched/templ-playground/internal/handlers"
)

const port = ":9080"
func main() {

	// using Go's built in router
	// http.HandleFunc("GET /", handlers.Home)
	// http.HandleFunc("GET /about", handlers.About)
	//
	// fmt.Printf("Starting application on %s\n", port)
	// err := http.ListenAndServe(port, nil)

	
	// using chi
	srv := &http.Server{
		Addr: 		port,
		Handler: 	routes(),
	}

	fmt.Printf("Starting application on %s\n", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
	
}