package comment_pictures

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/models"
)

type CommentPicture struct {
	PictureID uint      `json:"picture_id"`
	URL       string    `json:"url"`
	CommentID string    `json:"comment_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CommentPictureResponse(cPicture models.CommentPicture) *CommentPicture {
	return &CommentPicture{
		PictureID: cPicture.PictureID,
		URL:       cPicture.URL,
		CommentID: cPicture.CommentID,
		CreatedAt: cPicture.CreatedAt,
		UpdatedAt: cPicture.UpdatedAt,
	}
}

func CommentPicturesResponse(cPictures []models.CommentPicture) *[]CommentPicture {
	var cPicturesResponse []CommentPicture
	for _, picture := range cPictures {
		response := CommentPicture{
			PictureID: picture.PictureID,
			URL:       picture.URL,
			CommentID: picture.CommentID,
			CreatedAt: picture.CreatedAt,
			UpdatedAt: picture.UpdatedAt,
		}
		cPicturesResponse = append(cPicturesResponse, response)
	}
	return &cPicturesResponse
}
