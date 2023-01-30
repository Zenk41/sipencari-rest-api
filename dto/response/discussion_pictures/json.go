package discussion_pictures

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/models"
)

type DiscussionPicture struct {
	PictureID    uint      `json:"picture_id"`
	URL          string    `json:"url"`
	DiscussionID string    `json:"discussion_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func DiscussionPictureResponse(dPicture models.DiscussionPicture) *DiscussionPicture {
	return &DiscussionPicture{
		PictureID:    dPicture.PictureID,
		URL:          dPicture.URL,
		DiscussionID: dPicture.DiscussionID,
		CreatedAt:    dPicture.CreatedAt,
		UpdatedAt:    dPicture.UpdatedAt,
	}
}

func DiscussionPicturesResponse(dPictures []models.DiscussionPicture) *[]DiscussionPicture {
	var dPicturesResponse []DiscussionPicture
	for _, picture := range dPictures {
		response := DiscussionPicture{
			PictureID:    picture.PictureID,
			URL:          picture.URL,
			DiscussionID: picture.DiscussionID,
			CreatedAt:    picture.CreatedAt,
			UpdatedAt:    picture.UpdatedAt,
		}
		dPicturesResponse = append(dPicturesResponse, response)
	}
	return &dPicturesResponse
}
