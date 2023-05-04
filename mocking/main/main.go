package main

import (
	"mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{MyDuration: 1 * time.Second, MySleep: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
