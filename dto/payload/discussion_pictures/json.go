package discussion_pictures

import ()

type DiscussionPicture struct {
	URL          string `json:"url" form:"url" validate:"required"`
}

type AddMultipleDiscussionPicture struct {
	URL          []string `json:"url" form:"url" validate:"required"`
}

