package model

import (
	"fmt"
	"net/http"

	"github.com/nitishgalaxy/go-mvc/utils"
)

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Tom", LastName: "Cruise", Email: "tomcruise@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	// *User allow us to return nil in case of user==nil
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User %v not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}

/*
Alternative approach:
func GetUser(userId int64) (*User, error) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, errors.New(fmt.Fprintf("user %v was not found", userId))

}

*/
