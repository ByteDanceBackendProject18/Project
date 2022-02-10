package Routers

import (
	"Project/Controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

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
	g.POST("/course/create")
	g.GET("/course/get")

	g.POST("/teacher/bind_course")
	g.POST("/teacher/unbind_course")
	g.GET("/teacher/get_course")
	g.POST("/course/schedule")

	// 抢课
	g.POST("/student/book_course")
	g.GET("/student/course")
}
