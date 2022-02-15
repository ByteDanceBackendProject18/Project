package SecKillService

import (
	"Project/Dao/CapDao"
	"Project/Dao/TCourseDao"
	"Project/Types"
	"errors"
	"sync"
)

var lock = sync.Mutex{}

//检查余量
func CheckResidue(CourseID string) (bool, int) {
	//不存在返回false
	if _, e := TCourseDao.FindCourseByID(CourseID); e != Types.OK {
		return false, 0
	}

	//存在返回true
	return true, CapDao.FindCapByCourseID(CourseID)
}

//检查学生是否选择课程
func StudentHasCourse(UserID string, CourseID string) bool {
	//未选返回true

	//选择返回false
	return false
}

//具体操作
func HandleSecKill(CourseID string, UserID string) error {
	//调用dao，返回error
	return errors.New("")
}

//高并发锁
func HandleSecKillWithLock(CourseID string, UserID string) error {
	lock.Lock()
	err := HandleSecKill(CourseID, UserID)
	lock.Unlock()
	return err
}

//获取学生课程表
func CheckStudentCourse(StudentID string) (courses []Types.TCourse) {
	return
}
