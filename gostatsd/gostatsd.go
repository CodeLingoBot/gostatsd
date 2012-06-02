package main

import (
	"log"
	"github.com/kisielk/gostatsd/statsd"
)

func main() {
	err := statsd.ListenAndServe(":8125", ":8126")
	if err != nil {
		log.Fatal(err)
	}
}