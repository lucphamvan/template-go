package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Quiz struct
type Quiz struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerId   string             `json:"owner_id,omitempty" bson:"owner_id"`
	Code      string             `json:"code,omitempty" bson:"code"`
	IsPublish bool               `json:"is_publish,omitempty" bson:"is_publish"`

	QuestionIds []string    `json:"question_ids,omitempty" bson:"question_ids"`
	Setting     QuizSetting `json:"setting,omitempty" bson:"setting"`

	CreatedAt float64 `json:"created_at,omitempty" bson:"created_at"`
	Deleted   bool    `json:"deleted,omitempty" `
}

type QuizSetting struct {
	Name          string   `json:"name,omitempty" bson:"name" validate:"required"`
	StartTime     float64  `json:"start_time,omitempty" bson:"start_time" validate:"required"`
	EndTime       float64  `json:"end_time,omitempty" bson:"end_time" validate:"required,gtefield=StartTime"`
	Duration      int      `json:"duration,omitempty" bson:"duration" validate:"required"`
	AllowedEmails []string `json:"allowed_emails,omitempty" bson:"allowed_emails"`
}

type QuizAnswer struct {
	Id              primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	QuizId          string             `json:"quiz_id,omitempty" bson:"quiz_id"`
	UserId          string             `json:"user_id,omitempty" bson:"user_id"`
	QuestionAnswers []QuestionAnswer   `json:"question_answers,omitempty" bson:"question_answers"`
}

// struct for Request and Response
type CreateQuizInput struct {
	Setting QuizSetting `json:"setting,omitempty" bson:"setting" validate:"required"`
	OwnerId string      `json:"owner_id,omitempty" bson:"owner_id"`
}

type InsertQuestionInput struct {
	Question CreateQuestionInput `json:"question,omitempty" bson:"question"`
}

type RemoveQuestionInput struct {
	QuestionId string `json:"question_id,omitempty" bson:"question_id"`
}

type GetListQuizzesResponse struct {
	Items []Quiz `json:"items"`
	Total int64  `json:"total"`
}

type CreateAndInsertQuestionToQuizResponse struct {
	Quiz     Quiz     `json:"quiz,omitempty"`
	Question Question `json:"question,omitempty"`
}
