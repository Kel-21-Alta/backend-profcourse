package locals

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"profcourse/business/locals"
	controller "profcourse/controllers"
	"profcourse/helpers"
	"profcourse/helpers/randomString"
	"strings"
)

type Locals struct {
}

func (l Locals) UploadImage(ctx context.Context, header multipart.FileHeader, destination string) (locals.Domain, error) {
	domain := locals.Domain{}
	domain.File = header

	src, err := domain.File.Open()

	if err != nil {
		return locals.Domain{}, err
	}
	defer src.Close()

	fileName := domain.File.Filename
	splitFileName := strings.Split(fileName, ".")

	// Mengambil extention lalu cek
	extention := splitFileName[len(splitFileName)-1]
	validExtensionImage := []string{"jpg", "jpeg", "png"}
	if !helpers.CheckItemInSlice(validExtensionImage, extention) {
		return locals.Domain{}, controller.INVALID_FILE
	}
	newFileName := randomString.RandomString(10) + "." + extention

	dstFile, err := os.Create("./public" + destination + domain.Destination + newFileName)
	if err != nil {
		return locals.Domain{}, err
	}
	defer dstFile.Close()

	if _, err := io.Copy(dstFile, src); err != nil {
		return locals.Domain{}, err
	}

	domain.ResultUrl = dstFile.Name()
	domain.FileName = newFileName

	return domain, nil
}

func NewLocalRepository() locals.Repository {
	return &Locals{}
}
