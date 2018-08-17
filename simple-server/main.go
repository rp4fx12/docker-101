package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	defaultPort = "8080"
	defaultName = "world"
)

func main() {
	var port string
	var name string
	if len(os.Args) > 1 && os.Args[1] != "" {
		port = ":" + os.Args[1]
	} else {
		port = ":" + defaultPort
	}
	if len(os.Args) > 2 && os.Args[2] != "" {
		name = os.Args[2]
	} else {
		name = defaultName
	}

	fmt.Printf("Server started on port: %v with name: %v", port, name)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", name)
	})

	log.Fatal(http.ListenAndServe(port, nil))

}
