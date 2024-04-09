package main

import (
	"fmt"
	"os"

	"gormtest/db"
	"gormtest/model"
)

func main() {
	dbconn, err := db.InitDBSqlite("test.db")
	if err != nil {
		fmt.Println("Failed to connect to sqlite db:", err)
		os.Exit(-1)
	}

	/*
		err = dbconn.AutoMigrate(&model.Product{}, &model.User{})
		if err != nil {
			panic("Create table failed")
		}
			err = db.CreateTableProduct(dbconn, &model.User{})
			if err != nil {
				panic("Product create table failed")
			}
	*/

	err = db.CreateTable(dbconn,
		&model.Product{},
		&model.User{})
	if err != nil {
		panic("Product create table failed")
	}

	p1 := model.Product{
		ID:    1,
		Name:  "product1",
		Desc:  "A good product",
		Price: 12.5,
	}
	p2 := &model.Product{
		ID:    2,
		Name:  "product2",
		Desc:  "A good product 2",
		Price: 11.5,
	}
	p3 := &model.Product{
		ID:    3,
		Name:  "product3",
		Desc:  "A good product 3",
		Price: 22.5,
	}
	err = db.AddProductItem(dbconn, &p1)
	if err != nil {
		panic("Product 1 insert failed")
	}

	err = db.AddProductItem(dbconn, &p2)
	if err != nil {
		panic("Product 2 insert failed")
	}

	err = db.AddProductItem(dbconn, &p3)
	if err != nil {
		panic("Product 3 insert failed")
	}
	fmt.Println("Find product 2:")
	p, err := db.LookupProductItemByPrice(dbconn, 11.5)
	if err != nil {
		panic("find p2 failed")
	}
	fmt.Println(p)

	fmt.Println("Update product1 price:")
	db.UpdateProductItemPriceByName(dbconn, "product1", 15.4)
	fmt.Println("Find product 1:")
	p, err = db.LookupProductItemByPrice(dbconn, 15.4)
	if err != nil {
		panic("find p1 failed")
	}
	fmt.Println(p)

	fmt.Println("find products with price > 10.1:")

	var pp []model.Product
	tx := dbconn.Find(&pp, "price > ?", 10.1)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		panic("find price > 10 failed")
	}
	fmt.Println(pp)

	/*
		tx := dbconn.Create(&p1)
		if tx.Error != nil {
			panic("Product 1 insert failed")
		}
		tx = dbconn.Create(&p2)
		if tx.Error != nil {
			panic("Product 2 insert failed")
		}
		tx = dbconn.Create(&p3)
		if tx.Error != nil {
			panic("Product 3 insert failed")
		}

		fmt.Println("Find product 2:")
		var p model.Product
		tx = dbconn.First(&p, "price = ?", 11.5)
		if tx.Error != nil {
			panic("find p2 failed")
		}
		fmt.Println(p)
	*/
}
