package container

import "net/http"

func (c *Container) handleIndex() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to imgResize"))
	})
}

func (c *Container) handleResize() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
