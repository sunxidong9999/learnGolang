package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint32 `gorm: "primaryKey"`
	Name  string
	Email string
	Age   uint8
}
