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
		Addr:         "127.0.0.1:8080",
		ReadTimeout:  15 * time.Millisecond,
		WriteTimeout: 15 * time.Millisecond,
	}
	log.Fatal(server.ListenAndServe())
}
