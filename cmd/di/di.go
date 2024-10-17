package di

import (
	"cdn/cmd/api"
	"cdn/services/image"
)

// Inject all dependencies
func InjectDependencies() {
	// Storage
	api.Controllers.ImageController = image.NewController(
		image.NewService(),
	)
}
