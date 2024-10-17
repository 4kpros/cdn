package types

type AssetUrl struct {
	Url string `json:"url" path:"url" required:"true" minLength:"3" doc:"Asset url" example:""`
}
