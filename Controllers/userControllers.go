package Controllers

import (
	"Project/Dao/TMemberDao"
	"Project/Dao/UserDao"
	"Project/Types"
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

type UserController struct {
}

func (con UserController) CreateMember(c *gin.Context) {
	createMemberRequest := &Types.CreateMemberRequest{}
	createMemberResponse := &Types.CreateMemberResponse{}

	// 判断管理员身份
	if cookie, err := c.Cookie("camp-session"); err != nil {
		if curMember, e := TMemberDao.FindMemberByUserName(cookie); e == Types.OK && curMember.UserType == Types.Admin {
			createMemberResponse.Code = Types.PermDenied
			createMemberResponse.Data = struct{ UserID string }{UserID: curMember.UserID}
			c.JSON(http.StatusOK, createMemberResponse)
			return
		}
	}

	if err := c.ShouldBindJSON(&createMemberRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if !check(4, 20, createMemberRequest.Nickname) || !check(8, 20, createMemberRequest.Username) || !checkPassword(8, 20, createMemberRequest.Password) {
		createMemberResponse.Code = Types.ParamInvalid
		createMemberResponse.Data = struct{ UserID string }{UserID: string(0)}
		c.JSON(http.StatusOK, createMemberResponse)
		return
	}

	if member, isExisted := TMemberDao.FindMemberByUserName(createMemberRequest.Username); isExisted == Types.OK {
		createMemberResponse.Code = Types.UserHasExisted
		createMemberResponse.Data = struct{ UserID string }{UserID: member.UserID}
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
	ID := s.NextVal()
	userID := strconv.FormatInt(ID, 10)

	member := Types.TMember{
		UserID:   userID,
		Nickname: createMemberRequest.Nickname,
		Username: createMemberRequest.Username,
		UserType: createMemberRequest.UserType,
	}

	TMemberDao.InsertMember(member)
	UserDao.InsertUser(userID, createMemberRequest.Password)

	createMemberResponse.Code = Types.OK
	createMemberResponse.Data = struct{ UserID string }{UserID: userID}
	c.JSON(http.StatusOK, createMemberResponse)
	return
}

func (con UserController) UpdateMember(c *gin.Context) {
	updateMemberRequest := Types.UpdateMemberRequest{}
	updateMemberResponse := Types.UpdateMemberResponse{}

	if err := c.ShouldBindJSON(&updateMemberRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	member, e := TMemberDao.FindMemberByID(updateMemberRequest.UserID)
	if e != Types.OK {
		updateMemberResponse.Code = Types.UserNotExisted
		c.JSON(http.StatusOK, updateMemberResponse)
		return
	}

	if e := TMemberDao.UpdateNickNameByID(member.UserID, updateMemberRequest.Nickname); e != Types.OK {
		updateMemberResponse.Code = Types.UnknownError
	} else {
		updateMemberResponse.Code = Types.OK
	}
	c.JSON(http.StatusOK, updateMemberResponse)
	return
}

func (con UserController) DeleteMember(c *gin.Context) {
	deleteMemberRequest := Types.DeleteMemberRequest{}
	deleteMemberResponse := Types.DeleteMemberResponse{}

	if err := c.ShouldBindJSON(&deleteMemberRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	member, e := TMemberDao.FindMemberByID(deleteMemberRequest.UserID)
	if e != Types.OK {
		deleteMemberResponse.Code = Types.UserNotExisted
		c.JSON(http.StatusOK, deleteMemberResponse)
		return
	}

	e = TMemberDao.DeleteMemberByID(member.UserID)
	deleteMemberResponse.Code = e
	c.JSON(http.StatusOK, deleteMemberResponse)
	return
}

func (con UserController) GetMember(c *gin.Context) {
	getMemberRequest := Types.GetMemberRequest{}
	getMemberResponse := Types.GetMemberResponse{}

	if err := c.ShouldBindQuery(&getMemberRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	member, e := TMemberDao.FindMemberByID(getMemberRequest.UserID)
	if e != Types.OK {
		getMemberResponse.Code = Types.UserNotExisted
	} else {
		getMemberResponse.Code = Types.OK
		getMemberResponse.Data = member
	}

	c.JSON(http.StatusOK, getMemberResponse)
	return
}

func (con UserController) GetMemberList(c *gin.Context) {
	//getMemberListRequest := Types.GetMemberListRequest{}
	//getMemberListResponse := Types.GetMemberListResponse{}
	//
	//if err := c.ShouldBindQuery(&getMemberListRequest); err != nil {
	//	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	//	return
	//}
	//
	//
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
