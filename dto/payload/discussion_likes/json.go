package discussion_likes

import ()

type DiscussionLike struct {
	DiscussionID string `json:"discussion_id"  form:"discussion_id" validate:"required"`
}
