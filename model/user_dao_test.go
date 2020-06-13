package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	// Initialization

	// Execution
	user, err := GetUser(0)

	// Validation

	// user should be nil
	assert.Nil(t, user, "we were not expecting a user with id 0")

	// err should be not nil
	assert.NotNil(t, err)

	// err.StatusCode should be equal  http.StatusNotFound
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)

	assert.EqualValues(t, "not_found", err.Code)

	assert.EqualValues(t, "User 0 not found", err.Message)

	/*
		// Using Plain Golang
		if user != nil {
			t.Error("we were not expecting a user with id 0")
		}

		if err == nil {
			t.Error("we were expecting an error when user id is 0")
		}

		if err.StatusCode != http.StatusNotFound {
			t.Error("we ere expecting 404 when user is not found")
		}
	*/
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Tom", user.FirstName)
	assert.EqualValues(t, "Cruise", user.LastName)
}
