package app

import (
	"fmt"
	"github.com/fausto4911/GO-USER-API/controllers"
)

const (
	USERS_PATH = "/users"
)

func mapUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET(fmt.Sprintf("%s%s", USERS_PATH, "/:user_id"), controllers.GetUser)
	router.POST(USERS_PATH, controllers.CreateUser)
}
