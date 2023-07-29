package discussions

import ()

type CreateDiscussion struct {
	Title              string   `json:"title" form:"title"  validate:"required"`
	Category           string   `json:"category" form:"category" validate:"required"`
	Content            string   `json:"content" form:"content" validate:"required"`
	DiscussionPictures []string `json:"discussion_pictures" form:"discussion_pictures"`
	Lat                float64  `json:"lat" form:"lat"`
	Lng                float64  `json:"lng" form:"lng"`
	Type               string   `json:"type" form:"type" validate:"required"`
	Status             string   `json:"status" form:"status" validate:"required"`
	Contact            string   `json:"contact" form:"contact" validate:"required"`
	Privacy            string   `json:"privacy" form:"privacy" validate:"required"`
}

type UpdateDiscussion struct {
	Title    string  `json:"title" form:"title" validate:"required"`
	Category string  `json:"category" form:"category" validate:"required"`
	Content  string  `json:"content" form:"content" validate:"required"`
	Lat      float64 `json:"lat" form:"lat" validate:"required"`
	Lng      float64 `json:"lng" form:"lng" validate:"required"`
	Type     string  `json:"type" form:"type" validate:"required"`
	Status   string  `json:"status" form:"status" validate:"required"`
	Contact  string  `json:"contact" form:"contact" validate:"required"`
	Privacy  string  `json:"privacy" form:"privacy" validate:"required"`
}
