package discussion_likes

import (
	"time"

	resUser "github.com/Zenk41/sipencari-rest-api/dto/response/users"
	"github.com/Zenk41/sipencari-rest-api/models"
)

type DiscussionLike struct {
	UserID       string       `json:"user_id"`
	User         resUser.User `json:"user"`
	DiscussionID string       `json:"discussion_id"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
}

func DiscussionLikeResponse(dLike models.DiscussionLike) *DiscussionLike {
	return &DiscussionLike{
		UserID:       dLike.UserID,
		DiscussionID: dLike.DiscussionID,
		User: resUser.User{
			UserID:  dLike.User.UserID,
			Name:    dLike.User.Name,
			Email:   dLike.User.Email,
			Picture: dLike.User.Picture,
		},
		CreatedAt: dLike.CreatedAt,
		UpdatedAt: dLike.UpdatedAt,
	}
}

func DiscussionLikesResponse(dLikes []models.DiscussionLike) *[]DiscussionLike {
	var dLikesResponse []DiscussionLike
	for _, like := range dLikes {
		response := DiscussionLike{
			UserID:       like.UserID,
			DiscussionID: like.DiscussionID,
			User: resUser.User{
				UserID:  like.User.UserID,
				Name:    like.User.Name,
				Email:   like.User.Email,
				Picture: like.User.Picture,
			},
			CreatedAt: like.CreatedAt,
			UpdatedAt: like.UpdatedAt,
		}
		dLikesResponse = append(dLikesResponse, response)
	}
	return &dLikesResponse
}
