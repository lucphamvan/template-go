package quiz

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"tchh.lucpham/pkg/db"
	"tchh.lucpham/pkg/model"
)

type ServiceInterface interface {
}

type Service struct {
}

func (s *Service) createQuestions(questions []model.CreateQuestionInput) (*[]string, error) {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)

	input := []interface{}{}
	for i := 0; i < len(questions); i++ {
		input = append(input, questions[i])
	}

	result, err := collection.InsertMany(context.Background(), input)
	if err != nil {
		return nil, err
	}

	var questionIds []string
	for i := 0; i < len(result.InsertedIDs); i++ {
		id := result.InsertedIDs[i].(primitive.ObjectID).Hex()
		questionIds = append(questionIds, id)
	}

	return &questionIds, nil
}

func (s *Service) CreateQuizzes(inputQuiz model.CreateQuizzesInput) (*model.Quizzes, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// create list questions
	questionIds, err := s.createQuestions(inputQuiz.Questions)
	if err != nil {
		return nil, err
	}

	// create quiz
	quiz := model.Quizzes{
		Name:       inputQuiz.Name,
		OwnerId:    inputQuiz.OwnerId,
		QuestionId: *questionIds,
	}
	result, err := collection.InsertOne(context.Background(), quiz)
	if err != nil {
		return nil, err
	}

	quiz.Id = result.InsertedID.(primitive.ObjectID)
	return &quiz, nil
}

var ServiceInstance = new(Service)
