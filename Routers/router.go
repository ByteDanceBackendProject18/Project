package Routers

import (
	"Project1/Controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	g := r.Group("/api/v1")

	// 成员管理
	g.POST("/member/create", Controllers.UserController{}.CreateUser)
	g.GET("/member", Controllers.UserController{}.GetSingleUser)
	g.GET("/member/list", Controllers.UserController{}.GetUserList)
	g.POST("/member/update", Controllers.UserController{}.UpdateUser)
	g.POST("/member/delete", Controllers.UserController{}.DeleteUser)

	// 登录
	g.POST("/auth/login")
	g.POST("/auth/logout")
	g.GET("/auth/whoami")

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
