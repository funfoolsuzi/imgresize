package container

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/funfoolsuzi/imgresize/mock"

	"github.com/golang/mock/gomock"
)

func TestResizedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewContainer()
	originalsRepo := mock.NewMockImageRepo(ctrl)
	resizedRepo := mock.NewMockImageRepo(ctrl)

	c.originalsRepo = originalsRepo
	c.resizedRepo = resizedRepo

	handler, _ := c.middlewareResizedValidation(c.handleResize()).(http.HandlerFunc)
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/resized/hello.jpg", bytes.NewReader([]byte{}))
	w := httptest.NewRecorder()

	handler(w, req)
}
