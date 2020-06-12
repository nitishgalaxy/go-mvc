package services

import (
	"github.com/nitishgalaxy/go-mvc/model"
	"github.com/nitishgalaxy/go-mvc/utils"
)

func GetUser(userId int64) (*model.User, *utils.ApplicationError) {
	return model.GetUser(userId)
}
