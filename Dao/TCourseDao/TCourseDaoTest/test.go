//package TCourseDaoTest
package main

import (
	"Project1/Dao/TCourseDao"
	"github.com/ByteDanceBackendProject18/Project/Types"
)

type Test struct {
	a int
}

func TestInsertCourse() {
	var course Types.TCourse
	course.CourseID = string(rune(0x0000))
	course.Name = "Goland Backend"
	TCourseDao.InsertCourse(course)
}
func main() {
	//testInsertCourse()
}
