package data

import "github.com/danielgtaylor/huma/v2"

type DocumentQuery struct {
}

type DocumentData struct {
	Document huma.FormFile `form:"document" required:"true" doc:"Supported files type: txt, pdf, doc, docx, csv, xls, xlsx, ppt, pptx, odp, ods, odt"`
}
