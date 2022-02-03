//package TCourseDaoTest
package main

import (
	"Project1/Dao/TCourseDao"
	"github.com/ByteDanceBackendProject18/Project/Dao/DBAccessor"
	"github.com/ByteDanceBackendProject18/Project/Types"
	_ "github.com/go-sql-driver/mysql"
)

func TestInsertCourse() {
	var course Types.TCourse
	course.CourseID = string(rune(0x0000))
	course.Name = "Goland Backend"
	TCourseDao.InsertCourse(course)
}
func main() {
	DBAccessor.MySqlInit()
	print(1)
	//TestInsertCourse()
}
