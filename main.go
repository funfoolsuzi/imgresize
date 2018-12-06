package main

import (
	"github.com/funfoolsuzi/imgresize/container"
)

func main() {
	s := container.NewContainer()

	err := s.StartListening(":8080")
	if err != nil {
		panic(err)
	}
}
