package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"tchh.lucpham/pkg/common"
	"tchh.lucpham/pkg/model"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Create(u model.CreateUserInput) (*model.User, error) {
	args := m.Called(u)
	user, ok := args.Get(0).(*model.User)
	if ok {
		return user, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockService) Get(id string) (*model.User, error) {
	return nil, nil
}

func (m *MockService) GetList(limit int64, offset int64) (*[]model.User, error) {
	return nil, nil
}

func (m *MockService) Update(id string, u model.UpdateUserInput) (*model.User, error) {
	return nil, nil
}

func (m *MockService) Delete(id string) error {
	return nil
}

func (m *MockService) IsUserExist(email string) bool {
	return false
}

func (m *MockService) Login(email, password string) (*model.AuthenResponse, error) {
	return nil, nil
}

func (m *MockService) Access() (*model.User, error) {
	return nil, nil
}

func TestHandler_Create(t *testing.T) {
	t.Run("Create user failed to parse request body", func(t *testing.T) {
		// mock gin context
		w := httptest.NewRecorder()
		context := common.MockGinContext(w)

		// call api
		handler := NewHandler(ServiceInstance)
		handler.Create(context)
		// expected and output
		expected := model.Error{
			Error: common.ERROR_BIND_JSON,
		}
		var output model.Error
		json.Unmarshal(w.Body.Bytes(), &output)
		// assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.EqualValues(t, expected, output)
	})

	t.Run("Create user failed with invalid request body", func(t *testing.T) {
		// mock context & request body
		w := httptest.NewRecorder()
		context := common.MockGinContext(w)
		data := model.CreateUserInput{
			Name:     "Aka",
			Password: "AkaPassword123!",
		}
		common.MockRequestBody(context, data)
		// call api
		handler := NewHandler(ServiceInstance)
		handler.Create(context)

		// expected
		err := validator.New().Struct(data)
		expected := model.Error{
			Error: err.Error(),
		}
		var output model.Error
		json.Unmarshal(w.Body.Bytes(), &output)
		// assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.EqualValues(t, expected, output)
	})

	t.Run("Create user failed because service create failed", func(t *testing.T) {
		// mock context & request body
		w := httptest.NewRecorder()
		context := common.MockGinContext(w)
		data := model.CreateUserInput{
			Email:    "Aka@gmail.com",
			Name:     "Aka",
			Password: "AkaPassword123!",
		}
		common.MockRequestBody(context, data)

		// mock service
		err := errors.New("Error create user")
		mockService := new(MockService)
		mockService.On("Create", data).Return(nil, err)

		// call api
		handler := NewHandler(mockService)
		handler.Create(context)

		// expected
		expected := model.Error{
			Error: err.Error(),
		}
		var output model.Error
		json.Unmarshal(w.Body.Bytes(), &output)
		// assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.EqualValues(t, expected, output)
	})

	t.Run("Create user success with valid body data", func(t *testing.T) {
		// mock context + request body
		w := httptest.NewRecorder()
		context := common.MockGinContext(w)
		data := model.CreateUserInput{
			Email:    "Aka@gmail.com",
			Name:     "Aka",
			Password: "AkaPassword123!",
		}
		common.MockRequestBody(context, data)

		// mock service
		user := model.User{
			Id:    primitive.NewObjectID(),
			Email: data.Email,
			Name:  data.Name,
		}
		mockService := new(MockService)
		mockService.On("Create", data).Return(&user, nil)
		handler := NewHandler(mockService)
		handler.Create(context)

		// output
		var outputUser model.User
		json.Unmarshal(w.Body.Bytes(), &outputUser)
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.EqualValues(t, user, outputUser)
	})
}
