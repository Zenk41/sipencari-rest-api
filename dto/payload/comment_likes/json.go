package comment_likes

import ()

type CommentLike struct {
	CommentID     string `json:"comment_id"  form:"comment_id" validate:"required"`
}
