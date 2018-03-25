package main

import (
	"github.com/kataras/iris"
	"./pm"
	"./pm/routes"
  "./pm/models"
  "./pm/dbInterface"
)


func main() {
  models.Migrate(dbInterface.Psql)

	routes.Register()
	pm.App().Run(iris.Addr(":8080"))
}
