package models

import (
	"github.com/jinzhu/gorm"
  "../dbInterface"
)

type Pair struct {
	gorm.Model

	UserOne uint `gorm:foreignkey`
	UserTwo uint `gorm:foreignkey`

	Count uint `gorm:"default:0"`
}

func (p *Pair) IncrementCount() error {
	p.Count += 1
	return p.Save()
}

func (p *Pair) Save() error {
	return dbInterface.Psql.Save(&p).Error
}
