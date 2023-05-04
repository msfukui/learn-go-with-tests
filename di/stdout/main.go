package main

import (
	"os"

	"learn-go-with-tests/di/di"
)

func main() {
	di.Greet(os.Stdout, "Elodie")
}
