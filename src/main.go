package main

import (
	"database/sql"
	"fmt"
	"k8s.io/kubernetes/pkg/kubelet/configmap"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello there " + message
	w.Write([]byte(message))
}

func main() {
	fmt.Printf("Running on port 8080")
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
