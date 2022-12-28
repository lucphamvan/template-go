package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/db"
	"tchh.lucpham/pkg/model"
)

type ServiceInterface interface {
	Create(u model.CreateUserInput) (*model.User, error)
	Get(id string) (*model.User, error)
	GetList(limit int64, offset int64) (*[]model.User, error)
	Update(id string, u model.UpdateUserInput) (*model.User, error)
	Delete(id string) error
	IsUserExist(email string) bool
	Login(email, password string) (*model.AuthenResponse, error)
	Access() (*model.User, error)
}

type Service struct {
}

/* Create user */
func (s *Service) Create(input model.CreateUserInput) (*model.User, error) {
	// check user existed or not
	isExist := s.IsUserExist(input.Email)
	if isExist {
		return nil, errors.New(common.ERROR_USER_EXISTED)
	}

	collection := db.Client.Database(db.DATABASE).Collection(db.USER_COLLECTION)
	// generate password
	hashpass, err := common.GeneratePassword(input.Password)
	if err != nil {
		return nil, err
	}
	user := model.User{
		Email:     input.Email,
		Name:      input.Name,
		Password:  hashpass,
		CreatedAt: float64(time.Now().UnixMilli()),
		UpdatedAt: float64(time.Now().UnixMilli()),
	}

	// insert into database
	result, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	user.Id = result.InsertedID.(primitive.ObjectID)
	user.Password = ""

	return &user, nil
}

// Get list user with pagination
func (s *Service) GetList(limit int64, offset int64) (*[]model.User, error) {
	collection := db.Client.Database(db.DATABASE).Collection(db.USER_COLLECTION)

	// find option with skip, limit
	var findOption *options.FindOptions
	if limit != 0 {
		findOption = options.Find().SetSkip(limit * offset).SetLimit(limit)
	}

	cursor, err := collection.Find(context.Background(), bson.M{}, findOption)
	if err != nil {
		return nil, err
	}
	// mapping value
	var users []model.User
	err = cursor.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	// return value
	return &users, nil
}

// get user by id
func (s *Service) Get(id string) (*model.User, error) {
	collection := db.Client.Database(db.DATABASE).Collection(db.USER_COLLECTION)

	fmt.Println("id", id)
	objectId, _ := primitive.ObjectIDFromHex(id)
	result := collection.FindOne(context.Background(), bson.M{"_id": objectId})

	user := new(model.User)
	err := result.Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Update(id string, u model.UpdateUserInput) (*model.User, error) {
	return nil, errors.New("not implement")
}

func (s *Service) Delete(id string) error {
	return nil
}

func (s *Service) IsUserExist(email string) bool {
	collection := db.Client.Database(db.DATABASE).Collection(db.USER_COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"email": email})
	return result.Err() == nil
}

func (s *Service) Login(email, password string) (*model.AuthenResponse, error) {
	// find user in db
	collection := db.Client.Database(db.DATABASE).Collection(db.USER_COLLECTION)
	result := collection.FindOne(context.Background(), bson.M{"email": email})
	var user model.User
	err := result.Decode(&user)
	if err != nil {
		return nil, err
	}
	// verify password
	isValid := common.VerifyPassword(password, user.Password)
	if !isValid {
		return nil, errors.New("password is incorrect")
	}
	// generate token
	accToken, err := common.GenerateAccToken(user.Id.Hex())
	if err != nil {
		return nil, err
	}
	refToken, err := common.GenerateRefToken(user.Id.Hex())
	if err != nil {
		return nil, err
	}
	// response
	response := &model.AuthenResponse{
		AccessToken:  accToken,
		RefreshToken: refToken,
		User:         user,
	}
	return response, nil
}

func (s *Service) Access() (*model.User, error) {
	return nil, nil
}

var ServiceInstance = new(Service)
