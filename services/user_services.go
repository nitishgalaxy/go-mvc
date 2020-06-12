package services

import "github.com/nitishgalaxy/go-mvc/model"

func GetUser(userId int64) (*model.User, error) {
	return model.GetUser(userId)
}
