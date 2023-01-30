package discussion_locations

import ()

type DiscussionLocation struct {
	Lat          float64 `json:"lat"  form:"lat" validate:"required"`
	Lng          float64 `json:"lng"  form:"lng" validate:"required"`
}


