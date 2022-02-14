package TCourseDaoTest

//package main

import (
	"Project/Dao/TCourseDao"
	"Project/Types"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type TCourseDaoTest struct {
}

var courseID = "0001"
var courseName = "Goland Backend"
var teacherID = "1001"
var newCourse Types.TCourse

func TestInsertCourse() {
	var course Types.TCourse
	course.CourseID = courseID
	course.Name = courseName
	course.TeacherID = teacherID
	TCourseDao.InsertCourse(course)
}
func TestDeleteCourseByID() {
	TCourseDao.DeleteCourseByID(courseID)
}
func TestDeleteCourseByName() {
	TCourseDao.DeleteCoursesByName(courseName)
}
func TestUpdateCourseByID() {
	TCourseDao.UnsafeUnbindTeacherIDOfCourse("")
	TCourseDao.UnsafeDeleteCourses("", "", "")
	TCourseDao.UnbindCourseByTeacherID("")
	TCourseDao.UnsafeUpdateCourseByID(courseID, newCourse)
}
func TestUpdateTeacherIDOfCourse() Types.ErrNo {
	return TCourseDao.UpdateTeacherIDOfCourse(courseID, "1002")
}
func TestFindCourseByID() {
	fmt.Println(TCourseDao.FindCourseByID("0001"))
}
func TestFindCourseByTeacherID() {
	fmt.Println(TCourseDao.FindCourseByTeacherID("1002"))
}
func TestFindCoursesByName() {
	fmt.Println(TCourseDao.FindCoursesByName(courseName))
}
func TestDao() {
	newCourse.CourseID = courseID
	newCourse.Name = courseName
	dao := TCourseDao.MakeTCourseDao(newCourse)
	TCourseDao.InsertCourseByDao(dao)
	newCourse.TeacherID = "1002"
	TCourseDao.UpdateCourseByDao(dao, newCourse)
	TCourseDao.DeleteCourseByDao(dao)
}
func TestInsertCourses() {
	var courses []Types.TCourse
	newCourse.CourseID = courseID
	newCourse.Name = courseName
	courses = append(courses, newCourse)
	newCourse.CourseID = "0002"
	newCourse.TeacherID = teacherID
	courses = append(courses, newCourse)
	TCourseDao.InsertCourses(courses)
}
func (t TCourseDaoTest) Test(c *gin.Context) {
	TestInsertCourse()                         //test successfully
	TestDeleteCourseByID()                     //test successfully
	TestDeleteCourseByName()                   //test successfully
	TestUpdateCourseByID()                     //test successfully
	fmt.Println(TestUpdateTeacherIDOfCourse()) //test successfully
	TestFindCourseByID()                       //test successfully
	TestFindCourseByTeacherID()                //test successfully
	TestFindCoursesByName()                    //test successfully
	TestDao()                                  //test successfully
	TestInsertCourses()                        //test successfully

	c.JSON(http.StatusOK, gin.H{})
}
