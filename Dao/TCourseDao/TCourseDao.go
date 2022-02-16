package TCourseDao

import (
	"Project/Dao/DBAccessor"
	"Project/Types"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type TCourseDao struct {
	CourseID  string `gorm:"type:varchar(128)"`
	Name      string `gorm:"type:varchar(256)"`
	TeacherID string `gorm:"type:varchar(256)"`
	gorm.Model
}

// TableName 改表名为“courses”
func (TCourseDao) TableName() string {
	return "courses"
}

// convertCourseDaoToCourse 将Dao转换为course
func convertCourseDaoToCourse(dao TCourseDao) Types.TCourse {
	var course Types.TCourse
	course.CourseID = dao.CourseID
	course.Name = dao.Name
	course.TeacherID = dao.TeacherID
	return course
}

// MakeTCourseDao 提供MakeTCourseDao接口，如果需要对同一个Course反复操作，可以使用该接口获取Dao类型指针
func MakeTCourseDao(course Types.TCourse) *TCourseDao {
	var res *TCourseDao = new(TCourseDao)
	res.CourseID = course.CourseID
	res.Name = course.Name
	res.TeacherID = course.TeacherID
	return res
}

// makeCourseTable 如果表不存在就建表,并返回最终是否有该表
func makeCourseTable() bool {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		if !db.HasTable(&TCourseDao{}) {
			db.AutoMigrate(&TCourseDao{})
		}
	}
	if db.HasTable(&TCourseDao{}) {
		return true
	} else {
		return false
	}
}

// InsertCourse 添加一门课程
func InsertCourse(course Types.TCourse) {
	dao := MakeTCourseDao(course)
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		db.Create(dao)
	}
}

// InsertCourses 添加多门课程
func InsertCourses(courses []Types.TCourse) {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		for _, course := range courses {
			db.Create(MakeTCourseDao(course))
		}
	}
}

// InsertCourseByDao 使用Dao指针添加一门课程
func InsertCourseByDao(dao *TCourseDao) {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		db.Create(dao)
	}
}

// FindCoursesByName 根据课程的名字找到对应的课程
func FindCoursesByName(name string) ([]Types.TCourse, Types.ErrNo) {
	var res []TCourseDao
	var res1 []Types.TCourse
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		db.Where(&TCourseDao{Name: name}).Find(&res)
		for _, course := range res {
			res1 = append(res1, convertCourseDaoToCourse(course))
		}
	}
	return res1, 0
}

// FindCourseByID 根据CourseID找到对应的唯一的课程
func FindCourseByID(id string) (Types.TCourse, Types.ErrNo) {
	var res TCourseDao
	var res1 Types.TCourse
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		db.Where(&TCourseDao{CourseID: id}).Find(&res)
		return convertCourseDaoToCourse(res), Types.OK
	}
}

// FindCourseByTeacherID 根据TeacherID找到对应的课程
func FindCourseByTeacherID(id string) ([]*Types.TCourse, Types.ErrNo) {
	var res []TCourseDao
	var res1 []*Types.TCourse
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		db.Where(&TCourseDao{TeacherID: id}).Find(&res)
		for _, course := range res {
			cse := convertCourseDaoToCourse(course)
			res1 = append(res1, &cse)
		}
	}
	return res1, 0
}

// UpdateTeacherIDOfCourse 将ID对应的course的执教教师ID更新为传入的老师的ID(对是否已经绑定做检查)
func UpdateTeacherIDOfCourse(courseId string, teacherID string) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var res TCourseDao
		db.Where(&TCourseDao{CourseID: courseId}).Find(&res)
		if res.Name == "" {
			return Types.CourseNotExisted
		}
		if res.TeacherID != "" {
			return Types.CourseHasBound
		}
		db.Model(&res).Update(map[string]interface{}{"TeacherID": teacherID})
		return Types.OK
	}
}

// UnbindTeacherIDOfCourse 将ID对应的course的执教教师ID解绑(检查teacherID是否正确)
func UnbindTeacherIDOfCourse(courseId string, teacherID string) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var res TCourseDao
		db.Where(&TCourseDao{CourseID: courseId}).Find(&res)
		if res.Name == "" {
			return Types.CourseNotExisted
		}
		if res.TeacherID == "" {
			return Types.CourseNotBind
		}
		// 如果这个课程的教师ID与传入参数不符返回已绑定
		if res.TeacherID != teacherID {
			return Types.CourseHasBound
		}
		db.Model(&res).Update(map[string]interface{}{"TeacherID": ""})
		return Types.OK
	}
}

// UnsafeUnbindTeacherIDOfCourse 将ID对应的course的执教教师ID解绑
func UnsafeUnbindTeacherIDOfCourse(courseId string) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var res TCourseDao
		db.Where(&TCourseDao{CourseID: courseId}).Find(&res)
		if res.Name == "" {
			return Types.CourseNotExisted
		}
		if res.TeacherID == "" {
			return Types.CourseHasBound
		}
		res.TeacherID = ""
		db.Model(&res).Updates(res)
		return Types.OK
	}
}

// UnsafeUpdateCourseByID 根据CourseID更新对应课程的信息(不对是否已经绑定做检查)
func UnsafeUpdateCourseByID(courseID string, course Types.TCourse) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var res TCourseDao
		db.Where(&TCourseDao{CourseID: courseID}).Find(&res)
		if res.Name == "" {
			return Types.CourseNotExisted
		}
		res.CourseID = course.CourseID
		res.Name = course.Name
		res.TeacherID = course.TeacherID
		db.Model(&res).Updates(res)
		return 0
	}
}

// UnbindCourseByTeacherID 解绑这个老师的所有课程（只解绑一个课程参考其他Unbind函数）
func UnbindCourseByTeacherID(teacherID string) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var res TCourseDao
		db.Where(&TCourseDao{TeacherID: teacherID}).Find(&res)
		if res.Name == "" {
			return Types.CourseNotExisted
		}
		if res.TeacherID == "" {
			return Types.CourseNotBind
		}
		res.TeacherID = ""
		db.Model(&res).Updates(res)
		return Types.OK
	}
}

// UpdateCourseByDao 根据Dao指针更新对应课程的信息(对是否已经绑定做检查)
func UpdateCourseByDao(dao *TCourseDao, course Types.TCourse) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		if dao.TeacherID != "" {
			return Types.CourseHasBound
		}
		if dao.Name == "" {
			return Types.CourseNotExisted
		}
		db.Model(dao).Updates(TCourseDao{CourseID: course.CourseID, Name: course.Name, TeacherID: course.TeacherID})
		return 0
	}
}

// DeleteCourseByID 将ID对应的course软删除
func DeleteCourseByID(courseID string) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var res TCourseDao
		db.Where(&TCourseDao{CourseID: courseID}).Find(&res)
		fmt.Println(res)
		if res.Name == "" {
			return Types.CourseNotExisted
		}
		db.Delete(&res)
		return 0
	}
}

// DeleteCoursesByName 将Name对应courses软删除
func DeleteCoursesByName(courseName string) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		var res []TCourseDao
		db.Where(&TCourseDao{Name: courseName}).Find(&res)
		if len(res) == 0 {
			return Types.CourseNotExisted
		}
		for _, i := range res {
			db.Delete(&i)
		}
		return 0
	}
}

// UnsafeDeleteCourses 将满足任一条件的Courses从数据库中软删除
func UnsafeDeleteCourses(courseID string, courseName string, teacherID string) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		db.Where("courseID ?", courseID).Or("name ?", courseName).Or("teacherID ?", teacherID).Delete(TCourseDao{})
		return 0
	}
}

// DeleteCourseByDao 根据Dao指针软删除
func DeleteCourseByDao(dao *TCourseDao) Types.ErrNo {
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
			if makeCourseTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'courses'.Please check the database.")
			}
		}
		db.Delete(dao)
		return 0
	}
}
