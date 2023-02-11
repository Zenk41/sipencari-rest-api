package comments

import (
	payload "github.com/Zenk41/sipencari-rest-api/dto/payload/comments"
	response "github.com/Zenk41/sipencari-rest-api/dto/response/comments"
	"github.com/Zenk41/sipencari-rest-api/models"

	repository "github.com/Zenk41/sipencari-rest-api/repositories/comments"
)

type CommentService interface {
	Create(userID string, discussionID string, pictures []string, payload payload.Comment, receiverID string) (response.Comment, error)
	Update(commentID string, payload payload.UpdateComment, receiverID string) (response.Comment, error)
	GetByID(commentID string, receiverID string) (response.Comment, error)
	GetAll(discussionID string, receiverID string) ([]response.Comment, error)
	Delete(commentID string) (bool, error)
}

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) CommentService {
	return &commentService{repository: repository}
}

func (cs *commentService) Create(userID string, discussionID string, pictures []string, payload payload.Comment, receiverID string) (response.Comment, error) {
	var Comment models.Comment
	Comment.Message = payload.Message
	Comment.ParrentComment = payload.ParrentComment
	Comment.UserID = userID
	Comment.DiscussionID = discussionID

	comment, err := cs.repository.Create(Comment, pictures)

	if err != nil {
		return response.Comment{}, err
	}
	return *response.CommentResponse(comment, receiverID), err

}
func (cs *commentService) Update(commentID string, payload payload.UpdateComment, receiverID string) (response.Comment, error) {
	comment, err := cs.repository.GetByID(commentID)
	if err != nil {
		return response.Comment{}, err
	}
	comment.Message = payload.Message
	updatedComment, err := cs.repository.Update(comment, commentID)
	if err != nil {
		return response.Comment{}, err
	}
	return *response.CommentResponse(updatedComment, receiverID), err
}

func (cs *commentService) GetByID(commentID string, receiverID string) (response.Comment, error) {
	comment, err := cs.repository.GetByID(commentID)
	if err != nil {
		return response.Comment{}, err
	}
	return *response.CommentResponse(comment, receiverID), nil
}
func (cs *commentService) GetAll(discussionID string, receiverID string) ([]response.Comment, error) {
	comments, err := cs.repository.GetAll(discussionID)
	if err != nil {
		return []response.Comment{}, err
	}
	return *response.CommentsResponse(comments, receiverID), err
}
func (cs *commentService) Delete(commentID string) (bool, error) {

	isDeleted, err := cs.repository.Delete(commentID)
	if err != nil {
		return isDeleted, err
	}

	return isDeleted, nil
}
