package app

import (
	"fmt"
	"net/http"

	"github.com/nitishgalaxy/go-mvc/controllers"
)

// StartApp - Start application
func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	fmt.Println("Server running at port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
