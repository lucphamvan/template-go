package quiz

import (
	"tchh.lucpham/pkg/model"
)

type ServiceInterface interface {
}

type Service struct {
}

func (s *Service) createQuestions(questions []model.CreateQuestionInput) (*[]string, error) {
	return nil, nil
	// collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)

	// input := []interface{}{}
	// for i := 0; i < len(questions); i++ {
	// 	input = append(input, questions[i])
	// }

	// result, err := collection.InsertMany(context.Background(), input)
	// if err != nil {
	// 	return nil, err
	// }

	// var questionIds []string
	// for i := 0; i < len(result.InsertedIDs); i++ {
	// 	id := result.InsertedIDs[i].(primitive.ObjectID).Hex()
	// 	questionIds = append(questionIds, id)
	// }

	// return &questionIds, nil
}

func (s *Service) CreateQuizzes(inputQuiz model.CreateQuizzesInput) (*model.Quizzes, error) {
	return nil, nil
	// collection
	// collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// // create list questions
	// questionIds, err := s.createQuestions(inputQuiz.Questions)
	// if err != nil {
	// 	return nil, err
	// }

	// // create quiz
	// quiz := model.Quizzes{
	// 	// Name:       inputQuiz.Name,
	// 	OwnerId:    inputQuiz.OwnerId,
	// 	QuestionId: *questionIds,
	// }
	// result, err := collection.InsertOne(context.Background(), quiz)
	// if err != nil {
	// 	return nil, err
	// }

	// quiz.Id = result.InsertedID.(primitive.ObjectID)
	// return &quiz, nil
}

var ServiceInstance = new(Service)
