package Controllers

import (
	"Project/Dao/TMemberDao"
	"Project/Dao/UserDao"
	"Project/Service/UserService"
	"Project/Types"
	"net/http"
	"strconv"

	"github.com/GUAIK-ORG/go-snowflake/snowflake"
	"github.com/gin-gonic/gin"
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

	if !UserService.CheckNickName(createMemberRequest.Nickname) || !UserService.CheckUserName(createMemberRequest.Username) || !UserService.CheckPassword(createMemberRequest.Password) {
		createMemberResponse.Code = Types.ParamInvalid
		createMemberResponse.Data = struct{ UserID string }{UserID: strconv.Itoa(0)}
		c.JSON(http.StatusOK, createMemberResponse)
		return
	}

	if member, isExisted := TMemberDao.FindMemberByUserName(createMemberRequest.Username); isExisted == Types.OK {
		createMemberResponse.Code = Types.UserHasExisted
		createMemberResponse.Data = struct{ UserID string }{UserID: member.UserID}
		c.JSON(http.StatusOK, createMemberResponse)
		return
	}

	if !UserService.CheckUserType(createMemberRequest.UserType) {
		createMemberResponse.Code = Types.ParamInvalid
		createMemberResponse.Data = struct{ UserID string }{UserID: strconv.Itoa(0)}
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

	if !UserService.CheckNickName(updateMemberRequest.Nickname) {
		updateMemberResponse.Code = Types.ParamInvalid
		return
	}

	if e := TMemberDao.UpdateNickNameByID(member.UserID, updateMemberRequest.Nickname); e != Types.OK {
		updateMemberResponse.Code = Types.UnknownError
	} else {
		updateMemberResponse.Code = Types.OK
	}
	c.JSON(http.StatusOK, updateMemberResponse)
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
		getMemberResponse.Code = TMemberDao.TellMemberExistedBefore(member.Username)
	} else {
		getMemberResponse.Code = Types.OK
		getMemberResponse.Data = member
	}

	c.JSON(http.StatusOK, getMemberResponse)
}

func (con UserController) GetMemberList(c *gin.Context) {
	getMemberListRequest := Types.GetMemberListRequest{}
	getMemberListResponse := Types.GetMemberListResponse{}

	if err := c.ShouldBindQuery(&getMemberListRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	memberList, e := TMemberDao.GetMemberList(getMemberListRequest.Offset, getMemberListRequest.Limit)

	getMemberListResponse.Code = e
	getMemberListResponse.Data = struct{ MemberList []Types.TMember }{MemberList: memberList}
	c.JSON(http.StatusOK, getMemberListResponse)
}
