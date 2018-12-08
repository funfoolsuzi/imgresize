package container

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/funfoolsuzi/reqid"
)

func (c *Container) middlewareAccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		st := time.Now()

		ctx := reqid.AttachReqID(r.Context())
		r = r.WithContext(ctx)
		rID := reqid.GetReqID(ctx)

		log.Println(r.Method, r.URL.Path, rID)

		next.ServeHTTP(w, r)

		log.Printf("Finished %s in %v", rID, time.Since(st))
	})
}

func (c *Container) middlewareResizedValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		reqID := reqid.GetReqID(ctx)

		// validations:
		// 1. name format
		// 2. original exist

		p := r.URL.Path
		p = p[8:] // remove "/resized" in the path
		re := regexp.MustCompile(`(?P<name>.+)_h(?P<height>\d+)_w(?P<width>\d+)\.(?P<suffix>\w+)`)

		if !re.MatchString(p) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Path failed regex test"))
			log.Printf("%s %s failed regex text", reqID, p)
			return
		}

		height := re.ReplaceAllString(p, "${height}")
		width := re.ReplaceAllString(p, "${width}")
		fp := re.ReplaceAllString(p, "${name}.${suffix}")

		if !c.originalsRepo.Exist(fp) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("original image %s doesn't exist.", fp)))
			log.Printf("%s %s doesn't exist in originals", reqID, fp)
			return
		}

		log.Printf("%s original: %s, height: %s, width: %s", reqID, fp, height, width)
		ctx = context.WithValue(ctx, contextKeyOriginal, fp)
		ctx = context.WithValue(ctx, contextKeyHeight, height)
		ctx = context.WithValue(ctx, contextKeyWidth, width)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
