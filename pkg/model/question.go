package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Question struct
type Question struct {
	Id               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Content          string             `json:"content,omitempty" validate:"required"`
	OwnerId          string             `json:"owner_id,omitempty" bson:"owner_id"`
	Choices          []Choice           `json:"choices,omitempty" bson:"choices" validate:"required,min=1"`
	CorrectChoiceIds []string           `json:"correct_choice_ids,omitempty" bson:"correct_choice_ids" validate:"required,min=1"`
	Deleted          bool               `json:"deleted,omitempty"`

	CreatedAt float64 `json:"created_at,omitempty"`
}

// Choice struct
type Choice struct {
	Id      primitive.ObjectID `json:"id,omitempty" validate:"required"`
	Content string             `json:"content,omitempty" validate:"required"`
}

// QuestionAnswer struct
type QuestionAnswer struct {
	QuestionId       string   `json:"question_id,omitempty" bson:"question_id"`
	CorrectChoiceIds []string `json:"correct_choice_ids,omitempty" bson:"correct_choice_ids"`
}

// struct for Request and Response
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
