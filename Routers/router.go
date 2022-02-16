package Routers

import (
	"Project/Controllers"
	"Project/Dao/RedisAccessor"
	"Project/Dao/TCourseDao/TCourseDaoTest"
	"bytes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1", ginBodyLogMiddleware())

	//使用session

	// 成员管理
	g.POST("/member/create", Controllers.UserController{}.CreateMember)
	g.GET("/member", Controllers.UserController{}.GetMember)
	g.GET("/member/list", Controllers.UserController{}.GetMemberList)
	g.POST("/member/update", Controllers.UserController{}.UpdateMember)
	g.POST("/member/delete", Controllers.UserController{}.DeleteMember)

	// 登录
	g.POST("/auth/login", Controllers.AuthController{}.Login)
	g.POST("/auth/logout", Controllers.AuthController{}.Logout)
	g.GET("/auth/whoami", Controllers.AuthController{}.WhoAmI)

	// 排课
	g.POST("/course/create", Controllers.CourseController{}.CreateCourse)
	g.GET("/course/get", Controllers.CourseController{}.GetOneCourse)

	g.POST("/teacher/bind_course", Controllers.CourseController{}.BindCourse)
	g.POST("/teacher/unbind_course", Controllers.CourseController{}.UnBindCourse)
	g.GET("/teacher/get_course", Controllers.CourseController{}.GetTeacherCourse)
	g.POST("/course/schedule", Controllers.ScheCourseController{}.ScheduleCourse)

	// 抢课
	g.POST("/student/book_course", Controllers.SecKillController{}.SecKill)
	g.GET("/student/course", Controllers.SecKillController{}.GetStudentCourse)

	g.GET("/test", TCourseDaoTest.TCourseDaoTest{}.Test)

	g.GET("/testRedis", RedisAccessor.RedisAccessorTest{}.Test)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func ginBodyLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		fmt.Println("Response body: " + blw.body.String())
	}
}
