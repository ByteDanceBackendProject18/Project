package DBAccessor

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func MySqlInit() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/irisapp?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused")
	}
	return db, err
}
