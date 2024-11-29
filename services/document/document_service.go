package document

import (
	"cdn/common/constants"
	"cdn/common/utils"
	"cdn/config"
	"cdn/services/document/data"
	"context"
	"fmt"
	"net/http"
)

type Service struct {
}

const subDir = "/documents"

var documentTypes = []string{
	"text/plain",
	"application/pdf",

	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.template",

	"text/csv",
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
	invalidInput := true
	for _, inputType := range documentTypes {
		if fileType == inputType {
			invalidInput = false
			break
		}
	}
	if invalidInput {
		errCode = http.StatusBadRequest
		err = fmt.Errorf("%s", fmt.Sprintf("Unsupported file type! Please enter valid information."))
	}

	// Create the file
	fileName, err := utils.SaveFile(buffer, constants.ASSET_UPLOADS_PATH+subDir)
	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("save document")
		return
	}
	return &data.UploadDocumentResponse{
		Url:  "https://" + config.Env.Hostname + config.Env.ApiGroup + "/documents/" + fileName,
		Path: fileName,
	}, 0, nil
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
	err = constants.HTTP_404_ERROR_MESSAGE("Resource")
	return
}

// Delete document with matching id and return affected rows
func (service *Service) Delete(ctx *context.Context, url string) (result bool, errCode int, err error) {
	result, err = utils.DeleteFile(constants.ASSET_UPLOADS_PATH + subDir + "/" + url)
	if err != nil || !result {
		errCode = http.StatusNotFound
		err = constants.HTTP_404_ERROR_MESSAGE("Resource")
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
		err = constants.HTTP_404_ERROR_MESSAGE("Resource")
	}

	if err != nil {
		errCode = http.StatusInternalServerError
		err = constants.HTTP_500_ERROR_MESSAGE("resize document")
		return
	}
	return
}
