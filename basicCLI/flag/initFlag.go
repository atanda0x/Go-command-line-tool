package main

import (
	"flag"
	"log"
)

var name string
var age int

func init() {
	flag.StringVar(&name, "name", "stranger", "you wounderful name")
	flag.IntVar(&age, "age", 0, "Your graceful age ")
}

func main() {
	flag.Parse()
	log.Printf("Hello %s (%d years), welcome to this cmmand-line world", name, age)
}
