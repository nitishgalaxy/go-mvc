package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/nitishgalaxy/go-mvc/services"
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
		return
	}

	user, err := services.GetUser(userID)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("user_id must be a number"))
		//Handle the err and return to client
		return
	}

	// Return user to client
	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
