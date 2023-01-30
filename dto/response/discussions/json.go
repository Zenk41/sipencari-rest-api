package discussions

import (
	"time"

	"github.com/Zenk41/sipencari-rest-api/constant"
	resComLike "github.com/Zenk41/sipencari-rest-api/dto/response/comment_likes"
	resComPic "github.com/Zenk41/sipencari-rest-api/dto/response/comment_pictures"
	resComReaction "github.com/Zenk41/sipencari-rest-api/dto/response/comment_reactions"
	resCommentD "github.com/Zenk41/sipencari-rest-api/dto/response/comments"
	resLikesD "github.com/Zenk41/sipencari-rest-api/dto/response/discussion_likes"
	resLocD "github.com/Zenk41/sipencari-rest-api/dto/response/discussion_locations"
	resPicD "github.com/Zenk41/sipencari-rest-api/dto/response/discussion_pictures"
	resUser "github.com/Zenk41/sipencari-rest-api/dto/response/users"
	"github.com/Zenk41/sipencari-rest-api/models"
)

type Discussion struct {
	DiscussionID       string                      `json:"discussion_id"`
	Title              string                      `json:"title"`
	Category           string                      `json:"category"`
	Content            string                      `json:"content"`
	DiscussionPictures []resPicD.DiscussionPicture `json:"discussion_pictures"`
	DiscussionLocation resLocD.DiscussionLocation  `json:"discussion_location"`
	DiscussionLikes    []resLikesD.DiscussionLike  `json:"discussion_likes"`
	UserID             string                      `json:"user_id"`
	User               resUser.User                `json:"user"`
	Comments           []resCommentD.Comment       `json:"comments"`
	Status             string                      `json:"status"`
	Privacy            string                      `json:"privacy"`
	LikeTotal          *int                        `json:"like_total"`
	CommentTotal       *int                        `json:"comment_total"`
	CreatedAt          time.Time                   `json:"created_at"`
	UpdatedAt          time.Time                   `json:"updated_at"`
}

func DiscussionResponse(discussion models.Discussion) *Discussion {
	var CommentTotalD int = 0
	var LikeTotalD int = 0
	// picture
	var pictures []resPicD.DiscussionPicture
	picturesData := discussion.DiscussionPictures
	for _, picture := range picturesData {
		pictures = append(pictures, resPicD.DiscussionPicture{
			PictureID:    picture.PictureID,
			URL:          picture.URL,
			DiscussionID: picture.DiscussionID,
			CreatedAt:    picture.CreatedAt,
			UpdatedAt:    picture.UpdatedAt,
		})
	}

	// Like
	var likes []resLikesD.DiscussionLike
	likesData := discussion.DiscussionLikes
	for _, like := range likesData {
		LikeTotalD += 1
		likes = append(likes, resLikesD.DiscussionLike{
			UserID: like.UserID,
			User: resUser.User{
				UserID:  like.User.UserID,
				Name:    like.User.Name,
				Email:   like.User.Email,
				Picture: like.User.Picture,
			},
			DiscussionID: like.DiscussionID,
			CreatedAt:    like.CreatedAt,
			UpdatedAt:    like.UpdatedAt,
		})
	}

	// Comment
	var comments []resCommentD.Comment
	commentsData := discussion.Comments
	for _, comment := range commentsData {
		var ReactionTotalC int = 0
		var LikeTotalC int = 0
		var HelpfulNo int = 0
		var HelpfulYes int = 0
		CommentTotalD += 1
		// Picture in Comment
		var picturesC []resComPic.CommentPicture
		picturesCData := comment.CommentPictures
		for _, picture := range picturesCData {
			picturesC = append(picturesC, resComPic.CommentPicture{
				PictureID: picture.PictureID,
				URL:       picture.URL,
				CommentID: picture.CommentID,
				CreatedAt: picture.CreatedAt,
				UpdatedAt: picture.UpdatedAt,
			})
		}

		// Like in Comment
		var likesC []resComLike.CommentLike
		likesCData := comment.CommentLikes
		for _, like := range likesCData {
			LikeTotalC += 1
			likesC = append(likesC, resComLike.CommentLike{
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

		// Reaction in Comment
		var reactionsC []resComReaction.CommentReaction
		reactionsCData := comment.CommentReactions
		for _, reaction := range reactionsCData {
			ReactionTotalC += 1
			if reaction.Helpful.String() == constant.HelpfulNo.String() {
				HelpfulNo += 1
			} else if reaction.Helpful.String() == constant.HelpfulYes.String() {
				HelpfulYes += 1
			}
			reactionsC = append(reactionsC, resComReaction.CommentReaction{
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

		comments = append(comments, resCommentD.Comment{
			CommentID:       comment.CommentID,
			Message:         comment.Message,
			DiscussionID:    comment.DiscussionID,
			CommentPictures: picturesC,
			CommentLikes:    likesC,
			CommentReaction: reactionsC,
			ParrentComment:  comment.ParrentComment,
			UserID:          comment.UserID,
			User: resUser.User{
				UserID:  comment.UserID,
				Name:    comment.User.Name,
				Picture: comment.User.Picture,
				Email:   comment.User.Email,
			},
			LikeTotal:       &LikeTotalC,
			TotalReaction:   &ReactionTotalC,
			TotalHelpfulYes: &HelpfulYes,
			TotalHelpfulNo:  &HelpfulNo,
			CreatedAt:       comment.CreatedAt,
			UpdatedAt:       comment.UpdatedAt,
		})

	}

	return &Discussion{
		DiscussionID:       discussion.DiscussionID,
		Title:              discussion.Title,
		Category:           string(discussion.Category),
		Content:            discussion.Content,
		DiscussionPictures: pictures,
		DiscussionLocation: resLocD.DiscussionLocation{
			LocationID:   discussion.DiscussionLocation.LocationID,
			Lat:          discussion.DiscussionLocation.Lat,
			Lng:          discussion.DiscussionLocation.Lng,
			LocationName: discussion.DiscussionLocation.LocationName,
			DiscussionID: discussion.DiscussionID,
		},
		DiscussionLikes: likes,
		UserID:          discussion.UserID,
		User: resUser.User{
			UserID:  discussion.User.UserID,
			Name:    discussion.User.Name,
			Picture: discussion.User.Picture,
			Email:   discussion.User.Email,
		},
		Comments:     comments,
		Status:       string(discussion.Status),
		Privacy:      string(discussion.Privacy),
		LikeTotal:    &LikeTotalD,
		CommentTotal: &CommentTotalD,
		CreatedAt:    discussion.CreatedAt,
		UpdatedAt:    discussion.UpdatedAt,
	}
}

func DiscussionsResponse(discussions []models.Discussion) *[]Discussion {
	var discussionsResponse []Discussion
	for _, discussion := range discussions {
		var CommentTotalD int = 0
		var LikeTotalD int = 0

		// picture
		var pictures []resPicD.DiscussionPicture
		picturesData := discussion.DiscussionPictures
		for _, picture := range picturesData {
			pictures = append(pictures, resPicD.DiscussionPicture{
				PictureID:    picture.PictureID,
				URL:          picture.URL,
				DiscussionID: picture.DiscussionID,
				CreatedAt:    picture.CreatedAt,
				UpdatedAt:    picture.UpdatedAt,
			})
		}

		// Like
		var likes []resLikesD.DiscussionLike
		likesData := discussion.DiscussionLikes
		for _, like := range likesData {
			LikeTotalD += 1
			likes = append(likes, resLikesD.DiscussionLike{
				UserID: like.UserID,
				User: resUser.User{
					UserID:  like.User.UserID,
					Name:    like.User.Name,
					Email:   like.User.Email,
					Picture: like.User.Picture,
				},
				DiscussionID: like.DiscussionID,
				CreatedAt:    like.CreatedAt,
				UpdatedAt:    like.UpdatedAt,
			})
		}

		// Comment
		var comments []resCommentD.Comment
		commentsData := discussion.Comments
		for _, comment := range commentsData {
			var ReactionTotalC int = 0
			var LikeTotalC int = 0
			var HelpfulNo int = 0
			var HelpfulYes int = 0
			CommentTotalD += 1
			// Picture in Comment
			var picturesC []resComPic.CommentPicture
			picturesCData := comment.CommentPictures
			for _, picture := range picturesCData {
				picturesC = append(picturesC, resComPic.CommentPicture{
					PictureID: picture.PictureID,
					URL:       picture.URL,
					CommentID: picture.CommentID,
					CreatedAt: picture.CreatedAt,
					UpdatedAt: picture.UpdatedAt,
				})
			}

			// Like in Comment
			var likesC []resComLike.CommentLike
			likesCData := comment.CommentLikes
			for _, like := range likesCData {
				LikeTotalC += 1
				likesC = append(likesC, resComLike.CommentLike{
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

			// Reaction in Comment
			var reactionsC []resComReaction.CommentReaction
			reactionsCData := comment.CommentReactions
			for _, reaction := range reactionsCData {
				ReactionTotalC += 1
				if reaction.Helpful.String() == constant.HelpfulNo.String() {
					HelpfulNo += 1
				} else if reaction.Helpful.String() == constant.HelpfulYes.String() {
					HelpfulYes += 1
				}
				reactionsC = append(reactionsC, resComReaction.CommentReaction{
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

			comments = append(comments, resCommentD.Comment{
				CommentID:       comment.CommentID,
				Message:         comment.Message,
				DiscussionID:    comment.DiscussionID,
				CommentPictures: picturesC,
				CommentLikes:    likesC,
				CommentReaction: reactionsC,
				ParrentComment:  comment.ParrentComment,
				UserID:          comment.UserID,
				User: resUser.User{
					UserID:  comment.UserID,
					Name:    comment.User.Name,
					Picture: comment.User.Picture,
					Email:   comment.User.Email,
				},
				LikeTotal:       &LikeTotalC,
				TotalReaction:   &ReactionTotalC,
				TotalHelpfulYes: &HelpfulYes,
				TotalHelpfulNo:  &HelpfulNo,
				CreatedAt:       comment.CreatedAt,
				UpdatedAt:       comment.UpdatedAt,
			})
		}

		response := Discussion{
			DiscussionID:       discussion.DiscussionID,
			Title:              discussion.Title,
			Category:           string(discussion.Category),
			Content:            discussion.Content,
			DiscussionPictures: pictures,
			DiscussionLocation: resLocD.DiscussionLocation{
				LocationID:   discussion.DiscussionLocation.LocationID,
				Lat:          discussion.DiscussionLocation.Lat,
				Lng:          discussion.DiscussionLocation.Lng,
				LocationName: discussion.DiscussionLocation.LocationName,
				DiscussionID: discussion.DiscussionID,
			},
			DiscussionLikes: likes,
			UserID:          discussion.UserID,
			User: resUser.User{
				UserID:  discussion.User.UserID,
				Name:    discussion.User.Name,
				Picture: discussion.User.Picture,
				Email:   discussion.User.Email,
			},
			Comments:     comments,
			Status:       string(discussion.Status),
			Privacy:      string(discussion.Privacy),
			LikeTotal:    &LikeTotalD,
			CommentTotal: &CommentTotalD,
			CreatedAt:    discussion.CreatedAt,
			UpdatedAt:    discussion.UpdatedAt,
		}
		discussionsResponse = append(discussionsResponse, response)
	}
	return &discussionsResponse
}
