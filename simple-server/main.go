package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

const (
	default_port = "8080"
	default_name = "world"
)

func main() {
	var port string
	var name string
	if (len(os.Args) > 1 && os.Args[1] != "") {
		port = ":" + os.Args[1]
	} else {
		port = ":" + default_port
	}
	if (len(os.Args) > 2 && os.Args[2] != "") {
		name = os.Args[2]
	} else {
		name = default_name
	}

	fmt.Printf("Server started on port: %v with name: %v", port, name)
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", name)
	})
	
	log.Fatal(http.ListenAndServe(port, nil)) 

}
