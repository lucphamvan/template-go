package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Email     string             `json:"email,omitempty" validate:"email,required"`
	Name      string             `json:"name,omitempty"`
	Avatar    string             `json:"avatar,omitempty"`
	Password  string             `json:"-"`
	CreatedAt float64            `json:"created_at,omitempty"`
	UpdatedAt float64            `json:"updated_at,omitempty"`
}

type CreateUserInput struct {
	Email    string `json:"email,omitempty" validate:"email,required"`
	Name     string `json:"name,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Password string `json:"password,omitempty"`
}

type UpdateUserInput struct {
	Email  string `json:"email,omitempty" validate:"email,required"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type ListUserResponse struct {
	Items []User `json:"items"`
	Total int64  `json:"total"`
}
