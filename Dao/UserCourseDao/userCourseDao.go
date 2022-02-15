package UserCourseDao

import (
	"Project/Dao/DBAccessor"
	"Project/Types"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type UserCourseDao struct {
	CourseIDs []string `gorm:"type:varchar(128)"`
	UserID    string   `gorm:"type:varchar(128)"`
	gorm.Model
}

// TableName 改表名为“userCourses”
func (UserCourseDao) TableName() string {
	return "userCourses"
}

// makeUserCourseTable 如果表不存在就建表,并返回最终是否有该表
func makeUserCourseTable() bool {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		if !db.HasTable(&UserCourseDao{}) {
			db.AutoMigrate(&UserCourseDao{})
		} else {
			return true
		}
	}
	if db.HasTable(&UserCourseDao{}) {
		return true
	} else {
		return false
	}
}
func InsertUserCourse(courseIDs []string, userID string) {
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
			if makeUserCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'userCourses'.Please check the database.")
			}
		}
		db.Create(&UserCourseDao{UserID: userID, CourseIDs: courseIDs})
	}
}
func FindUserCoursesByUserID(userID string) []string {
	var res UserCourseDao
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
			if makeUserCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'userCourses'.Please check the database.")
			}
		}
		db.Where(&UserCourseDao{UserID: userID}).Find(&res)
	}
	return res.CourseIDs
}
func UpdateUserCoursesByCourseID(courseIDs []string, userID string) Types.ErrNo {
	var res UserCourseDao
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
			if makeUserCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'userCourses'.Please check the database.")
			}
		}
		db.Where(&UserCourseDao{UserID: userID}).Find(&res)
		res.CourseIDs = courseIDs
		db.Save(res)
		return Types.OK
	}
}
func DeleteUserCoursesByCourseID(userID string) Types.ErrNo {
	var res UserCourseDao
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
			if makeUserCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'userCourses'.Please check the database.")
			}
		}
		db.Where(&UserCourseDao{UserID: userID}).Find(&res)
		db.Delete(res)
		return Types.OK
	}
}
