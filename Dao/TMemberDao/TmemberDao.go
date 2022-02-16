package TMemberDao

import (
	"Project/Dao/DBAccessor"
	"Project/Types"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TMemberDao struct {
	UserID   string         `gorm:"type:varchar(128)"`
	UserName string         `gorm:"type:varchar(256)"`
	Nickname string         `gorm:"type:varchar(256)"`
	UserType Types.UserType `gorm:"type:varchar(256)"`
	gorm.Model
}

// TableName 改表名为“members”
func (TMemberDao) TableName() string {
	return "members"
}

// makeMemberTable 如果表不存在就建表,并返回最终是否有该表
func makeMemberTable() bool {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		if !db.HasTable(&TMemberDao{}) {
			db.AutoMigrate(&TMemberDao{})
		}
	}
	if db.HasTable(&TMemberDao{}) {
		return true
	} else {
		return false
	}
}

// MakeTMemberDao 提供MakeTMemberDao接口，如果需要对同一个member反复操作，可以使用该接口获取Dao类型指针
func MakeTMemberDao(member Types.TMember) *TMemberDao {
	var res *TMemberDao = new(TMemberDao)
	res.UserName = member.Username
	res.UserID = member.UserID
	res.UserType = member.UserType
	res.Nickname = member.Nickname
	return res
}

// convertMemberDaoToMember 将Dao转换为member
func convertMemberDaoToMember(dao TMemberDao) Types.TMember {
	var res Types.TMember
	res.UserID = dao.UserID
	res.UserType = dao.UserType
	res.Nickname = dao.Nickname
	res.Username = dao.UserName
	return res
}

// InsertMember 插入一条Member
func InsertMember(member Types.TMember) {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Create(MakeTMemberDao(member))
	}
}

// InsertMembers 批量插入Member
func InsertMembers(members []Types.TMember) {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		for _, member := range members {
			db.Create(MakeTMemberDao(member))
		}
	}
}

// InsertMemberByDao 使用Dao指针插入一条数据
func InsertMemberByDao(dao *TMemberDao) {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Create(dao)
	}
}

// FindMemberByID 根据memberID找到对应的唯一的课程
func FindMemberByID(id string) (Types.TMember, Types.ErrNo) {
	var res TMemberDao
	var res1 Types.TMember
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Where(&TMemberDao{UserID: id}).Find(&res)
		res1 = convertMemberDaoToMember(res)
		if res1.Username == "" {
			return res1, Types.UserNotExisted
		}
		return res1, Types.OK
	}
}

// FindMemberByUserName 根据member的Username找对应的唯一的课程
func FindMemberByUserName(name string) (Types.TMember, Types.ErrNo) {
	var res TMemberDao
	var res1 Types.TMember
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Where(&TMemberDao{UserName: name}).Find(&res)
		res1 = convertMemberDaoToMember(res)
		if res1.Username == "" {
			return res1, Types.UserNotExisted
		}
		return res1, Types.OK
	}
}

// FindMemberByNickName 根据UserType找到对应的的课程
func FindMemberByNickName(name string) ([]Types.TMember, Types.ErrNo) {
	var res []TMemberDao
	var res1 []Types.TMember
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Where(&TMemberDao{UserName: name}).Find(&res)
		if len(res) == 0 {
			return res1, Types.UserNotExisted
		}
		for _, i := range res {
			res1 = append(res1, convertMemberDaoToMember(i))
		}
		return res1, Types.OK
	}
}

// FindMemberByUserType 根据UserType找到对应的的课程
func FindMemberByUserType(usertype Types.UserType) ([]Types.TMember, Types.ErrNo) {
	var res []TMemberDao
	var res1 []Types.TMember
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Where(&TMemberDao{UserType: usertype}).Find(&res)
		if len(res) == 0 {
			return res1, Types.UserNotExisted
		}
		for _, i := range res {
			res1 = append(res1, convertMemberDaoToMember(i))
		}
		return res1, Types.OK
	}
}

// UpdateNickNameByName 根据UserName更新Nickname
func UpdateNickNameByName(name string, nickname string) Types.ErrNo {
	var res TMemberDao
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Where(&TMemberDao{UserName: name}).Find(&res)
		if res.UserType == 0 {
			return Types.UserNotExisted
		}
		res.Nickname = nickname
		db.Model(&res).Update("nickname", nickname)
		return 0
	}
}

// UpdateNickNameByID 根据UserID更新NickName
func UpdateNickNameByID(userid string, nickname string) Types.ErrNo {
	var res TMemberDao
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Where(&TMemberDao{UserID: userid}).Find(&res)
		if res.UserType == 0 {
			return Types.UserNotExisted
		}
		res.Nickname = nickname
		db.Model(&res).Update("nickname", nickname)
		return 0
	}
}

// UpdateNickNameByDao 根据Dao指针更新NickName
func UpdateNickNameByDao(dao *TMemberDao, nickname string) Types.ErrNo {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		dao.Nickname = nickname
		db.Model(dao).Update("nickname", nickname)
		return 0
	}
}

// DeleteMemberByID 根据UserID软删除
func DeleteMemberByID(memberID string) Types.ErrNo {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		var res TMemberDao
		db.Where(&TMemberDao{UserID: memberID}).Find(&res)
		if res.UserType == 0 {
			return Types.UserNotExisted
		}
		db.Delete(&res)
		return 0
	}
}

// DeleteMemberByUserName 根据UserName软删除
func DeleteMemberByUserName(name string) Types.ErrNo {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		var res TMemberDao
		db.Where(&TMemberDao{UserName: name}).Find(&res)
		if res.UserType == 0 {
			return Types.UserNotExisted
		}
		db.Delete(&res)
		return 0
	}
}

// DeleteMemberByDao 根据Dao指针软删除
func DeleteMemberByDao(dao *TMemberDao) Types.ErrNo {
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Delete(dao)
		return 0
	}
}

// TellMemberExistedBefore 判断UserName是否曾经存在过
func TellMemberExistedBefore(name string) Types.ErrNo {
	var res TMemberDao
	var res1 Types.TMember
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Unscoped().Where(&TMemberDao{UserName: name}).Find(&res)
		res1 = convertMemberDaoToMember(res)
		if res1.Username == "" {
			return Types.UserNotExisted
		}
		return Types.UserHasDeleted
	}
}

// GetMemberList 从起始id——offset开始查看往后limit个数据直至查询不到
func GetMemberList(offset int, limit int) ([]Types.TMember, Types.ErrNo) {
	var res = make([]TMemberDao, limit, limit)
	var res1 = make([]Types.TMember, limit, limit)
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		db.Limit(limit).Offset(offset).Order("id").Find(&res)
		for key := range res {
			res1[key] = convertMemberDaoToMember(res[key])
		}
		return res1, Types.OK
	}
}

// GetAllMemberList 返回数据库中所有的信息
func GetAllMemberList() ([]Types.TMember, Types.ErrNo) {
	var res TMemberDao
	var res1 = make([]Types.TMember, 10, 10)
	db, err := DBAccessor.MySqlInit()
	defer func(db *gorm.DB) {
		_ = db.Close()
	}(db)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Database connection refused.")
		return res1, Types.UnknownError
	} else {
		// 直到建表成功才继续
		for true {
			if makeMemberTable() {
				break
			} else {
				// 如果建表失败，停4s并输出提示信息
				time.Sleep(time.Duration(4))
				fmt.Println("Something happened when trying to establish the table--'members'.Please check the database.")
			}
		}
		var offset int = 0
		for true {
			db.Unscoped().Where("id = ?", offset).Find(&res)
			if res.UserName == "" {
				break
			} else {
				offset++
				res1 = append(res1, convertMemberDaoToMember(res))
			}
		}
		return res1, Types.OK
	}
}
