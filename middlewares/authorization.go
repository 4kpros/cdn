package middlewares

import (
	"cdn/common/constants"
	"cdn/common/helpers"
	"cdn/config"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/zap"
)

// AuthMiddleware Handles authentication for API requests.
func AuthMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	var errMessage string
	return func(ctx huma.Context, next func(huma.Context)) {
		// Check if current endpoint require authorization
		isAuthorizationRequired := false
		for _, opScheme := range ctx.Operation().Security {
			if _, ok := opScheme[constants.SECURITY_AUTH_NAME]; ok {
				isAuthorizationRequired = true
				break
			}
		}
		if !isAuthorizationRequired {
			next(ctx)
			return
		}

		// Parse and decode the token
		apiKey := helpers.ExtractApiKeyHeader(&ctx)
		helpers.Logger.Warn(
			"Request api key",
			zap.String("Value: ", apiKey),
		)
		helpers.Logger.Warn(
			"CDN api key",
			zap.String("Value: ", config.Env.ApiKey),
		)
		if len(apiKey) < 1 {
			errMessage = "Missing or bad authorization header! Please enter valid information."
			_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, errMessage, fmt.Errorf("%s", errMessage))
			return
		}

		// Validate the token
		if apiKey == config.Env.ApiKey {
			next(ctx)
			return
		}
		tempErr := constants.HTTP_401_INVALID_TOKEN_ERROR_MESSAGE()
		_ = huma.WriteErr(api, ctx, http.StatusUnauthorized, tempErr.Error())
	}
}
