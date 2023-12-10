package main

import (
	"net/http"
	"time"

	"github.com/om205/array-sorter/pkg/server"
)

func main() {
	http.HandleFunc("/process-single", server.ProcessSingleHandler)
	http.HandleFunc("/process-concurrent", server.ProcessConcurrentHandler)

	port := ":8000"
	serverInstance := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	println("Server listening on port", port)
	err := serverInstance.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
