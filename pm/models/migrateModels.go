package models

import (
  "github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) {
  db.AutoMigrate(&User{})
  db.AutoMigrate(&Pair{})
  db.Model(Pair{}).AddForeignKey("user_one", "users (id)", "CASCADE", "CASCADE")
  db.Model( Pair{}).AddForeignKey("user_two", "users (id)", "CASCADE", "CASCADE")
}

