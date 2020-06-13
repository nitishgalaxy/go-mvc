package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/nitishgalaxy/go-mvc/services"
	"github.com/nitishgalaxy/go-mvc/utils"
)

// GetUser - Gets User
func GetUser(c *gin.Context) {
	// http://localhost:8080/users?user_id=4
	// http://localhost:8080/users?user_id=123

	userID, err := strconv.ParseInt(c.Param("userid"), 10, 64)

	if err != nil {
		// return bad request to client
		apiErr := &utils.ApplicationError{
			Message:    fmt.Sprintf("User_id must be a number"),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		//c.JSON(apiErr.StatusCode, apiErr)
		//utils.Respond(c, apiErr.StatusCode, apiErr)
		utils.RespondError(c, apiErr)

		return // If we dont put return, code keeps executing.
		// Cool feature... if you want to write metrics, logs etc.
	}

	user, apiErr := services.UserService.GetUser(userID)

	if apiErr != nil {

		//c.JSON(apiErr.StatusCode, apiErr)
		//utils.Respond(c, apiErr.StatusCode, apiErr)
		utils.RespondError(c, apiErr)

		return
	}

	// Return user to client
	//c.JSON(http.StatusOK, user)
	utils.Respond(c, http.StatusOK, user)
}
