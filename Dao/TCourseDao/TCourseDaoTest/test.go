//package TCourseDaoTest
package main

import (
	"Project1/Dao/DBAccessor"
	_ "github.com/go-sql-driver/mysql"
)

//func TestInsertCourse() {
//	var course Types.TCourse
//	course.CourseID = string(rune(0x0000))
//	course.Name = "Goland Backend"
//	TCourseDao.InsertCourse(course)
//}
func main() {
	DBAccessor.MySqlInit()
	//db, err := DBAccessor.MySqlInit()
	//defer func(db *gorm.DB) {
	//	_ = db.Close()
	//}(db)
	//if err != nil {
	//	fmt.Println(err)
	//	fmt.Println("Database connection refused.")
	//} else {
	//	// 直到建表成功才继续
	//}
}
