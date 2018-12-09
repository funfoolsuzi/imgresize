package imaginaryclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Resizer can resize image
type Resizer interface {
	Resize(originalImage []byte, width, height string) ([]byte, error)
}

// ImaginaryClient is used to call imaginary server for image processing
type ImaginaryClient struct {
	url string
}

// NewImaginaryClient creates an instance of ImaginaryClient
func NewImaginaryClient(imaginaryURL string) *ImaginaryClient {
	return &ImaginaryClient{
		url: imaginaryURL,
	}
}

// Resize sends a http request with original image to imaginary to get it resized
func (ic *ImaginaryClient) Resize(originalImage []byte, width, height string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/resize?width=%s&height=%s", ic.url, width, height), bytes.NewReader(originalImage))
	if err != nil {
		return nil, fmt.Errorf("Faield to create http request for imaginary. %v", err)
	}
	req.Header.Set("Content-Type", "image/*")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to send http request to imaginary. %v", err)
	}
	defer res.Body.Close()

	resizedBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read from imaginary response body. %v", err)
	}

	return resizedBytes, nil
}
