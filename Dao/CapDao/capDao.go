package CapDao

import (
	"Project/Dao/DBAccessor"
	"Project/Types"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type CapDao struct {
	CourseID string `gorm:"type:varchar(128)"`
	Cap      int    `gorm:"type:varchar(128)"`
	gorm.Model
}

// TableName 改表名为“caps”
func (CapDao) TableName() string {
	return "caps"
}

// makeCapTable 如果表不存在就建表,并返回最终是否有该表
func makeCapTable() bool {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		if !db.HasTable(&CapDao{}) {
			db.AutoMigrate(&CapDao{})
		} else {
			return true
		}
	}
	if db.HasTable(&CapDao{}) {
		return true
	} else {
		return false
	}
}

func InsertCap(courseID string, cap int) {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		// 直到建表成功才继续
		for true {
			if makeCapTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'caps'.Please check the database.")
			}
		}
		db.Create(&CapDao{Cap: cap, CourseID: courseID})
	}
}

func FindCapByCourseID(courseID string) int {
	var res CapDao
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		// 直到建表成功才继续
		for true {
			if makeCapTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'caps'.Please check the database.")
			}
		}
		db.Where(&CapDao{CourseID: courseID}).Find(&res)
	}
	return res.Cap
}

func UpdateCapByCourseID(courseID string, cap int) Types.ErrNo {
	var res CapDao
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeCapTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'caps'.Please check the database.")
			}
		}
		db.Where(&CapDao{CourseID: courseID}).Find(&res)
		res.Cap = cap
		db.Save(res)
		return Types.OK
	}
}

func DeleteCapByCourseID(courseID string) Types.ErrNo {
	var res CapDao
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeCapTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'caps'.Please check the database.")
			}
		}
		db.Where(&CapDao{CourseID: courseID}).Find(&res)
		db.Delete(res)
		return Types.OK
	}
}
