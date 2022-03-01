package service

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/google/uuid"
)

type ImageStore interface {
	// Save saves a given image into a destination and returns the generated
	// uuid of the image that was saved.
	Save(laptopID string, imageType string, imageData io.WriterTo) (string, error)
}

type DiskImageStore struct {
	sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}

type ImageInfo struct {
	LaptopID string
	Type     string
	Path     string
}

func NewDiskImageStore(imageFolder string) *DiskImageStore {
	return &DiskImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}

func (s *DiskImageStore) Save(
	laptopID string,
	imageType string,
	imageData io.WriterTo,
) (string, error) {
	imageID, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	imagePath := filepath.Join(s.imageFolder, fmt.Sprintf("%s%s", imageID, imageType))

	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image: %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	s.Lock()
	defer s.Unlock()

	s.images[imageID.String()] = &ImageInfo{
		LaptopID: laptopID,
		Type:     imageType,
		Path:     imagePath,
	}

	return imageID.String(), nil
}
