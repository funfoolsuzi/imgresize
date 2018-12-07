package main

import (
	"log"
	"os"

	"github.com/funfoolsuzi/imgresize/container"
)

func main() {
	s := container.NewContainer()

	// set up log output destination
	f, err := os.Create("./log/app.log")
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)

	err = s.StartListening(":8080")
	if err != nil {
		panic(err)
	}
}
