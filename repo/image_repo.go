package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// ErrorNonExist is the error for file not existing
var ErrorNonExist = fmt.Errorf("file already exist")

// ImageRepo is ...
type ImageRepo interface {
	Create()
	Get(path string) ([]byte, error)
	Exist(p string) bool
}

// ImageRepository is the repository to store resized images
type ImageRepository struct {
	parentDir string
}

// NewImageRepository creates an instance of ImageRepository
func NewImageRepository(parentDir string) *ImageRepository {
	return &ImageRepository{
		parentDir,
	}
}

// Create will create an image in repo
func (ir *ImageRepository) Create() {

}

// Get will retrieve an image by path
func (ir *ImageRepository) Get(p string) ([]byte, error) {

	data, err := ioutil.ReadFile(path.Join(".", ir.parentDir, p))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrorNonExist
		}
		return nil, fmt.Errorf("Couldn't read file. %v", err)
	}

	return data, nil
}

// Exist check if a file exists
func (ir *ImageRepository) Exist(p string) bool {
	_, err := os.Open(path.Join(".", ir.parentDir, p))
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
