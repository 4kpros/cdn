package data

import (
	"github.com/danielgtaylor/huma/v2"
)

type ImageData struct {
	Image huma.FormFile `form:"image" content-type:"image/png,image/jpeg" required:"true" doc:"Select you image"`
}
