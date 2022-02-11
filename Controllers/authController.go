package Controllers

import (
	"Project/Dao/TMemberDao"
	"Project/Dao/UserDao"
	"Project/Types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
}

func (con AuthController) Login(c *gin.Context) {
	loginRequest := &Types.LoginRequest{}
	loginResponse := &Types.LoginResponse{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	//验证用户
	if member, e := TMemberDao.FindMemberByUserName(loginRequest.Username); e != Types.OK {
		loginResponse.Code = Types.WrongPassword
		c.JSON(http.StatusOK, loginResponse)
		return
	} else {
		if isTrue, _ := UserDao.CheckUser(member.UserID, loginRequest.Password); isTrue {
			loginResponse.Code = Types.OK
			loginResponse.Data = struct{ UserID string }{UserID: member.UserID}
		} else {
			loginResponse.Code = Types.WrongPassword
			c.JSON(http.StatusOK, loginResponse)
			return
		}
	}

	c.SetCookie("camp-session", loginRequest.Username, 0, "/", "", false, false)

	c.JSON(http.StatusOK, loginResponse)
}

func (con AuthController) Logout(c *gin.Context) {
	logoutResponse := &Types.LogoutResponse{}

	cookie, err := c.Request.Cookie("camp-session")
	if err == nil {
		c.SetCookie(cookie.Name, cookie.Value, -1, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		logoutResponse.Code = Types.OK
	} else {
		logoutResponse.Code = Types.LoginRequired
	}

	c.JSON(http.StatusOK, logoutResponse)
}

func (con AuthController) WhoAmI(c *gin.Context) {
	whoAmIResponse := Types.WhoAmIResponse{}

	cookie, err := c.Request.Cookie("camp-session")
	if err != nil {
		whoAmIResponse.Code = Types.LoginRequired
		return
	}

	if member, e := TMemberDao.FindMemberByUserName(cookie.Value); e == Types.OK {
		whoAmIResponse.Code = Types.OK
		whoAmIResponse.Data = member
	} else {
		whoAmIResponse.Code = Types.OK
		whoAmIResponse.Data = member
	}

	c.JSON(http.StatusOK, whoAmIResponse)
}
