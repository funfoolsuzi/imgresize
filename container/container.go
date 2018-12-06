package container

import (
	"github.com/funfoolsuzi/imgresize/repo"
	"github.com/gorilla/mux"
)

// Container is the IoC container for the server
type Container struct {
	router mux.Router
	repo   repo.ImageRepo
}
