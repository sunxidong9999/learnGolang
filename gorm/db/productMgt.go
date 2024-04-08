package db

import (
	"gormtest/model"

	"gorm.io/gorm"
)

func AddProductItem(db *gorm.DB, item interface{}) error {
	tx := db.Create(item)
	return tx.Error
}

func LookupProductItemByPrice(db *gorm.DB, price float32) (model.Product, error) {
	var p model.Product
	tx := db.First(&p, "price = ?", price)
	return p, tx.Error
}

func UpdateProductItemPriceByName(db *gorm.DB, name string, newPrice float32) error {
	var p model.Product
	tx := db.First(&p, "name = ?", name)
	if tx.Error != nil {
		return tx.Error
	}
	tx = db.Model(&p).Update("Price", newPrice)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
