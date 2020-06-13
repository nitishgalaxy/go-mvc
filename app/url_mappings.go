package app

import (
	"github.com/nitishgalaxy/go-mvc/controllers"
)

func mapUrls() {
	router.GET("/users/:userid", controllers.GetUser)
}
