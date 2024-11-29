package document

import (
	"cdn/common/constants"
	"context"
	"fmt"
	"net/http"

	"cdn/common/types"
	"cdn/services/document/data"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterEndpoints(
	humaApi *huma.API,
	controller *Controller,
) {
	var endpointConfig = types.APIEndpointConfig{
		Group: "/documents",
		Tag:   []string{"Documents"},
	}

	// Upload document
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "post-document",
			Summary:     "Upload document",
			Description: "Upload new document.",
			Method:      http.MethodPost,
			Path:        endpointConfig.Group,
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SECURITY_AUTH_NAME: {}}, // Used to require authentication
			},
			MaxBodyBytes:  (1024 * 1000) * 100, // (1 KiB * 1000) * 5 = 5MiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden},
		},
		func(
			ctx context.Context,
			input *struct {
				data.DocumentQuery
				RawBody huma.MultipartFormFiles[data.DocumentData]
			},
		) (*struct{ Body *data.UploadDocumentResponse }, error) {
			result, errCode, err := controller.Create(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error())
			}
			return &struct{ Body *data.UploadDocumentResponse }{Body: result}, nil
		},
	)

	// Update document
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "update-document",
			Summary:     "Update document",
			Description: "Update existing document.",
			Method:      http.MethodPut,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SECURITY_AUTH_NAME: {}}, // Used to require authentication
			},
			MaxBodyBytes:  (1024 * 1000) * 100, // (1 KiB * 1000) * 5 = 5MiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				types.AssetUrl
				data.DocumentQuery
				RawBody huma.MultipartFormFiles[data.DocumentData]
			},
		) (*struct{ Body *data.UploadDocumentResponse }, error) {
			result, errCode, err := controller.Update(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error())
			}
			return &struct{ Body *data.UploadDocumentResponse }{Body: result}, nil
		},
	)

	// Delete document
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID: "delete-document",
			Summary:     "Delete document",
			Description: "Delete existing document.",
			Method:      http.MethodDelete,
			Path:        fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:        endpointConfig.Tag,
			Security: []map[string][]string{
				{constants.SECURITY_AUTH_NAME: {}}, // Used to require authentication
			},
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				types.AssetUrl
			},
		) (*struct{ Body *types.DeletedResponse }, error) {
			result, errCode, err := controller.Delete(&ctx, input)
			if err != nil {
				return nil, huma.NewError(errCode, err.Error())
			}
			return &struct{ Body *types.DeletedResponse }{Body: &types.DeletedResponse{Deleted: result}}, nil
		},
	)

	// Get document
	huma.Register(
		*humaApi,
		huma.Operation{
			OperationID:   "get-document",
			Summary:       "Get document",
			Description:   "Return existing document data",
			Method:        http.MethodGet,
			Path:          fmt.Sprintf("%s/{url}", endpointConfig.Group),
			Tags:          endpointConfig.Tag,
			MaxBodyBytes:  1024, // 1 KiB
			DefaultStatus: http.StatusOK,
			Errors:        []int{http.StatusInternalServerError, http.StatusBadRequest, http.StatusForbidden, http.StatusNotFound},
		},
		func(
			ctx context.Context,
			input *struct {
				types.AssetUrl
				data.DocumentQuery
			},
		) (
			*huma.StreamResponse,
			error,
		) {
			result, errCode, err := controller.Get(&ctx, input)
			return &huma.StreamResponse{
				Body: func(ctx huma.Context) {
					// Add response headers
					ctx.SetHeader("Content-Type", "document/webp")
					ctx.SetHeader("Content-Length", fmt.Sprint(len(result)))

					// Check errors
					if err != nil {
						ctx.SetStatus(errCode)
						return
					}
					if len(result) < 1 {
						ctx.SetStatus(http.StatusNotFound)
						return
					}

					// Write some data to the stream.
					writer := ctx.BodyWriter()
					_, err = writer.Write(result)
					if err != nil {
						ctx.SetStatus(http.StatusNoContent)
						return
					}
				},
			}, nil
		},
	)
}
