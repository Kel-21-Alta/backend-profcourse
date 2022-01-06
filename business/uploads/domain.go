package uploads

import (
	"context"
	"mime/multipart"
)

type Domain struct {
	File        *multipart.FileHeader
	Destination string
	ResultUrl   string
	FileName    string
}

type Repository interface {
	UploadImage(ctx context.Context, header *multipart.FileHeader, destination string) (Domain, error)
}
