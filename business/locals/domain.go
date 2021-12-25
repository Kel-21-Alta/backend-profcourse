package locals

import (
	"context"
	"mime/multipart"
)

type Domain struct {
	File        multipart.FileHeader
	Destination string
	ResultUrl   string
	FileName    string
}

type Repository interface {
	UploadImage(ctx context.Context, header multipart.FileHeader) (Domain, error)
}
