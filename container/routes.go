package container

import (
	"net/http"
)

func (c *Container) routes() {

	c.router.Use(c.middlewareAccessLog)
	c.router.Path("/").Methods(http.MethodGet).Handler(c.handleIndex())
	c.router.PathPrefix("/originals/").Handler(http.StripPrefix("/originals/", http.FileServer(http.Dir("./originals"))))

	subResized := c.router.PathPrefix("/resized/").Subrouter()
	subResized.Use(c.middlewareResizedValidation)
	subResized.Methods(http.MethodGet).Handler(http.StripPrefix("/resized/", c.handleResize()))
}
