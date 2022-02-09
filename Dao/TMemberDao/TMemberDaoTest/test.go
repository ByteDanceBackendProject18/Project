package TMemberDaoTest

//package main

import (
	"Project/Dao/TMemberDao"
	"Project/Types"
	"fmt"
)

var id = "0001"
var username = "OceanCT"
var member Types.TMember = Types.TMember{UserType: Types.Student, UserID: id, Username: username, Nickname: username}

func TestInsertMember() {
	TMemberDao.InsertMember(member)
}
func TestInsertMembers() {
	var members []Types.TMember
	members = append(members, member)
	member.UserID = "0002"
	members = append(members, member)
	TMemberDao.InsertMembers(members)
}
func TestDeleteMemberByID() {
	TMemberDao.DeleteMemberByID(id)
}
func TestDeleteMemberByUserName() {
	TMemberDao.DeleteMemberByUserName(username)
}
func TestFindMemberByID() {
	fmt.Println(TMemberDao.FindMemberByID(id))
}
func TestFindMemberByUserName() {
	fmt.Println(TMemberDao.FindMemberByUserName(username))
}
func TestFindMemberByNickName() {
	fmt.Println(TMemberDao.FindMemberByNickName(username))
}
func TestFindMemberByUserType() {
	TestInsertMembers()
	fmt.Println(TMemberDao.FindMemberByUserType(Types.Student))
}
func TestUpdateNickNameByName() {
	TMemberDao.UpdateNickNameByName(username, "OceanRE")
}
func TestUpdateNickNameByID() {
	TMemberDao.UpdateNickNameByID(id, "OceanKT")
}
func TestDao() {
	dao := TMemberDao.MakeTMemberDao(member)
	TMemberDao.InsertMemberByDao(dao)
	TMemberDao.UpdateNickNameByDao(dao, "OceanRE")
	TMemberDao.DeleteMemberByDao(dao)
}
func main() {
	//TestInsertMember() //test successfully
	//TestDeleteMemberByID()//test successfully
	//TestInsertMembers() //test successfully
	//TestDeleteMemberByUserName() //test successfully
	//TestFindMemberByID() //test successfully
	//TestFindMemberByUserName() //test successfully
	//TestFindMemberByNickName() //test successfully
	//TestFindMemberByUserType() //test successfully
	//TestUpdateNickNameByName() //test successfully
	//TestUpdateNickNameByID() //test successfully
	//TestDao()//test successfully
}
