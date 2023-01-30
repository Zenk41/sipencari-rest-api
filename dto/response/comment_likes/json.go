package comment_likes

import (
	"time"

	resUser "github.com/Zenk41/sipencari-rest-api/dto/response/users"
	"github.com/Zenk41/sipencari-rest-api/models"
)

type CommentLike struct {
	UserID    string       `json:"user_id"`
	User      resUser.User `json:"user"`
	CommentID string       `json:"comment_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func CommentLikeResponse(cLike models.CommentLike) *CommentLike {
	return &CommentLike{
		UserID:    cLike.UserID,
		CommentID: cLike.CommentID,
		User: resUser.User{
			UserID:  cLike.User.UserID,
			Name:    cLike.User.Name,
			Email:   cLike.User.Email,
			Picture: cLike.User.Picture,
		},
		CreatedAt: cLike.CreatedAt,
		UpdatedAt: cLike.UpdatedAt,
	}
}

func CommentLikesResponse(cLikes []models.CommentLike) *[]CommentLike {
	var cLikesResponse []CommentLike
	for _, like := range cLikes {
		response := CommentLike{
			UserID:    like.UserID,
			CommentID: like.CommentID,
			User: resUser.User{
				UserID:  like.User.UserID,
				Name:    like.User.Name,
				Email:   like.User.Email,
				Picture: like.User.Picture,
			},
			CreatedAt: like.CreatedAt,
			UpdatedAt: like.UpdatedAt,
		}
		cLikesResponse = append(cLikesResponse, response)
	}
	return &cLikesResponse
}
