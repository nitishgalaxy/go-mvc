package model

import (
	"errors"
	"fmt"
)

var (
	users = map[int64]*User{
		123: &User{ID: 123, FirstName: "Tom", LastName: "Cruise", Email: "tomcruise@gmail.com"},
	}
)

func GetUser(userId int64) (*User, error) {
	// *User allow us to return nil in case of user==nil
	user := users[userId]
	if user == nil {
		return nil, errors.New(fmt.Sprintf("user %v was not found", userId))
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
