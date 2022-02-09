package UserDaoTest

import (
	"Project/Dao/UserDao"
	"fmt"
)

//package main

var userID = "0001"
var pw = "123456"

func TestInsertUser() {
	UserDao.InsertUser(userID, pw)
}
func TestCheckUser() {
	fmt.Println(UserDao.CheckUser(userID, "123"))
	fmt.Println(UserDao.CheckUser(userID, pw))
}
func TestUpdatePassword() {
	UserDao.UpdatePassword(userID, "123")
}
func main() {
	//TestInsertUser()//test successfully
	//TestCheckUser() //test successfully
	//TestUpdatePassword() //test successfully
}
