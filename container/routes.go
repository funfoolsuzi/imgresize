package container

import (
	"net/http"
)

func (c *Container) routes() {
	c.router.Path("/").Methods(http.MethodGet).Handler(c.handleIndex())
}
