package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Quizzes struct
type Quizzes struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerId   string             `json:"owner_id,omitempty" bson:"owner_id"`
	Code      string             `json:"code,omitempty" bson:"code"`
	IsPublish bool               `json:"is_publish,omitempty" bson:"is_publish"`

	QuestionId []string       `json:"question_id,omitempty" bson:"question_id"`
	Setting    QuizzesSetting `json:"setting,omitempty" bson:"setting"`

	CreatedAt float64 `json:"created_at,omitempty"`
	Deleted   bool    `json:"deleted,omitempty"`
}

type QuizzesSetting struct {
	Name          string   `json:"name,omitempty" bson:"name" validate:"required"`
	StartTime     float64  `json:"start_time,omitempty" bson:"start_time" validate:"required"`
	EndTime       float64  `json:"end_time,omitempty" bson:"end_time" validate:"required"`
	Duration      int      `json:"duration,omitempty" bson:"duration" validate:"required"`
	AllowedEmails []string `json:"allowed_emails,omitempty" bson:"allowed_emails"`
}

type QuizzesAnswer struct {
	Id              primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	QuizzesId       string             `json:"quizzes_id,omitempty" bson:"quizzes_id"`
	UserId          string             `json:"user_id,omitempty" bson:"user_id"`
	QuestionAnswers []QuestionAnswer   `json:"question_answers,omitempty" bson:"question_answers"`
}

// struct for Request and Response
type CreateQuizzesInput struct {
	Setting QuizzesSetting `json:"setting,omitempty" bson:"setting" validate:"required"`
	OwnerId string         `json:"owner_id,omitempty" bson:"owner_id"`
}

type InsertQuestionInput struct {
	Question CreateQuestionInput `json:"question,omitempty" bson:"question"`
}

type RemoveQuestionInput struct {
	QuestionId string `json:"question_id,omitempty" bson:"question_id"`
}
