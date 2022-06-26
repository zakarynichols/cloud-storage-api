package video_stuff

import (
	"context"
	"mime/multipart"
)

type UploadService interface {
	Upload(ctx context.Context, file multipart.File) error
}
