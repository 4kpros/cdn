package config

import (
	"cdn/common/constants"
	"cdn/common/helpers"
	"cdn/common/utils"

	"go.uber.org/zap"
)

type OpenAPITemplate struct {
	Redocly   *string
	Scalar    *string
	Stoplight *string
	Swagger   *string
}

const subDir = "/openapi"

var OpenAPITemplates = &OpenAPITemplate{}

// LoadOpenAPITemplates Loads OpenAPI templates from a specified location resources.
func LoadOpenAPITemplates() error {
	var err error
	var errRead error

	// Redocly
	OpenAPITemplates.Redocly, errRead = utils.ReadFileToString(constants.ASSET_TEMPLATES_PATH + subDir + "/redocly.html")
	if errRead != nil {
		err = errRead
		helpers.Logger.Warn(
			"Failed to load OpenAPI Redocly template",
			zap.String("Error", errRead.Error()),
		)
	} else {
		helpers.Logger.Info("OpenAPI template Redocly loaded!")
	}

	// Scalar
	OpenAPITemplates.Scalar, err = utils.ReadFileToString(constants.ASSET_TEMPLATES_PATH + subDir + "/scalar.html")
	if errRead != nil {
		helpers.Logger.Warn(
			"Failed to load OpenAPI Scalar template",
			zap.String("Error", errRead.Error()),
		)
	} else {
		helpers.Logger.Info("OpenAPI template Scalar loaded!")
	}

	// Stoplight
	OpenAPITemplates.Stoplight, errRead = utils.ReadFileToString(constants.ASSET_TEMPLATES_PATH + subDir + "/stoplight.html")
	if errRead != nil {
		err = errRead
		helpers.Logger.Warn(
			"Failed to load OpenAPI Stoplight template",
			zap.String("Error", errRead.Error()),
		)
	} else {
		helpers.Logger.Info("OpenAPI template Stoplight loaded!")
	}

	// Swagger
	OpenAPITemplates.Swagger, errRead = utils.ReadFileToString(constants.ASSET_TEMPLATES_PATH + subDir + "/swagger.html")
	if errRead != nil {
		err = errRead
		helpers.Logger.Warn(
			"Failed to load OpenAPI Swagger template",
			zap.String("Error", errRead.Error()),
		)
	} else {
		helpers.Logger.Info("OpenAPI template Swagger loaded!")
	}

	return err
}
