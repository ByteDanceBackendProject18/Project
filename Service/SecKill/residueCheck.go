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
func CheckResidue(courseID string) (int, error) {
	residue := CapDao.FindCapByCourseID(courseID)
	if residue < 1 {
		//没库存了
		err := CapDao.HasNoCapForRedis(courseID)
		if err != nil {
			return 0, errors.New("fail")
		}
	}
	//存在返回true
	return residue, nil
}

// StudentHasCourse 检查学生是否已选该课程
func StudentHasCourse(userID string, courseID string) bool {
	//未选返回true
	//选择返回false

	return !UserCourseDao.CheckStudentCourseIsExisted(userID, courseID)
}

// StudentInsertCourse 学生增加课程
func StudentInsertCourse(userID, courseID string) {
	if StudentHasCourse(userID, courseID) {
		UserCourseDao.InsertUserCourse(courseID, userID)
	}
}

// HandleSecKill 具体操作
func HandleSecKill(courseID string, userID string) error {
	//检查余量
	residue, err := CheckResidue(courseID)
	if err != nil {
		return errors.New("fail")
	}
	if residue < 1 {
		//没有余量
		return errors.New("HasNoCap")
	} else {
		StudentInsertCourse(userID, courseID)
		CapDao.UpdateCapByCourseID(courseID, residue-1)
	}
	//调用dao，返回error
	return nil
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
