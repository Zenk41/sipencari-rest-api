package comments

import ()

type Comment struct {
	Message         string   `json:"message"  form:"message" validate:"required"`
	CommentPictures []string `json:"comment_pictures" form:"comment_pictures" `
	// ParrentComment  string   `json:"parrent_comment"  form:"comment_pictures"`
}
type UpdateComment struct {
	Message         string   `json:"message"  form:"message" validate:"required"`
}
