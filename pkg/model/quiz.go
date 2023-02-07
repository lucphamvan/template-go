package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// test
type Quizzes struct {
	Id            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerId       string             `json:"owner_id,omitempty" bson:"owner_id"`
	Name          string             `json:"name,omitempty" bson:"name"`
	Code          string             `json:"code,omitempty" bson:"code"`
	AllowedEmails []string           `json:"allowed_emails,omitempty" bson:"allowed_emails"`
	QuestionId    []string           `json:"question_id,omitempty" bson:"question_id"`

	Deleted   bool    `json:"deleted,omitempty"`
	StartTime float64 `json:"start_time,omitempty" bson:"start_time"`
	EndTime   float64 `json:"end_time,omitempty" bson:"end_time"`
	Duration  int     `json:"duration,omitempty" bson:"duration"`

	CreatedAt float64 `json:"created_at,omitempty"`
}

// option answer for a question
type Choice struct {
	Id      primitive.ObjectID `json:"id,omitempty" validate:"required"`
	Content string             `json:"content,omitempty" validate:"required"`
}

// question
type Question struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content          string             `json:"content,omitempty" validate:"required"`
	OwnerId          string             `json:"owner_id,omitempty" bson:"owner_id"`
	Choices          []Choice           `json:"choices,omitempty" bson:"choices" validate:"required,min=1"`
	CorrectChoiceIds []string           `json:"correct_choice_ids,omitempty" bson:"correct_choice_ids" validate:"required,min=1"`
	Deleted          bool               `json:"deleted,omitempty"`

	CreatedAt float64 `json:"created_at,omitempty"`
}

type QuizzesAnswer struct {
	Id              primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	QuizzesId       string             `json:"quizzes_id,omitempty" bson:"quizzes_id"`
	UserId          string             `json:"user_id,omitempty" bson:"user_id"`
	QuestionAnswers []QuestionAnswer   `json:"question_answers,omitempty" bson:"question_answers"`
}

// struct for Question Answer
type QuestionAnswer struct {
	QuestionId       string   `json:"question_id,omitempty" bson:"question_id"`
	CorrectChoiceIds []string `json:"correct_choice_ids,omitempty" bson:"correct_choice_ids"`
}

// struct for Request and Response
type CreateQuizzesInput struct {
	Name      string                `json:"name,omitempty" validate:"required"`
	OwnerId   string                `json:"owner_id,omitempty" bson:"owner_id" validate:"required"`
	Questions []CreateQuestionInput `json:"questions,omitempty" bson:"questions" validate:"required,min=1"`
}

type CreateQuestionInput struct {
	Content          string   `json:"content,omitempty" validate:"required"`
	Choices          []Choice `json:"choices,omitempty" validate:"required,min=1" bson:"choices"`
	CorrectChoiceIds []string `json:"correct_choice_ids,omitempty" validate:"required,min=1" bson:"correct_choice_ids"`
	OwnerId          string   `json:"owner_id,omitempty" bson:"owner_id"`
	Deleted          bool     `json:"deleted,omitempty"`
}

type UpdateQuestionInput struct {
	Content          string   `json:"content,omitempty" validate:"required"`
	Choices          []Choice `json:"choices,omitempty" validate:"required,min=1" bson:"choices"`
	CorrectChoiceIds []string `json:"correct_choice_ids,omitempty" validate:"required,min=1" bson:"correct_choice_ids"`
}

type ListQuestionResponse struct {
	Items []Question `json:"items,omitempty"`
	Total int64      `json:"total,omitempty"`
}
