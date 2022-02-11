package UserService

import (
	"Project/Types"
	"regexp"
)

func CheckNickName(nickname string) bool {
	reg, _ := regexp.Compile("^.{4,20}$")
	return reg.MatchString(nickname)
}

func CheckUserName(username string) bool {
	reg, _ := regexp.Compile("^[A-Za-z]{8,20}$")
	return reg.MatchString(username)
}

func CheckPassword(password string) bool {
	reg, _ := regexp.Compile("^[A-Za-z0-9]{8,20}$")
	return reg.MatchString(password)
}

func CheckUserType(userType Types.UserType) bool {
	if userType < 1 || userType > 3 {
		return false
	}
	return true
}
