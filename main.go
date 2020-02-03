package main

import (
	db "db"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello there " + message
	w.Write([]byte(message))
}

func main() {
	fmt.Printf("Running on port 8080\n")
	fmt.Printf("DB Name: ")
	fmt.Printf(os.Getenv("POSTGRES_DB"))
	fmt.Printf("\n")
	dbConnect()
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
