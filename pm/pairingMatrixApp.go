package pm

import (
  "github.com/kataras/iris"
  "github.com/iris-contrib/middleware/cors"
)

var app *iris.Application = nil

func setupApp() *iris.Application {
  var crs = cors.New(cors.Options{
    AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
    AllowCredentials: true,
  })
  var App = iris.Default()
  App.Use(crs)
  return App
}

func App() *iris.Application {
  if app == nil {
    app = setupApp()
  }
  return app
}
