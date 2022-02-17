package UserCourseDao

import (
	"Project/Dao/DBAccessor"
	"Project/Types"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type UserCourseDao struct {
	UserID   string `gorm:"type:varchar(128)"`
	CourseID string `gorm:"type:varchar(128)"`
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

func InsertUserCourse(courseID string, userID string) {
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
		db.Create(&UserCourseDao{UserID: userID, CourseID: courseID})
	}
}

func FindUserCoursesByUserID(userID string) []string {
	var res []UserCourseDao
	var courses []string
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
	for _, v := range res {
		courses = append(courses, v.CourseID)
	}
	return courses
}

func CheckStudentCourseIsExisted(studentID string, courseID string) bool {
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
		db.Where(&UserCourseDao{UserID: studentID, CourseID: courseID}).Find(&res)
	}
	if res.CourseID != "" {
		return true
	}
	return false
}

// func UpdateUserCoursesByCourseID(courseID string, userID string) Types.ErrNo {
// 	var res UserCourseDao
// 	db, err := DBAccessor.MySqlInit()
// 	defer func(db *gorm.DB) {
// 		_ = db.Close()
// 	}(db)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("Database connection refused.")
// 		return Types.UnknownError
// 	} else {
// 		// 直到建表成功才继续
// 		for true {
// 			if makeUserCourseTable() {
// 				break
// 			} else {
// 				// 如果建表失败，停4s并输出提示信息
// 				time.Sleep(time.Duration(4))
// 				fmt.Println("Something happened when trying to establish the table--'userCourses'.Please check the database.")
// 			}
// 		}
// 		db.Where(&UserCourseDao{UserID: userID}).Find(&res)
// 		if res.UserID == "" {
// 			return Types.UnknownError
// 		}
// 		res.CourseIDs = courseIDs
// 		db.Save(res)
// 		return Types.OK
// 	}
// }

func DeleteUserCoursesByUserID(userID string) Types.ErrNo {
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
