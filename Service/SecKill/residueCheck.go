package SecKillService

import (
	"Project/Dao/CapDao"
	"Project/Dao/TCourseDao"
	"Project/Dao/UserCourseDao"
	"Project/Types"
	"errors"
	"sync"
)

var lock = sync.Mutex{}

// CheckResidue 检查余量
func CheckResidue(courseID string) (bool, int) {
	//不存在返回false
	if _, e := TCourseDao.FindCourseByID(courseID); e != Types.OK {
		return false, 0
	}

	//存在返回true
	return true, CapDao.FindCapByCourseID(courseID)
}

// StudentHasCourse 检查学生是否选择课程
func StudentHasCourse(userID string, courseID string) (bool, []string) {
	coursesList := UserCourseDao.FindUserCoursesByUserID(userID)
	//未选返回true
	//选择返回false
	for _, v := range coursesList {
		if v == courseID {
			return false, coursesList
		}
	}

	return true, coursesList
}

// StudentInsertCourse 学生增加课程
func StudentInsertCourse(userID, courseID string) {
	if b, courses := StudentHasCourse(userID, courseID); b {
		if UserCourseDao.UpdateUserCoursesByCourseID(courses, userID) != Types.OK {
			var courseIDs []string
			UserCourseDao.InsertUserCourse(courseIDs, userID)
		} else {
			courses = append(courses, courseID)
			UserCourseDao.UpdateUserCoursesByCourseID(courses, userID)
		}
	}
}

// HandleSecKill 具体操作
func HandleSecKill(courseID string, userID string) error {
	//调用dao，返回error
	return errors.New("")
}

// HandleSecKillWithLock 高并发锁
func HandleSecKillWithLock(courseID string, userID string) error {
	lock.Lock()
	err := HandleSecKill(courseID, userID)
	lock.Unlock()
	return err
}

// CheckStudentCourse 获取学生课程表
func CheckStudentCourse(studentID string) []Types.TCourse {
	coursesList := UserCourseDao.FindUserCoursesByUserID(studentID)
	var courses []Types.TCourse
	for _, v := range coursesList {
		c, _ := TCourseDao.FindCourseByID(v)
		courses = append(courses, c)
	}

	return courses
}
