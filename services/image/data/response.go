package data

type ImageResponse struct {
	Name string     `json:"name" required:"false" doc:"Asset name" example:"simple-avatar.png"`
	Urls []ImageUrl `json:"urls" required:"false" doc:"Asset URLs" example:"[]"`
}

type ImageUrl struct {
	Name    string `json:"name" required:"false" doc:"Asset variation name" example:"low"`
	Url     string `json:"url" required:"false" doc:"Asset URL for the data" example:""`
	FileUrl string `json:"fileUrl" required:"false" doc:"Asset URL for the file explorer" example:""`
	Size    int64  `json:"size" required:"false" doc:"Size in KiB" example:""`
	Width   int    `json:"width" required:"false" doc:"Width in pixels" example:"128"`
	Height  int    `json:"height" required:"false" doc:"Height in pixels" example:"128"`
}
