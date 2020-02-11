package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	log.Println("Electric Boogaloo")

	log.SetFlags(0)

	for i := 0; i < 100; i += 1 {
		go log.Println(i)
	}

	for i := 0; i < 100; i += 1 {
		go fmt.Println(i)
	}

	log.SetOutput(ioutil.Discard)

	log.Println("Electric Bugaloo 2")

	defer log.Println("Will not be logged")
	log.Fatal("Exit")
}
