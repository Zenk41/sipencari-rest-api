package comments

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/constant"
	resLikesC "github.com/Zenk41/sipencari-rest-api/dto/response/comment_likes"
	resPicC "github.com/Zenk41/sipencari-rest-api/dto/response/comment_pictures"
	resReactC "github.com/Zenk41/sipencari-rest-api/dto/response/comment_reactions"
	resUser "github.com/Zenk41/sipencari-rest-api/dto/response/users"
	"github.com/Zenk41/sipencari-rest-api/models"
)

type Comment struct {
	CommentID       string                      `json:"comment_id"`
	Message         string                      `json:"message"`
	DiscussionID    string                      `json:"discussion_id"`
	CommentPictures []resPicC.CommentPicture    `json:"comment_pictures"`
	CommentLikes    []resLikesC.CommentLike     `json:"comment_likes"`
	CommentReaction []resReactC.CommentReaction `json:"comment_reactions"`
	// ParentComment  string                      `json:"parent_comment"`
	UserID          string                      `json:"user_id"`
	User            resUser.User                `json:"user"`
	IsLike          *bool                       `json:"is_like"`
	IsHelpfulYes    *bool                       `json:"is_helpful_yes"`
	IsHelpfulNo     *bool                       `json:"is_helpful_no"`
	LikeTotal       *int                        `json:"like_total"`
	TotalHelpfulYes *int                        `json:"total_helpful_yes"`
	TotalHelpfulNo  *int                        `json:"total_helpful_no"`
	TotalReaction   *int                        `json:"total_reaction"`
	CreatedAt       time.Time                   `json:"created_at"`
	UpdatedAt       time.Time                   `json:"updated_at"`
}

func CommentResponse(comment models.Comment, userID string) *Comment {
	var ReactionTotalC int = 0
	var LikeTotalC int = 0
	var HelpfulNo int = 0
	var HelpfulYes int = 0
	var ICommentLike bool = false
	var CommentReactionYes bool = false
	var CommentReactionNo bool = false

	// Like in Comment
	var likesC []resLikesC.CommentLike
	likesCData := comment.CommentLikes
	for _, like := range likesCData {
		LikeTotalC += 1
		if userID == like.UserID {
			ICommentLike = true
		} 
		likesC = append(likesC, resLikesC.CommentLike{
			UserID: like.UserID,
			User: resUser.User{
				UserID:  like.User.UserID,
				Name:    like.User.Name,
				Email:   like.User.Email,
				Picture: like.User.Picture,
			},
			CommentID: like.CommentID,
			CreatedAt: like.CreatedAt,
			UpdatedAt: like.UpdatedAt,
		})
	}
	// Picture in Comment
	var picturesC []resPicC.CommentPicture
	picturesCData := comment.CommentPictures
	for _, picture := range picturesCData {
		picturesC = append(picturesC, resPicC.CommentPicture{
			PictureID: picture.PictureID,
			URL:       picture.URL,
			CommentID: picture.CommentID,
			CreatedAt: picture.CreatedAt,
			UpdatedAt: picture.UpdatedAt,
		})
	}

	// Reaction in Comment
	var reactionsC []resReactC.CommentReaction
	reactionsCData := comment.CommentReactions
	for _, reaction := range reactionsCData {
		ReactionTotalC += 1
		if reaction.Helpful.String() == constant.HelpfulNo.String() {
			HelpfulNo += 1
		} else if reaction.Helpful.String() == constant.HelpfulYes.String() {
			HelpfulYes += 1
		}
		if userID == reaction.UserID && reaction.Helpful.String() == constant.HelpfulNo.String() {
			CommentReactionNo = true
			CommentReactionYes = false
		} else if userID == reaction.UserID && reaction.Helpful.String() == constant.HelpfulYes.String() {
			CommentReactionNo = false
			CommentReactionYes = true
		}
		reactionsC = append(reactionsC, resReactC.CommentReaction{
			UserID: reaction.UserID,
			User: resUser.User{
				UserID:  reaction.UserID,
				Name:    reaction.User.Name,
				Picture: reaction.User.Picture,
			},
			Helpful:   string(reaction.Helpful),
			CommentID: reaction.CommentID,
			CreatedAt: reaction.CreatedAt,
			UpdatedAt: reaction.UpdatedAt,
		})
	}

	return &Comment{
		CommentID:       comment.CommentID,
		Message:         comment.Message,
		DiscussionID:    comment.DiscussionID,
		CommentPictures: picturesC,
		CommentLikes:    likesC,
		CommentReaction: reactionsC,
		// ParentComment:  comment.ParentComment,
		UserID:          comment.UserID,
		User: resUser.User{
			UserID:  comment.UserID,
			Name:    comment.User.Name,
			Picture: comment.User.Picture,
			Email:   comment.User.Email,
		},
		IsLike:          &ICommentLike,
		IsHelpfulYes:    &CommentReactionYes,
		IsHelpfulNo:     &CommentReactionNo,
		LikeTotal:       &LikeTotalC,
		TotalReaction:   &ReactionTotalC,
		TotalHelpfulYes: &HelpfulYes,
		TotalHelpfulNo:  &HelpfulNo,
		CreatedAt:       comment.CreatedAt,
		UpdatedAt:       comment.UpdatedAt,
	}
}

func CommentsResponse(comments []models.Comment, userID string) *[]Comment {
	// Comment
	var commentsResponse []Comment
	for _, comment := range comments {
		var ReactionTotalC int = 0
		var LikeTotalC int = 0
		var HelpfulNo int = 0
		var HelpfulYes int = 0
		var ICommentLike bool = false
		var CommentReactionYes bool = false
		var CommentReactionNo bool = false

		// Like in Comment
		var likesC []resLikesC.CommentLike
		likesCData := comment.CommentLikes
		for _, like := range likesCData {
			LikeTotalC += 1
			if userID == like.UserID {
				ICommentLike = true
			} 
			likesC = append(likesC, resLikesC.CommentLike{
				UserID: like.UserID,
				User: resUser.User{
					UserID:  like.User.UserID,
					Name:    like.User.Name,
					Email:   like.User.Email,
					Picture: like.User.Picture,
				},
				CommentID: like.CommentID,
				CreatedAt: like.CreatedAt,
				UpdatedAt: like.UpdatedAt,
			})
		}
		// Picture in Comment
		var picturesC []resPicC.CommentPicture
		picturesCData := comment.CommentPictures
		for _, picture := range picturesCData {
			picturesC = append(picturesC, resPicC.CommentPicture{
				PictureID: picture.PictureID,
				URL:       picture.URL,
				CommentID: picture.CommentID,
				CreatedAt: picture.CreatedAt,
				UpdatedAt: picture.UpdatedAt,
			})
		}

		// Reaction in Comment
		var reactionsC []resReactC.CommentReaction
		reactionsCData := comment.CommentReactions
		for _, reaction := range reactionsCData {
			ReactionTotalC += 1
			if reaction.Helpful.String() == constant.HelpfulNo.String() {
				HelpfulNo += 1
			} else if reaction.Helpful.String() == constant.HelpfulYes.String() {
				HelpfulYes += 1
			}
			if userID == reaction.UserID && reaction.Helpful.String() == constant.HelpfulNo.String() {
				CommentReactionNo = true
				CommentReactionYes = false
			} else if userID == reaction.UserID && reaction.Helpful.String() == constant.HelpfulYes.String() {
				CommentReactionNo = false
				CommentReactionYes = true
			}
			reactionsC = append(reactionsC, resReactC.CommentReaction{
				UserID: reaction.UserID,
				User: resUser.User{
					UserID:  reaction.UserID,
					Name:    reaction.User.Name,
					Picture: reaction.User.Picture,
				},
				Helpful:   string(reaction.Helpful),
				CommentID: reaction.CommentID,
				CreatedAt: reaction.CreatedAt,
				UpdatedAt: reaction.UpdatedAt,
			})
		}
		response := Comment{
			CommentID:       comment.CommentID,
			Message:         comment.Message,
			DiscussionID:    comment.DiscussionID,
			CommentPictures: picturesC,
			CommentLikes:    likesC,
			CommentReaction: reactionsC,
			// ParrentComment:  comment.ParrentComment,
			UserID:          comment.UserID,
			User: resUser.User{
				UserID:  comment.UserID,
				Name:    comment.User.Name,
				Picture: comment.User.Picture,
				Email:   comment.User.Email,
			},
			IsLike:          &ICommentLike,
			IsHelpfulYes:    &CommentReactionYes,
			IsHelpfulNo:     &CommentReactionNo,
			LikeTotal:       &LikeTotalC,
			TotalReaction:   &ReactionTotalC,
			TotalHelpfulYes: &HelpfulYes,
			TotalHelpfulNo:  &HelpfulNo,
			CreatedAt:       comment.CreatedAt,
			UpdatedAt:       comment.UpdatedAt,
		}
		commentsResponse = append(commentsResponse, response)
	}
	return &commentsResponse
}
