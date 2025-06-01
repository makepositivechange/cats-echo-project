package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type CatInfo struct {
	gorm.Model
	BreedName   string                      `gorm:"type:varchar(200); not null"`
	BreedOrigin string                      `gorm:"type:varchar(200); not null"`
	BreedType   string                      `gorm:"type:varchar(200); not null"`
	TypeInfo    *string                     `gorm:"type:varchar(200)"`
	BodyTypes   datatypes.JSONSlice[string] `gorm:"type:json"`
	CoatPattern string                      `gorm:"type:varchar(200);not null"`
}

func (ct *CatInfo) TableName() string {
	return "catinfo"
}
