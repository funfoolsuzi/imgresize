package container

import (
	"log"
	"net/http"
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
