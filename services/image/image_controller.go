package image

import (
	"cdn/common/types"
	"cdn/services/image/data"
	"context"
	"github.com/danielgtaylor/huma/v2"
)

type Controller struct {
	Service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{Service: service}
}

func (controller *Controller) Create(
	ctx *context.Context,
	input *struct {
		data.ImageQuery
		RawBody huma.MultipartFormFiles[data.ImageData]
	},
) (result *data.UploadImageResponse, errCode int, err error) {
	result, errCode, err = controller.Service.Create(ctx, &input.ImageQuery, input.RawBody.Data())
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		types.AssetUrl
		data.ImageQuery
		RawBody huma.MultipartFormFiles[data.ImageData]
	},
) (result *data.UploadImageResponse, errCode int, err error) {

	result, errCode, err = controller.Service.Update(ctx, input.AssetUrl.Url, &input.ImageQuery, input.RawBody.Data())
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		types.AssetUrl
	},
) (result bool, errCode int, err error) {
	result, errCode, err = controller.Service.Delete(ctx, input.AssetUrl.Url)
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		types.AssetUrl
		data.ImageQuery
	},
) (result []byte, errCode int, err error) {
	result, errCode, err = controller.Service.Get(ctx, input.AssetUrl.Url, &input.ImageQuery)
	return
}
