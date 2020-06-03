package app

import (
	"net/http"

	"github.com/nitishgalaxy/go-mvc/controllers"
)

// StartApp - Start application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
