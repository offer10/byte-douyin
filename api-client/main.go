package main

import (
	"log"

	"github.com/offer10/byte-douyin/api-client/initialization"
)

func main() {
	initialization.SetupService()
	r := initialization.SetupRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
