package services

import (
  "../models"
  "../dbInterface"
  "github.com/kataras/iris"
)

func CreatePair(first int, second int) (models.Pair, error) {
  pair := models.Pair{}
  return pair, dbInterface.Psql.
    FirstOrCreate(
    &pair,
    models.Pair{UserOne: uint(first), UserTwo: uint(second)},
  ).Error
}

func GetAllPairCounts() iris.Map {
  var users []models.User
  var pairs []models.Pair

  dbInterface.Psql.Find(&users).Order("id asc")
  dbInterface.Psql.Find(&pairs)
  return iris.Map{"users": users, "pairs": pairs}
}
