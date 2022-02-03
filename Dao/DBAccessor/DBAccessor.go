package DBAccessor

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func MySqlInit() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/Project1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused")
	}
	return db, err
}
