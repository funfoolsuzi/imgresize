package container

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

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
		ctx := r.Context()
		reqID := reqid.GetReqID(ctx)

		p := r.URL.Path[8:]

		resizedImage, err := c.resizedRepo.Get(p)

		if err != nil && err != repo.ErrorNonExist {
			// error occurred but not because file doesn't exist

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Failed to retrieve resized image %s", p)))
			log.Printf("%s failed to retrieve resized image %s. %v", reqID, p, err)
			return

		} else if err == nil {
			// no error

			w.Header().Set("Content-Type", "image/*") // TODO: be more specific
			w.Write(resizedImage)
			log.Printf("%s successfully retrieved resized image %s", reqID, p)
			return
		}

		// error because of image not exist in resized
		// use imaginary to save
		log.Printf("%s resized image %s doesn't exist yet. proceed to create", reqID, p)
		originalImage, err := c.originalsRepo.Get(getOriginalPath(ctx))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error getting original image"))
			log.Printf("%s failed to retrieve from original path that confirmed in middleware. %v", reqID, err)
			return
		}

		url := fmt.Sprintf("http://imaginary:9000/resize?width=%s&height=%s", getWidth(ctx), getHeight(ctx))
		reqResize, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(originalImage))
		reqResize.Header.Set("Content-Type", "image/*")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error resizng image"))
			log.Printf("%s failed to create request to imaginary. %v", reqID, err)
			return
		}

		log.Printf("%s sending http resize request to %s", reqID, url)

		resResize, err := http.DefaultClient.Do(reqResize)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error resizng image"))
			log.Printf("%s failed http request to imaginary. %v", reqID, err)
			return
		}
		defer resResize.Body.Close()
		log.Printf("%s http resp from imaginary %d Content-Length:%s", reqID, resResize.StatusCode, resResize.Header.Get("Content-Length"))

		resizedBytes, err := ioutil.ReadAll(resResize.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error resizng image"))
			log.Printf("%s failed to read from imaginary response body. %v", reqID, err)
			return
		}

		re := regexp.MustCompile(`(?P<name>.+)\.(?P<suffix>[^\.]+)$`)
		newName := re.ReplaceAllString(getOriginalPath(ctx), fmt.Sprintf("${name}_h%s_w%s.${suffix}", getHeight(ctx), getWidth(ctx)))
		err = c.resizedRepo.Create(resizedBytes, newName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error resizng image"))
			log.Printf("%s failed to save newly created image. %v", reqID, err)
			return
		}
		log.Printf("%s successfully save image", reqID)

		w.Write(resizedBytes)
	})
}
