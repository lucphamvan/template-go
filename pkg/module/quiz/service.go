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
	var quizzes = make([]model.Quiz, 0)
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

// get quiz
func (s *Service) GetQuiz(quizId string, ownerId string) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)
	// filter
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": ownerId, "deleted": false}
	// find quiz
	result := collection.FindOne(context.Background(), filter)
	quiz := new(model.Quiz)
	err := result.Decode(quiz)
	if err != nil {
		return nil, err
	}
	return quiz, nil
}

// get questions of quiz
func (s *Service) GetQuestions(quizId string, ownerId string) (*[]model.Question, error) {
	// get quiz
	quiz, err := s.GetQuiz(quizId, ownerId)
	if err != nil {
		return nil, err
	}

	listObjIds := make([]primitive.ObjectID, len(quiz.QuestionIds))
	for i, id := range quiz.QuestionIds {
		objId, _ := primitive.ObjectIDFromHex(id)
		listObjIds[i] = objId
	}
	// filter question have id in quiz.question_ids
	filter := bson.M{"_id": bson.M{"$in": listObjIds}}
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)
	// find questions
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	// decode questions
	var questions []model.Question
	err = cursor.All(context.Background(), &questions)
	if err != nil {
		return nil, err
	}
	return &questions, nil
}

// publish quiz
func (s *Service) PublishQuiz(quizId string, ownerId string) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)
	// update quiz
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": ownerId}
	update := bson.M{"$set": bson.M{"published": true}}
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
	return quiz, nil

}

// update quiz setting
func (s *Service) UpdateQuizSetting(quizId string, ownerId string, setting model.QuizSetting) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)
	// update quiz
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": ownerId}
	update := bson.M{"$set": bson.M{"setting": setting}}
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
	return quiz, nil
}

// delete quiz
func (s *Service) DeleteQuiz(quizId string, ownerId string) (*model.Quiz, error) {
	// collection
	collection := db.Client.Database(db.DATABASE).Collection(db.QUIZ_COLLECTION)
	// update quiz
	objectId, _ := primitive.ObjectIDFromHex(quizId)
	filter := bson.M{"_id": objectId, "owner_id": ownerId}
	update := bson.M{"$set": bson.M{"deleted": true}}
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
	return quiz, nil
}

var ServiceInstance = new(Service)
