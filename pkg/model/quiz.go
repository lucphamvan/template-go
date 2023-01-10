package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contest struct {
	Id         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	OwnerId    string             `json:"owner_id,omitempty" bson:"owner_id"`
	Name       string             `json:"name,omitempty" bson:"name"`
	Code       string             `json:"code,omitempty" bson:"code"`
	AllowEmail []string           `json:"allow_email,omitempty" bson:"allow_email"`

	StartTime float64 `json:"start_time,omitempty" bson:"start_time"`
	EndTime   float64 `json:"end_time,omitempty" bson:"end_time"`
	Duration  int     `json:"duration,omitempty" bson:"duration"`

	CreateAt float64 `json:"create_at,omitempty"`
}

type Quiz struct {
	Id         primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty"`
	OwnerId    string             `json:"owner_id,omitempty" bson:"owner_id"`
	QuestionId []string           `json:"question_id,omitempty" bson:"question_id"`
}

// struct for Question
type Question struct {
	Id              primitive.ObjectID   `json:"id,omitempty"  bson:"_id,omitempty"`
	Content         string               `json:"content,omitempty" validate:"required"`
	AnswerOption    []AnswerOption       `json:"answer_option,omitempty" validate:"required" bson:"answer_option"`
	AnswerCorrectId []primitive.ObjectID `json:"answer_correct_id,omitempty" validate:"required" bson:"answer_correct_id"`
}

// struct for Question Option Answer
type AnswerOption struct {
	Id      primitive.ObjectID `json:"id,omitempty" validate:"required"`
	Content string             `json:"content,omitempty" validate:"required"`
}

type QuizAnswer struct {
	Id             primitive.ObjectID `json:"id,omitempty"  bson:"_id,omitempty"`
	QuizId         string             `json:"quiz_id,omitempty" bson:"quiz_id"`
	OwnerId        string             `json:"owner_id,omitempty" bson:"owner_id"`
	ContestId      string             `json:"contest_id,omitempty" bson:"contest_id"`
	QuestionAnswer []QuestionAnswer   `json:"question_answer,omitempty" bson:"question_answer"`
}

// struct for Question Answer
type QuestionAnswer struct {
	Question        Question             `json:"question,omitempty"`
	AnswerCorrectId []primitive.ObjectID `json:"answer_correct_id,omitempty" bson:"answer_correct_id"`
}

type CreateQuizInput struct {
	Name     string                `json:"name,omitempty"`
	OwnerId  string                `json:"owner_id,omitempty" bson:"owner_id"`
	Question []CreateQuestionInput `json:"question,omitempty" bson:"question"`
}

type CreateQuestionInput struct {
	Content         string               `json:"content,omitempty" validate:"required"`
	AnswerOption    []AnswerOption       `json:"answer_option,omitempty" validate:"required" bson:"answer_option"`
	AnswerCorrectId []primitive.ObjectID `json:"answer_correct_id,omitempty" validate:"required" bson:"answer_correct_id"`
}

type ListQuestionResponse struct {
	Items []Question `json:"items,omitempty"`
	Total int64      `json:"total,omitempty"`
}
