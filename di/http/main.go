package main

import (
	"net/http"

	"learn-go-with-tests/di/di"
)

func MyGreeterHundler(w http.ResponseWriter, r *http.Request) {
	di.Greet(w, "world")
}

func main() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHundler))
}
