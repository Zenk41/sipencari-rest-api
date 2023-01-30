package feedbacks

import (
	"time"

	resUser "github.com/Zenk41/sipencari-rest-api/dto/response/users"
	"github.com/Zenk41/sipencari-rest-api/models"
)

type Feedback struct {
	FeedbackID string       `json:"feedback_id"`
	UserID     string       `json:"user_id"`
	User       resUser.User `json:"user"`
	Reaction   string       `json:"reaction"`
	Review     string       `json:"review"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

func FeedbackResponse(feedback models.Feedback) *Feedback {
	return &Feedback{
		UserID: feedback.UserID,
		User: resUser.User{
			UserID:    feedback.User.UserID,
			Name:      feedback.User.Name,
			Email:     feedback.User.Email,
			Picture:   feedback.User.Picture,
			CreatedAt: feedback.CreatedAt,
			UpdatedAt: feedback.UpdatedAt,
		},
		FeedbackID: feedback.FeedbackID,
		Reaction:   feedback.Reaction,
		Review:     feedback.Review,
		CreatedAt:  feedback.CreatedAt,
		UpdatedAt:  feedback.UpdatedAt,
	}
}

func FeedbacksResponse(feedbacks []models.Feedback) *[]Feedback {
	var feedbacksResponse []Feedback
	for _, feedback := range feedbacks {
		response := Feedback{
			UserID: feedback.UserID,
			User: resUser.User{
				UserID:  feedback.User.UserID,
				Name:    feedback.User.Name,
				Email:   feedback.User.Email,
				Picture: feedback.User.Picture,
			},
			FeedbackID: feedback.FeedbackID,
			Reaction:   feedback.Reaction,
			Review:     feedback.Review,
			CreatedAt:  feedback.CreatedAt,
			UpdatedAt:  feedback.UpdatedAt,
		}
		feedbacksResponse = append(feedbacksResponse, response)
	}
	return &feedbacksResponse
}
