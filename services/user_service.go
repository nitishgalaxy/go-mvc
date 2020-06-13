package services

import (
	"github.com/nitishgalaxy/go-mvc/model"
	"github.com/nitishgalaxy/go-mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (s *userService) GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	user, err := model.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
