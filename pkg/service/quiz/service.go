package quiz

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tchh.lucpham/pkg/db"
	"tchh.lucpham/pkg/model"
)

type ServiceInterface interface {
}

type Service struct {
}

// implement function create question return question id
func (s *Service) createQuestion(inputQuestion model.CreateQuestionInput) (*model.Question, error) {
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
		return nil, err
	}
	// return question id
	questionId := result.InsertedID.(primitive.ObjectID)
	question.Id = questionId
	return &question, nil
}

func (s *Service) CreateQuiz(inputQuiz model.CreateQuizInput) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// create quiz
	quiz := model.Quiz{
		OwnerId:   inputQuiz.OwnerId,
		Setting:   inputQuiz.Setting,
		Deleted:   false,
		CreatedAt: float64(time.Now().UnixMilli()),
	}
	result, err := collection.InsertOne(context.Background(), quiz)
	if err != nil {
		return nil, err
	}

	quiz.Id = result.InsertedID.(primitive.ObjectID)
	return &quiz, nil
}

func (s *Service) CreateAndInsertQuestionToQuiz(quizId string, createQuestionInput model.CreateQuestionInput) (*model.CreateAndInsertQuestionToQuizResponse, error) {
	// create question
	question, err := s.createQuestion(createQuestionInput)
	if err != nil {
		return nil, err
	}

	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// update quiz
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": createQuestionInput.OwnerId}

	// write condition $ifNull
	condition := bson.M{"$ifNull": bson.A{
		bson.M{"$concatArrays": bson.A{"$question_ids", bson.A{question.Id.Hex()}}},
		bson.A{question.Id.Hex()},
	}}
	// aggressive update, need wrap it in array
	update := bson.A{
		bson.M{"$set": bson.M{"question_ids": condition}},
	}
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

	data := model.CreateAndInsertQuestionToQuizResponse{
		Quiz:     *quiz,
		Question: *question,
	}
	// return quiz
	return &data, nil
}

// remove question from quiz
func (s *Service) RemoveQuestionFromQuiz(quizId string, questionId string, ownerId string) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)

	// update quiz
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": ownerId}
	update := bson.M{"$pull": bson.M{"question_ids": questionId}}

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

// get list quiz
func (s *Service) GetQuizzes(ownerId string, limit int64, offset int64) (*model.GetListQuizzesResponse, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)
	// filter
	filter := bson.M{"owner_id": ownerId, "deleted": false}
	// options
	options := options.Find().SetSkip(limit * offset).SetLimit(limit)
	// find quiz
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	// decode quiz
	var quizzes []model.Quiz
	err = cursor.All(context.Background(), &quizzes)
	if err != nil {
		return nil, err
	}
	// count quiz
	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, err
	}

	// data
	data := model.GetListQuizzesResponse{
		Total: count,
		Items: quizzes,
	}

	// return data
	return &data, nil
}

var ServiceInstance = new(Service)
