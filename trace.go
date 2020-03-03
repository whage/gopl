package main

import (
	"time"
	"log"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses!
	// ... lots of work here ...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}

