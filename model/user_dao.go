package model

import (
	"fmt"
	"net/http"

	"github.com/nitishgalaxy/go-mvc/utils"
)

/*
	1. We create a struct userDao
	2. We attach method GetUser to struct.
	3. struct userDao also implements userDaoInterface
	4. UserDao is a variable of userDaoInterface

*/

type userDao struct{} // Using interface instead of struct allows us to mock

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Tom", LastName: "Cruise", Email: "tomcruise@gmail.com"},
	}

	//UserDao userDao
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
	// UserDao is a pointer to blank struct userDao
}

// *userDao = struct userDao = also userDaoInterface
func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	// *User allow us to return nil in case of user==nil

	fmt.Println("Accessing database.")
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
