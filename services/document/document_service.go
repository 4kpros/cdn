package document

import (
	"cdn/common/constants"
	"cdn/common/utils"
	"cdn/config"
	"cdn/services/document/data"
	"context"
	"fmt"
	"net/http"
	"slices"
)

type Service struct {
}

const subDir = "/documents"

var documentTypes = []string{
	"text/plain",
	"application/pdf",
	"text/csv",

	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.template",

	"application/vnd.ms-excel",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	"application/vnd.openxmlformats-officedocument.spreadsheetml.template",

	"application/vnd.ms-powerpoint",
	"application/vnd.openxmlformats-officedocument.presentationml.presentation",
	"application/vnd.openxmlformats-officedocument.presentationml.template",
	"application/vnd.openxmlformats-officedocument.presentationml.slideshow",

	"application/vnd.oasis.opendocument.presentation", // .odp
	"application/vnd.oasis.opendocument.spreadsheet",  // .ods
	"application/vnd.oasis.opendocument.text",         // .odt
}

func NewService() *Service {
	return &Service{}
}

// Create new document
func (service *Service) Create(
	ctx *context.Context,
	option *data.DocumentQuery,
	input *data.DocumentData,
) (result *data.UploadDocumentResponse, errCode int, err error) {
	// Read buffer
	buffer, err := utils.ReadMultipartFile(input.Document.File)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("read document")
		return
	}

	// Check file type
	fileType := http.DetectContentType(buffer)
	if !slices.Contains(documentTypes, fileType) {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", fmt.Sprintf("Unsupported file type %s! Please enter valid information.", fileType))
		return
	}

	// Create the file
	fileName, err := utils.SaveFile(buffer, constants.ASSET_UPLOADS_PATH+subDir)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("save document")
		return
	}
	result = &data.UploadDocumentResponse{
		Url:  config.Env.Hostname + config.Env.ApiGroup + "/documents/" + fileName,
		Path: fileName,
	}
	return
}

// Update existing document
func (service *Service) Update(
	ctx *context.Context,
	url string,
	option *data.DocumentQuery,
	input *data.DocumentData,
) (result *data.UploadDocumentResponse, errCode int, err error) {
	isDeleted, err := utils.DeleteFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if isDeleted {
		result, errCode, err = service.Create(ctx, option, input)
		return
	}
	errCode = http.StatusNotFound
	err = constants.HTTP_404_ERROR_MESSAGE("resource")
	return
}

// Delete document with matching id and return affected rows
func (service *Service) Delete(ctx *context.Context, url string) (result bool, errCode int, err error) {
	result, err = utils.DeleteFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if err != nil || !result {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("resource")
	}
	return
}

// Get document with matching id
func (service *Service) Get(
	ctx *context.Context,
	url string,
	option *data.DocumentQuery,
) (result []byte, errCode int, err error) {
	// Read buffer
	buffer, err := utils.ReadFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if err != nil || len(buffer) < 1 {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("resource")
	}

	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("resize document")
		return
	}
	return
}
