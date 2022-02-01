package UserDao

import (
	"Project/Dao/DBAccessor"
	"Project/Types"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type UserDao struct {
	UserID   string `gorm:"type:varchar(128)"`
	Password string `gorm:"type:varchar(256)"`
	gorm.Model
}

// TableName 改表名为“courses”
func (UserDao) TableName() string {
	return "users"
}

// makeUserTable 如果表不存在就建表,并返回最终是否有该表
func makeUserTable() bool {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		if !db.HasTable(&UserDao{}) {
			db.AutoMigrate(&UserDao{})
		}
	}
	if db.HasTable(&UserDao{}) {
		return true
	} else {
		return false
	}
}

// InsertUser 插入一条数据
func InsertUser(userid string, password string) Types.ErrNo {
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
			if makeUserTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var user UserDao
		user.UserID = userid
		user.Password = password
		db.Create(&user)
		return Types.OK
	}
}

// CheckUser user和password是否正常匹配
func CheckUser(userid string, password string) (bool, Types.ErrNo) {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return false, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeUserTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var user UserDao
		db.Where(&UserDao{UserID: userid}).Find(&user)
		if user.UserID == "" {
			return false, Types.UserNotExisted
		}
		if user.Password == password {
			return true, Types.OK
		} else {
			return false, Types.OK
		}
	}
}

// UpdatePassword 根据用户id改密码
func UpdatePassword(userid string, password string) Types.ErrNo {
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
			if makeUserTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var user UserDao
		db.Where(&UserDao{UserID: userid}).Find(&user)
		if user.UserID == "" {
			return Types.UserNotExisted
		}
		db.Model(&user).Update("password", password)
		return Types.OK
	}
}
