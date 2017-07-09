package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Server Started")
	server := &http.Server{
		Handler:      createRoutes(),
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  15 * time.Millisecond,
		WriteTimeout: 15 * time.Millisecond,
	}
	log.Fatal(server.ListenAndServe())
}
