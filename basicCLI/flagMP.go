package main

import (
	"flag"
	"log"
)

var name = flag.String("name", "stranger", "your wonderful name")
var age = flag.Int("age", 0, "Your graaceful age")

func main() {
	flag.Parse()
	log.Printf("Hello %s (%d years), Welcome to the command line world", *name, *age)
}
