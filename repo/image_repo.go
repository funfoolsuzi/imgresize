package repo

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

// ErrorNonExist is the error for file not existing
var ErrorNonExist = fmt.Errorf("file already exist")

// ImageRepo is ...
type ImageRepo interface {
	Create(imageBytes []byte, p string) error
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
func (ir *ImageRepository) Create(imageBytes []byte, p string) error {

	re := regexp.MustCompile(`(?P<dir>.+/)[^/]+$`)
	dir := re.ReplaceAllString(p, "${dir}")
	err := os.MkdirAll(path.Join(".", ir.parentDir, dir), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to mkdir for save image.%v", err)
	}

	err = ioutil.WriteFile(path.Join(".", ir.parentDir, p), imageBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write to image file. %v", err)
	}

	return nil
}

// Get will retrieve an image by path
func (ir *ImageRepository) Get(p string) ([]byte, error) {

	data, err := ioutil.ReadFile(path.Join(".", ir.parentDir, p))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrorNonExist
		}
		return nil, fmt.Errorf("couldn't read file. %v", err)
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
