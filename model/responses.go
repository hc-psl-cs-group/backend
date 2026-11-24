package model

import (
	"time"

	"gorm.io/gorm"
)

type RecordResponseRequest struct {
	QuestionSetID string            `json:"questionSetId" example:"my-id" maxLength:"50" minLength:"3" validate:"required" binding:"required,min=3,max=50,printascii"`
	Answers       map[string]string `json:"answers" validate:"required" binding:"required,dive"`
}

type ResponseSet struct {
	ID            uint `gorm:"primaryKey"`
	UserID        uint
	QuestionSetID string
	CreatedAt     time.Time
	Answers       []QuestionAnswer
}

type QuestionAnswer struct {
	ID            uint `gorm:"primaryKey"`
	ResponseSetID uint
	QuestionID    string
	Answer        string
}

func (c *DBConnector) RecordResponse(req RecordResponseRequest, userid uint) (uint, bool) {
	responseSet := ResponseSet{
		QuestionSetID: req.QuestionSetID,
		UserID:        userid,
		CreatedAt:     time.Now(),
		Answers:       []QuestionAnswer{},
	}
	for questionId, answer := range req.Answers {
		responseSet.Answers = append(responseSet.Answers, QuestionAnswer{
			QuestionID: questionId,
			Answer:     answer,
		})
	}

	result := gorm.WithResult()
	err := gorm.G[ResponseSet](c.DB, result).Create(ctx, &responseSet)
	if err != nil {
		return 0, false
	}
	return responseSet.ID, true
}
