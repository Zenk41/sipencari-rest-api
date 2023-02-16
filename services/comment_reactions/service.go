package comment_reactions

import (
	"github.com/Zenk41/sipencari-rest-api/constant"
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/comment_reactions"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/comment_reactions"
	"github.com/Zenk41/sipencari-rest-api/models"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/comment_reactions"
)

type ComReactionService interface {
	React(payload payload.CommentReaction, commentID, userID string) (bool, response.CommentReaction, error)
	GetAll(commentID string) ([]response.CommentReaction, error)
	GetByID(commentID, userID string) (response.CommentReaction, error)
}

type comReactionService struct {
	repository repository.ComReactionRepository
}

func NewComReactionService(repository repository.ComReactionRepository) ComReactionService {
	return &comReactionService{repository: repository}
}

func (crs *comReactionService) React(payload payload.CommentReaction, commentID, userID string) (bool, response.CommentReaction, error) {
	react, _ := crs.repository.GetByID(commentID, userID)
	var reaction models.CommentReaction
	if react.UserID == "" {
		reaction.CommentID = commentID
		reaction.UserID = userID
		reaction.Helpful = constant.HelpfulEnum(payload.Helpful)
		reaction, err := crs.repository.Create(reaction)
		if err != nil {
			return false, response.CommentReaction{}, err
		}
		return true, *response.CommentReactionResponse(reaction), nil
	} else {
		if (react.Helpful == constant.HelpfulEnum(payload.Helpful) && react.Helpful == constant.HelpfulYes) ||
			(react.Helpful == constant.HelpfulEnum(payload.Helpful) && react.Helpful == constant.HelpfulNo) {
			isDeleted, err := crs.repository.Delete(commentID, userID)
			if err != nil {
				return isDeleted, response.CommentReaction{}, err
			}
			return isDeleted, response.CommentReaction{}, nil
		}
	}
	react.Helpful = constant.HelpfulEnum(payload.Helpful)
	reaction, err := crs.repository.Update(react, commentID, userID)
	if err != nil {
		return false, response.CommentReaction{}, err
	}

	return true, *response.CommentReactionResponse(reaction), nil
}

func (crs *comReactionService) GetAll(commentID string) ([]response.CommentReaction, error) {
	reactions, err := crs.repository.GetAll(commentID)
	if err != nil {
		return []response.CommentReaction{}, err
	}
	return *response.CommentReactionsResponse(reactions), nil
}
func (crs *comReactionService) GetByID(commentID, userID string) (response.CommentReaction, error) {
	reaction, err := crs.repository.GetByID(commentID, userID)
	if err != nil {
		return response.CommentReaction{}, err
	}
	return *response.CommentReactionResponse(reaction), nil
}
