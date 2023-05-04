package main

import (
	"di"
	"os"
)

func main() {
	di.Greet(os.Stdout, "Elodie")
}
