package comment_reactions

import (
	"time"

	resUser "github.com/Zenk41/sipencari-rest-api/dto/response/users"
	"github.com/Zenk41/sipencari-rest-api/models"
)

type CommentReaction struct {
	UserID    string       `json:"user_id"`
	User      resUser.User `json:"user"`
	Helpful   string       `json:"helpful"`
	CommentID string         `json:"comment_id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

func CommentReactionResponse(cReaction models.CommentReaction) *CommentReaction {
	return &CommentReaction{
		UserID:    cReaction.UserID,
		CommentID: cReaction.CommentID,
		Helpful:   string(cReaction.Helpful),
		User: resUser.User{
			UserID:  cReaction.User.UserID,
			Name:    cReaction.User.Name,
			Email:   cReaction.User.Email,
			Picture: cReaction.User.Picture,
		},
		CreatedAt: cReaction.CreatedAt,
		UpdatedAt: cReaction.UpdatedAt,
	}
}

func CommentReactionsResponse(cReactions []models.CommentReaction) *[]CommentReaction {
	var cReactionsResponse []CommentReaction
	for _, reaction := range cReactions {
		response := CommentReaction{
			UserID:    reaction.UserID,
			CommentID: reaction.CommentID,
			Helpful:   string(reaction.Helpful),
			User: resUser.User{
				UserID:  reaction.User.UserID,
				Name:    reaction.User.Name,
				Email:   reaction.User.Email,
				Picture: reaction.User.Picture,
			},
			CreatedAt: reaction.CreatedAt,
			UpdatedAt: reaction.UpdatedAt,
		}
		cReactionsResponse = append(cReactionsResponse, response)
	}
	return &cReactionsResponse
}
