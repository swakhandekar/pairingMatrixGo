package routes

import (
	"../../pm"
	"../controllers"
)

func UserRegister() {
	pm.App().Post("/user/new", controllers.NewUser)
}
