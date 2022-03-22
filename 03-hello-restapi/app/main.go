package main

import (
	"log"
	"potato/simple-rest/exports/rest"
)

func main() {
	if err := rest.NewServer().Run(":8080"); err != nil {
		log.Fatalf("Could not run HTTP Server with (%v)", err)
		return
	}
}
