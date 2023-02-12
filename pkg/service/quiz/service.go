package quiz

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tchh.lucpham/pkg/db"
	"tchh.lucpham/pkg/model"
)

type ServiceInterface interface {
}

type Service struct {
}

// implement function create question return question id
func (s *Service) createQuestion(inputQuestion model.CreateQuestionInput) (string, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)

	// create question
	question := model.Question{
		OwnerId:          inputQuestion.OwnerId,
		Content:          inputQuestion.Content,
		Choices:          inputQuestion.Choices,
		Deleted:          false,
		CorrectChoiceIds: inputQuestion.CorrectChoiceIds,
		CreatedAt:        float64(time.Now().UnixMilli()),
	}
	result, err := collection.InsertOne(context.Background(), question)
	if err != nil {
		return "", err
	}
	// return question id
	questionId := result.InsertedID.(primitive.ObjectID).Hex()
	return questionId, nil
}

func (s *Service) CreateQuiz(inputQuiz model.CreateQuizInput) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// create quiz
	quiz := model.Quiz{
		OwnerId: inputQuiz.OwnerId,
		Setting: inputQuiz.Setting,
	}
	result, err := collection.InsertOne(context.Background(), quiz)
	if err != nil {
		return nil, err
	}

	quiz.Id = result.InsertedID.(primitive.ObjectID)
	return &quiz, nil
}

func (s *Service) CreateAndInsertQuestionToQuiz(quizId string, question model.CreateQuestionInput) (*model.Quiz, error) {
	// create question
	questionId, err := s.createQuestion(question)
	if err != nil {
		return nil, err
	}

	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// update quiz
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": question.OwnerId}
	update := bson.M{"$push": bson.M{"questions": questionId}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	// find quiz
	result := collection.FindOne(context.Background(), bson.M{"_id": objectId})
	quiz := new(model.Quiz)
	err = result.Decode(quiz)
	if err != nil {
		return nil, err
	}

	// return quiz
	return quiz, nil
}

// remove question from quiz
func (s *Service) RemoveQuestionFromQuiz(quizId string, questionId string, ownerId string) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// update quiz
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": ownerId}
	update := bson.M{"$pull": bson.M{"questions": questionId}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	// find quiz
	result := collection.FindOne(context.Background(), bson.M{"_id": objectId})
	quiz := new(model.Quiz)
	err = result.Decode(quiz)
	if err != nil {
		return nil, err
	}

	// return quiz
	return quiz, nil
}

var ServiceInstance = new(Service)
