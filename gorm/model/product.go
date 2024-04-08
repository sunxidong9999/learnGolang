package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID    uint32 `gorm: "primaryKey"`
	Name  string
	Desc  string
	Price float32
}
