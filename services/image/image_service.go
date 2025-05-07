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
	result = &data.UploadImageResponse{
		Url:    config.Env.Hostname + config.Env.ApiGroup + "/images/" + fileName,
		Path:   fileName,
		Width:  size.Width,
		Height: size.Height,
	}
	return
}

// Create new images
func (service *Service) CreateMultiple(
	ctx *context.Context,
	option *data.ImageQuery,
	input *data.MultipleImageData,
) (result *data.UploadMultipleImageResponse, errCode int, err error) {
	if !(input != nil && input.Images != nil && len(input.Images) > 0) {
		return
	}
	images := make([]data.UploadImageResponse, len(input.Images))
	for i := 0; i < len(input.Images); i++ {
		tmpImg, _, _ := service.Create(ctx, option, &input.Images[i])
		if tmpImg != nil {
			images[i] = *tmpImg
		}
	}
	result.UploadImageResponse = images
	return
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
	err = constants.HTTP_404_ERROR_MESSAGE("resource")
	return
}

// Delete image with matching id and return affected rows
func (service *Service) Delete(ctx *context.Context, url string) (result bool, errCode int, err error) {
	result, err = utils.DeleteFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if err != nil || !result {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("resource")
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
		err = constants.HTTP_404_ERROR_MESSAGE("resource")
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
