package main

import (
	"log"

	"github.com/levigross/grequests"
)

func main() {
	res, err := grequests.Get("http://httpbin.org/get", nil)

	if err != nil {
		log.Fatalln("Unable to make request: ", err)
	}

	var returnData map[string]interface{}
	res.JSON(&returnData)
	log.Println(returnData)
}
