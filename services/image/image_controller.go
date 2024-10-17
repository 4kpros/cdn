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
		RawBody huma.MultipartFormFiles[data.ImageData]
	},
) (result *data.ImageResponse, errCode int, err error) {
	result, errCode, err = controller.Service.Create(ctx, input.RawBody.Data())
	return
}

func (controller *Controller) Update(
	ctx *context.Context,
	input *struct {
		types.AssetUrl
		RawBody huma.MultipartFormFiles[data.ImageData]
	},
) (result *data.ImageResponse, errCode int, err error) {

	result, errCode, err = controller.Service.Update(ctx, input.AssetUrl.Url, nil)
	return
}

func (controller *Controller) Delete(
	ctx *context.Context,
	input *struct {
		types.AssetUrl
	},
) (result int64, errCode int, err error) {
	affectedRows, errCode, err := controller.Service.Delete(ctx, input.AssetUrl.Url)
	if err != nil {
		return
	}
	result = affectedRows
	return
}

func (controller *Controller) Get(
	ctx *context.Context,
	input *struct {
		types.AssetUrl
	},
) (result []byte, errCode int, err error) {
	storage, errCode, err := controller.Service.Get(ctx, input.AssetUrl.Url)
	if err != nil {
		return
	}
	result = storage
	return
}
