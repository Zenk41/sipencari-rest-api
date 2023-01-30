package feedback

type Feedback struct {
	Reaction string `json:"reaction" form:"reaction" validate:"required"`
	Review   string `json:"review" form:"review" validate:"required"`
}
