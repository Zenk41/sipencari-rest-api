package comment_reactions

import ()

type CommentReaction struct {
	Helpful   string   `json:"helpful" form:"helpful"  validate:"required"`
}
