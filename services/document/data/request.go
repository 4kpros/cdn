package data

import "github.com/danielgtaylor/huma/v2"

type DocumentQuery struct {
}

type DocumentData struct {
	Document huma.FormFile `form:"document" required:"true" doc:"Select you document"`
}
