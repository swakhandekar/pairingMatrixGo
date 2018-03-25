package services

import (
	"../models"
  "../dbInterface"
)

func CreateUser(name string) (models.User, error) {
	user := models.User{Name:name}
	return user, dbInterface.Psql.Create(&user).Error
}
