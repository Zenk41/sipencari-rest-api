package comment_pictures

import ()

type CommentPicture struct {
	URL string `json:"url" form:"url" validate:"required"`
}

type AddMultipleDCommentPicture struct {
	URL []string `json:"url" form:"url" validate:"required"`
}
