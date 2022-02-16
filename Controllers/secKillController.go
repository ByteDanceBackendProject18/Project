package Controllers

import (
	TCourseDao "Project/Dao/TCourseDao"
	"Project/Dao/TMemberDao"
	"Project/Service/SecKill"
	"Project/Types"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type SecKillController struct {
}

var wg sync.WaitGroup

//抢课
func (con SecKillController) secKill(c *gin.Context) {
	secKillRequest := &Types.BookCourseRequest{}
	secKillResponse := &Types.BookCourseResponse{}

	err := c.ShouldBindJSON(&secKillRequest)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//获取课程信息
	course, e1 := TCourseDao.FindCourseByID(secKillRequest.CourseID)
	//判断用户身份
	cookie, err := c.Cookie("camp-session")
	curMember, e2 := TMemberDao.FindMemberByUserName(cookie)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	//当前用户不是学生
	if e2 != Types.OK || curMember.UserType != Types.Student {
		secKillResponse.Code = Types.ParamInvalid
		c.JSON(http.StatusOK, secKillResponse)
		return
	}

	//当前课程未被老师绑定
	if course.TeacherID == "" {
		secKillResponse.Code = Types.CourseNotBind
		c.JSON(http.StatusOK, secKillResponse)
		return
	}

	//课程不存在
	if e1 != Types.OK {
		secKillResponse.Code = Types.CourseNotExisted
		c.JSON(http.StatusOK, secKillResponse)
		return
	}

	//学生查找学生是否选过该课程
	//Todo:add args
	hasCourse, _ := SecKillService.StudentHasCourse(secKillRequest.StudentID, secKillRequest.CourseID)
	if !hasCourse {
		secKillResponse.Code = Types.StudentHasCourse
		return
	}

	//课程存在，检查余量
	//不知道余量参数，CheckResidue未实现
	residueJudge, residue := SecKillService.CheckResidue(secKillRequest.CourseID)
	if !residueJudge {
		//没有余量
		secKillResponse.Code = Types.CourseNotAvailable
		c.JSON(http.StatusOK, secKillResponse)
		return
	}

	wg.Add(residue)
	for i := 0; i < residue; i++ {
		err := SecKillService.HandleSecKillWithLock(course.CourseID, curMember.UserID)
		//抢课失败
		if err != nil {
			secKillResponse.Code = Types.UnknownError
			c.JSON(http.StatusOK, secKillResponse)
			wg.Done()
			return
		}
	}
	wg.Wait()

	secKillResponse.Code = Types.OK
	c.JSON(http.StatusOK, secKillResponse)
	return
}

//获取学生课表
func (con SecKillController) GetStudentCourse(c *gin.Context) {
	studentCourseRequest := &Types.GetStudentCourseRequest{}
	studentCourseResponse := &Types.GetStudentCourseResponse{}
	err := c.ShouldBindJSON(studentCourseRequest)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//检测学生是否存在
	if _, e := TMemberDao.FindMemberByID(studentCourseRequest.StudentID); e != Types.OK {
		studentCourseResponse.Code = Types.StudentNotExisted
		studentCourseResponse.Data = struct{ CourseList []Types.TCourse }{CourseList: nil}
		c.JSON(http.StatusOK, studentCourseResponse)
		return
	}
	//检查学生是否有课程
	course := SecKillService.CheckStudentCourse(studentCourseRequest.StudentID)
	if course == nil {
		studentCourseResponse.Code = Types.StudentHasNoCourse
		studentCourseResponse.Data = struct{ CourseList []Types.TCourse }{CourseList: nil}
		c.JSON(http.StatusOK, studentCourseResponse)
		return
	}
	studentCourseResponse.Code = Types.OK
	studentCourseResponse.Data = struct{ CourseList []Types.TCourse }{CourseList: course}
	c.JSON(http.StatusOK, studentCourseResponse)
	return

}
