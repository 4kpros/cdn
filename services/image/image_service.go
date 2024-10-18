package image

import (
	"cdn/common/constants"
	"cdn/common/utils"
	"cdn/config"
	"cdn/services/image/data"
	"context"
	"net/http"
)

type Service struct {
}

const subDir = "/images"

func NewService() *Service {
	return &Service{}
}

// Create new image
func (service *Service) Create(
	ctx *context.Context,
	option *data.ImageQuery,
	input *data.ImageData,
) (result *data.UploadImageResponse, errCode int, err error) {
	// Read buffer
	buffer, err := utils.ReadMultipartFile(input.Image.File)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("read image")
		return
	}

	// Resize and compress the image.
	bufferResized, size, err := utils.ResizeImage(
		buffer, option.Width, option.Height, option.Quality,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("resize image")
		return
	}

	// Create the file
	fileName, err := utils.SaveFile(bufferResized, constants.ASSET_UPLOADS_PATH+subDir)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("save image")
		return
	}
	return &data.UploadImageResponse{
		Url:    "https://" + config.Env.Hostname + config.Env.ApiGroup + "/images/" + fileName,
		Path:   fileName,
		Width:  size.Width,
		Height: size.Height,
	}, 0, nil
}

// Update existing image
func (service *Service) Update(
	ctx *context.Context,
	url string,
	option *data.ImageQuery,
	input *data.ImageData,
) (result *data.UploadImageResponse, errCode int, err error) {
	isDeleted, err := utils.DeleteFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if isDeleted {
		result, errCode, err = service.Create(ctx, option, input)
		return
	}
	errCode = http.StatusNotFound
	err = constants.HTTP_404_ERROR_MESSAGE("Resource")
	return
}

// Delete image with matching id and return affected rows
func (service *Service) Delete(ctx *context.Context, url string) (result bool, errCode int, err error) {
	result, err = utils.DeleteFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if err != nil || !result {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Resource")
	}
	return
}

// Get image with matching id
func (service *Service) Get(
	ctx *context.Context,
	url string,
	option *data.ImageQuery,
) (result []byte, errCode int, err error) {
	// Read buffer
	buffer, err := utils.ReadFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if err != nil || len(buffer) < 1 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Resource")
	}

	// Resize and compress the image.
	result, _, err = utils.ResizeImage(
		buffer, option.Width, option.Height, option.Quality,
	)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("resize image")
		return
	}
	return
}
