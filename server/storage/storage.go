package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
)

/*
	Cloud storage package for video files using the Cloudinary Upload API.
	If needed can swap out for another third party service.
*/

type StorageService struct {
	cld *cloudinary.Cloudinary
}

func NewStorageService() *StorageService {
	cld, err := cloudinary.NewFromParams(os.Getenv("CLOUD_NAME"), os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	if err != nil {
		log.Fatal(err)
	}

	return &StorageService{cld: cld}
}

func (s *StorageService) Upload(ctx context.Context, file multipart.File) error {
	// Closure around cloudinary upload functionality.
	resp, err := s.cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: uuid.New().String(),
	})

	if err != nil {
		log.Fatal(err)
	}

	err = PrintIndent(resp)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// PrintIndent pretty prints JSON for
// improved readability when debugging
// in the terminal.
func PrintIndent(v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	fmt.Print(string(b))

	return nil
}
