package data

type UploadImageResponse struct {
	Url    string `json:"url" required:"false" doc:"Asset URL for the data" example:""`
	Width  int    `json:"width" required:"false" doc:"Width in pixels" example:"128"`
	Height int    `json:"height" required:"false" doc:"Height in pixels" example:"128"`
}
