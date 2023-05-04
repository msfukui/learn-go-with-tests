package main

import (
	"os"
	"time"

	"learn-go-with-tests/mocking/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{MyDuration: 1 * time.Second, MySleep: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
