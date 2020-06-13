package services

import (
	"net/http"
	"testing"

	"github.com/nitishgalaxy/go-mvc/model"
	"github.com/nitishgalaxy/go-mvc/utils"
	"github.com/stretchr/testify/assert"
)

type usersDaoMock struct{}

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*model.User, *utils.ApplicationError)
)

func init() {
	model.UserDao = &usersDaoMock{}
}

func (u *usersDaoMock) GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int64) (*model.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "User 0 not found",
		}
	}

	user, err := UserService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User 0 not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {

	getUserFunction = func(userId int64) (*model.User, *utils.ApplicationError) {
		return &model.User{
			ID: 123,
		}, nil
	}

	user, err := UserService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
}
