package container

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/funfoolsuzi/imgresize/repo"

	"github.com/funfoolsuzi/imgresize/mock"

	"github.com/golang/mock/gomock"
)

func TestResizedHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := NewContainer()
	originalsRepo := mock.NewMockImageRepo(ctrl)
	resizedRepo := mock.NewMockImageRepo(ctrl)
	resizer := mock.NewMockResizer(ctrl)

	c.originalsRepo = originalsRepo
	c.resizedRepo = resizedRepo
	c.imageResizer = resizer
	handler, _ := c.middlewareResizedValidation(c.handleResize()).(http.HandlerFunc)

	// case 1: original image exists; resized doesn't exist; create resized image; save resized image
	req := httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/resized/hello_h20_w20.jpg", bytes.NewReader([]byte{}))
	w := httptest.NewRecorder()

	originalsRepo.EXPECT().Exist("/hello.jpg").Return(true)
	originalsRepo.EXPECT().Get("/hello.jpg").Return([]byte("original_hello"), nil)
	resizedRepo.EXPECT().Get("/hello_h20_w20.jpg").Return(nil, repo.ErrorNonExist)
	resizedRepo.EXPECT().Create([]byte("resized_hello"), "/hello_h20_w20.jpg").Return(nil)
	resizer.EXPECT().Resize([]byte("original_hello"), "20", "20").Return([]byte("resized_hello"), nil)

	handler(w, req)

	resized := w.Body.String()
	if resized != "resized_hello" {
		t.Errorf("result body(%s) doesn't match", resized)
	}

	// case 2: original image exists; resized image exists;
	req = httptest.NewRequest(http.MethodGet, "http://127.0.0.1:8080/resized/cube/square_h30_w30.png", bytes.NewReader([]byte{}))
	w = httptest.NewRecorder()

	originalsRepo.EXPECT().Exist("/cube/square.png").Return(true)
	resizedRepo.EXPECT().Get("/cube/square_h30_w30.png").Return([]byte("resized_square"), nil)

	handler(w, req)
	resized = w.Body.String()
	if resized != "resized_square" {
		t.Errorf("result body(%s) doesn't match", resized)
	}

	// case 3...: TODO: all other cases
}
