package image

import (
	"cdn/common/constants"
	"cdn/common/utils"
	"cdn/services/image/data"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

// Create new image
func (service *Service) Create(
	ctx *context.Context,
	input *data.ImageData,
) (result *data.ImageResponse, errCode int, err error) {
	defer func(File multipart.File) {
		err := File.Close()
		if err != nil {
			return
		}
	}(input.Image.File)

	//3. Create a temporary file to our directory
	tempFolderPath := fmt.Sprintf("%s%s", constants.RootPath, "/tempFiles")
	tempFileName := fmt.Sprintf("upload-%s-*.%s", utils.FileNameWithoutExtension(input.Image.Filename), filepath.Ext(input.Image.Filename))

	tempFile, err := os.CreateTemp(tempFolderPath, tempFileName)
	if err != nil {
		return
	}

	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			return
		}
	}(tempFile)

	buffer, err := io.ReadAll(input.Image.File)
	if err != nil {
		return
	}

	_, err = tempFile.Write(buffer)
	if err != nil {
		return
	}
	return &data.ImageResponse{
		Name: fmt.Sprint(input.Image.File),
	}, 0, nil
}

// Update existing image
func (service *Service) Update(ctx *context.Context, url string, data []byte) (result *data.ImageResponse, errCode int, err error) {

	return
}

// Delete image with matching id and return affected rows
func (service *Service) Delete(ctx *context.Context, url string) (affectedRows int64, errCode int, err error) {

	return
}

// Get Return image with matching id
func (service *Service) Get(ctx *context.Context, url string) (result []byte, errCode int, err error) {

	return
}
