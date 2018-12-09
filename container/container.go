package container

import (
	"fmt"
	"net/http"

	"github.com/funfoolsuzi/imgresize/imaginaryclient"
	"github.com/funfoolsuzi/imgresize/repo"
	"github.com/gorilla/mux"
)

// Container is the IoC container for the server
type Container struct {
	router        *mux.Router
	originalsRepo repo.ImageRepo
	resizedRepo   repo.ImageRepo
	imageResizer  imaginaryclient.Resizer
}

// NewContainer creates an initialized instance of Container
func NewContainer() *Container {
	c := &Container{
		router:        mux.NewRouter(),
		originalsRepo: repo.NewImageRepository("originals"),
		resizedRepo:   repo.NewImageRepository("resized"),
		imageResizer:  imaginaryclient.NewImaginaryClient("http://imaginary:9000"),
	}

	return c
}

// StartListening will make the server to start listening http traffic
func (c *Container) StartListening(addr string) error {
	c.routes()

	s := http.Server{
		Addr:    addr,
		Handler: c.router,
	}

	err := s.ListenAndServe()
	if err != nil {
		return fmt.Errorf("Failed to start server. %v", err)
	}

	return nil
}
