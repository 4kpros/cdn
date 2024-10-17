package middlewares

import (
	"fmt"
	"net/http"
	"slices"
	"strings"

	"cdn/config"

	"github.com/danielgtaylor/huma/v2"
)

func CORSMiddleware(api huma.API) func(huma.Context, func(huma.Context)) {
	return func(ctx huma.Context, next func(huma.Context)) {
		ctx.SetHeader("Access-Control-Allow-Origin", "*")
		ctx.SetHeader("Access-Control-Allow-Credentials", "true")
		ctx.SetHeader("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.SetHeader("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")

		// Check for allowed hosts
		if !isOriginKnown(ctx.Host()) {
			errMessage := "CORS error. Our system detected your request as malicious! Please fix that before."
			_ = huma.WriteErr(api, ctx, http.StatusForbidden, errMessage, fmt.Errorf("%s", errMessage))
			return
		}

		next(ctx)
	}
}

// Utility function for CORS
func isOriginKnown(host string) bool {
	hosts := strings.Split(config.Env.AllowedHosts, ",")
	return slices.Contains(hosts, host)
}
