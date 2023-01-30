package comment_pictures

import ()

type CommentPicture struct {
	URL string `json:"url" form:"url" validate:"required"`
}

type AddMultipleDiscussionPicture struct {
	URL []string `json:"url" form:"url" validate:"required"`
}
