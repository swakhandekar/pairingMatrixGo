package routes

import (
	"../controllers"
	"../../pm"
)

func PairingRegister() {
	pm.App().Post("/pairing/increment", controllers.IncrementPairingCount)
	pm.App().Get("/pairing/all", controllers.AllPairs)
}
