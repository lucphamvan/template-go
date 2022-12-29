package quiz

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"tchh.lucpham/pkg/db"
	"tchh.lucpham/pkg/model"
)

type ServiceInterface interface {
	CreateQuestions(questions []model.Question) error
}

type Service struct {
}

func (s *Service) createQuestions(questions []model.CreateQuestionInput, questionIds *[]string) error {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)

	input := []interface{}{}
	for i := 0; i < len(questions); i++ {
		input = append(input, questions[i])
	}

	result, err := collection.InsertMany(context.Background(), input)
	if err != nil {
		return err
	}

	for i := 0; i < len(result.InsertedIDs); i++ {
		id := result.InsertedIDs[i].(primitive.ObjectID).Hex()
		*questionIds = append(*questionIds, id)
	}
	return nil
}

func (s *Service) CreateQuiz(inputQuiz model.CreateQuizInput) error {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// create list questions
	var questionIds []string
	err := s.createQuestions(inputQuiz.Question, &questionIds)
	if err != nil {
		return err
	}

	// create quiz
	quiz := model.Quiz{
		Name:       inputQuiz.Name,
		OwnerId:    inputQuiz.OwnerId,
		QuestionId: questionIds,
	}
	_, err = collection.InsertOne(context.Background(), quiz)
	if err != nil {
		return err
	}

	return nil
}

var ServiceInstance = new(Service)
