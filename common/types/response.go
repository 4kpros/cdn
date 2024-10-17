package types

type DeletedResponse struct {
	AffectedRows int64 `json:"affectedRows" required:"false" doc:"Number of row affected with this delete" example:"1"`
}
