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
	GetList(limit, offset int64, ownerId string) (*model.ListQuestionResponse, error)
	Create(createQuestionInput model.CreateQuestionInput) (*model.Question, error)
	Get(id string) (*model.Question, error)
	Delete(id string, userId string) error
	Update(id string, updateQuestionInput model.UpdateQuestionInput, userId string) error
}

type Service struct {
}

// get list question
func (s *Service) GetList(limit, offset int64, ownerId string) (*model.ListQuestionResponse, error) {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)

	// filter
	filter := bson.M{"deleted": false, "owner_id": ownerId}
	// find
	options := options.Find().SetLimit(limit).SetSkip(limit * offset)
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}

	// questions
	var questions = make([]model.Question, 0)
	err = cursor.All(context.Background(), &questions)
	if err != nil {
		return nil, err
	}

	// get count
	count, err := collection.CountDocuments(context.Background(), filter)
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
		Id:               id,
		Content:          createQuestionInput.Content,
		OwnerId:          createQuestionInput.OwnerId,
		Choices:          createQuestionInput.Choices,
		CorrectChoiceIds: createQuestionInput.CorrectChoiceIds,
		Deleted:          createQuestionInput.Deleted,
	}
	return &question, nil
}

func (s *Service) Get(id string) (*model.Question, error) {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)

	objectId, _ := primitive.ObjectIDFromHex(id)
	result := collection.FindOne(context.Background(), bson.M{"_id": objectId})

	question := new(model.Question)
	err := result.Decode(question)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *Service) Delete(id string, userId string) error {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)
	objectId, _ := primitive.ObjectIDFromHex(id)

	// delete : update deleted field to 'true'
	filter := bson.M{"_id": objectId, "owner_id": userId}
	_, err := collection.UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"deleted": true}})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(id string, updateQuestionInput model.UpdateQuestionInput, userId string) error {
	collection := db.Client.Database(db.DATABASE).Collection(db.QUESTION_COLLECTION)
	objectId, _ := primitive.ObjectIDFromHex(id)

	// update question
	filter := bson.M{"_id": objectId, "owner_id": userId}
	_, err := collection.UpdateOne(context.Background(), filter, bson.M{
		"$set": updateQuestionInput,
	})

	if err != nil {
		return err
	}

	return nil
}

var ServiceInstance = new(Service)
