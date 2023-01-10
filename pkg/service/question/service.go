package question

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tchh.lucpham/pkg/db"
	"tchh.lucpham/pkg/model"
)

type IService interface {
	GetList(limit, offset int64) (*model.ListQuestionResponse, error)
	Create(createQuestionInput model.CreateQuestionInput) (*model.Question, error)
}

type Service struct {
}

func (s *Service) GetList(limit, offset int64) (*model.ListQuestionResponse, error) {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)

	// find
	options := options.Find().SetLimit(limit).SetSkip(limit * offset)
	cursor, err := collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}

	// questions
	var questions []model.Question
	err = cursor.All(context.Background(), &questions)
	if err != nil {
		return nil, err
	}

	// get count
	count, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	data := model.ListQuestionResponse{
		Items: questions,
		Total: count,
	}

	return &data, nil
}

// create question
func (s *Service) Create(createQuestionInput model.CreateQuestionInput) (*model.Question, error) {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)
	result, err := collection.InsertOne(context.Background(), createQuestionInput)
	if err != nil {
		return nil, err
	}
	id, _ := result.InsertedID.(primitive.ObjectID)
	question := model.Question{
		Id:              id,
		Content:         createQuestionInput.Content,
		OwnerId:         createQuestionInput.OwnerId,
		AnswerOption:    createQuestionInput.AnswerOption,
		AnswerCorrectId: createQuestionInput.AnswerCorrectId,
	}
	return &question, nil
}

var ServiceInstance = new(Service)
