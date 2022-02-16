package Controllers

import (
	"Project/Dao/CapDao"
	"Project/Dao/TCourseDao"
	"Project/Types"
	"net/http"
	"strconv"

	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
}

func (con CourseController) CreateCourse(c *gin.Context) {
	createCourseRequest := &Types.CreateCourseRequest{}
	createCourseResponse := &Types.CreateCourseResponse{}

	if err := c.ShouldBindJSON(&createCourseRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	s, _ := snowflake.NewSnowflake(int64(0), int64(0))
	courseID := s.NextVal()

	course := Types.TCourse{
		CourseID:  strconv.FormatInt(courseID, 10), //
		Name:      createCourseRequest.Name,
		TeacherID: "", //
	}

	TCourseDao.InsertCourse(course)
	CapDao.InsertCap(strconv.FormatInt(courseID, 10), createCourseRequest.Cap)

	createCourseResponse.Code = Types.OK
	createCourseResponse.Data = struct{ CourseID string }{CourseID: strconv.FormatInt(courseID, 10)}
	c.JSON(http.StatusOK, createCourseResponse)
}

func (con CourseController) GetOneCourse(c *gin.Context) {
	getCourseRequest := &Types.GetCourseRequest{}
	getCourseResponse := &Types.GetCourseResponse{}

	if err := c.ShouldBindQuery(&getCourseRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	course, e := TCourseDao.FindCourseByID(getCourseRequest.CourseID)
	if e != Types.OK {
		getCourseResponse.Code = Types.CourseNotExisted
		c.JSON(http.StatusOK, getCourseResponse)
		return
	}
	getCourseResponse.Data = course
	c.JSON(http.StatusOK, getCourseResponse)
}

func (con CourseController) BindCourse(c *gin.Context) {
	bindCourseRequest := &Types.BindCourseRequest{}
	bindCourseResponse := &Types.BindCourseResponse{}

	if err := c.ShouldBindJSON(&bindCourseRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	course, e := TCourseDao.FindCourseByID(bindCourseRequest.CourseID)
	if e != Types.OK {
		bindCourseResponse.Code = Types.CourseNotExisted
		c.JSON(http.StatusOK, bindCourseResponse)
		return
	}

	if course.TeacherID != "" {
		bindCourseResponse.Code = Types.CourseHasBound
		c.JSON(http.StatusOK, bindCourseResponse)
		return
	}
	TCourseDao.UpdateTeacherIDOfCourse(bindCourseRequest.CourseID, bindCourseRequest.TeacherID)
	bindCourseResponse.Code = Types.OK
	c.JSON(http.StatusOK, bindCourseResponse)
}

func (con CourseController) UnBindCourse(c *gin.Context) {
	unbindCourseRequest := &Types.UnbindCourseRequest{}
	unbindCourseResponse := &Types.UnbindCourseResponse{}

	if err := c.ShouldBindJSON(&unbindCourseRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	e := TCourseDao.UnbindTeacherIDOfCourse(unbindCourseRequest.CourseID, unbindCourseRequest.TeacherID)
	unbindCourseResponse.Code = e
	c.JSON(http.StatusOK, unbindCourseResponse)
}

func (con CourseController) GetTeacherCourse(c *gin.Context) {
	getTeacherCourseRequest := &Types.GetTeacherCourseRequest{}
	getTeacherCourseResponse := &Types.GetTeacherCourseResponse{}

	if err := c.ShouldBindQuery(&getTeacherCourseRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	course, e := TCourseDao.FindCourseByTeacherID(getTeacherCourseRequest.TeacherID)
	if e != Types.OK {
		getTeacherCourseResponse.Code = Types.CourseNotExisted
		c.JSON(http.StatusOK, getTeacherCourseResponse)
		return
	}

	getTeacherCourseResponse.Code = Types.OK
	getTeacherCourseResponse.Data = struct{ CourseList []*Types.TCourse }{CourseList: course}
	c.JSON(http.StatusOK, getTeacherCourseResponse)
}
