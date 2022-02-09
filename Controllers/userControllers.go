package Controllers

import (
	"Project/Dao/TMemberDao"
	"Project/Dao/UserDao"
	"Project/Types"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

type UserController struct {
}

func (con UserController) CreateMember(c *gin.Context) {
	createMemberRequest := &Types.CreateMemberRequest{}
	createMemberResponse := &Types.CreateMemberResponse{}

	// 判断管理员身份
	if cookie, err := c.Cookie("camp-session"); err != nil {
		if curMember, e := TMemberDao.FindMemberByUserName(cookie); e != Types.OK || curMember.UserType != Types.Admin {
			createMemberResponse.Code = Types.PermDenied
			createMemberResponse.Data = struct{ UserID string }{UserID: string(0)}
			c.JSON(http.StatusOK, createMemberResponse)
			return
		}
	}

	if err := c.ShouldBindJSON(createMemberRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if !check(4, 20, createMemberRequest.Nickname) || !check(8, 20, createMemberRequest.Username) || !checkPassword(8, 20, createMemberRequest.Password) {
		createMemberResponse.Code = Types.ParamInvalid
		createMemberResponse.Data = struct{ UserID string }{UserID: string(0)}
		c.JSON(http.StatusOK, createMemberResponse)
		return
	}

	if createMemberRequest.UserType != Types.Admin && createMemberRequest.UserType != Types.Student && createMemberRequest.UserType != Types.Teacher {
		createMemberResponse.Code = Types.ParamInvalid
		createMemberResponse.Data = struct{ UserID string }{UserID: string(0)}
		c.JSON(http.StatusOK, createMemberResponse)
		return
	}

	//生成用户id
	s, _ := snowflake.NewSnowflake(int64(0), int64(0))
	userID := s.NextVal()

	member := Types.TMember{
		UserID:   string(userID),
		Nickname: createMemberRequest.Nickname,
		Username: createMemberRequest.Username,
		UserType: createMemberRequest.UserType,
	}

	TMemberDao.InsertMember(member)
	UserDao.InsertUser(string(userID), createMemberRequest.Password)

	createMemberResponse.Code = Types.OK
	createMemberResponse.Data = struct{ UserID string }{UserID: string(userID)}
	c.JSON(http.StatusOK, createMemberResponse)
	return
}

func (con UserController) UpdateMember(c *gin.Context) {

}

func (con UserController) DeleteMember(c *gin.Context) {

}

func (con UserController) GetMember(c *gin.Context) {

}

func (con UserController) GetMemberList(c *gin.Context) {

}

func checkPassword(minLength, maxLength int, pwd string) bool {
	if len(pwd) < minLength || len(pwd) > maxLength {
		return false
	}

	level := 0
	patternList := []string{`[0-9]+`, `[a-z]+`, `[A-Z]+`}
	for _, pattern := range patternList {
		match, _ := regexp.MatchString(pattern, pwd)
		if match {
			level++
		}
	}
	if level == 3 {
		return true
	}
	return false
}

func check(minLength, maxLength int, pwd string) bool {
	if len(pwd) < minLength || len(pwd) > maxLength {
		return false
	}

	return true
}
