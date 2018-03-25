package controllers

import (
  "github.com/kataras/iris"
  "../services"
)

func NewUser(ctx iris.Context) {
  name := ctx.FormValue("Name")
  if name != "" {
    user, err := services.CreateUser(name)
    if err != nil {
      ctx.JSON(iris.Map{"status": "error"})
    } else {
      ctx.JSON(iris.Map{"status": "success", "id": user.ID})
    }
  } else {
    ctx.JSON(iris.Map{"status": "error", "message": "name of user cannot be empty!"})
  }
}
