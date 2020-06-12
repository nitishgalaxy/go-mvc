package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/nitishgalaxy/go-mvc/services"
	"github.com/nitishgalaxy/go-mvc/utils"
)

// GetUser - Gets User
func GetUser(res http.ResponseWriter, req *http.Request) {
	// http://localhost:8080/users?user_id=4
	// http://localhost:8080/users?user_id=123
	userIDParam := req.URL.Query().Get("user_id")
	log.Printf("Processing user_id=%v", userIDParam)
	userID, err := strconv.ParseInt(userIDParam, 10, 64)

	if err != nil {
		// return bad request to client
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("User_id must be a number"),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)

		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		//Handle the err and return to client
		return
	}

	user, apiErr := services.GetUser(userID)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)

		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		//Handle the err and return to client
		return
	}

	// Return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
