package container

import (
	"fmt"
	"log"
	"net/http"

	"github.com/funfoolsuzi/reqid"

	"github.com/funfoolsuzi/imgresize/repo"
)

func (c *Container) handleIndex() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to imgResize"))
	})
}

func (c *Container) handleResize() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqID := reqid.GetReqID(r.Context())

		resizedImage, err := c.resizedRepo.Get(r.URL.Path)

		if err != nil && err != repo.ErrorNonExist {
			// error occurred but not because file doesn't exist

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Failed to retrieve image %s", r.URL.Path)))
			log.Printf("%s failed to retrieve image %s. %v", reqID, r.URL.Path, err)
			return

		} else if err == nil {
			// no error

			w.Header().Set("Content-Type", "image/*") // TODO: be more specific
			w.Write(resizedImage)
			log.Printf("%s successfully retrieved image %s", reqID, r.URL.Path)
		}

		// error because of image not exist in resized
		// use imaginary to save
		// TODO: continue

	})
}
