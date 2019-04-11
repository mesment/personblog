package db

import (
	"github.com/mesment/personblog/models"
	"errors"
)

//添加一个用户
func AddUser(loginName string, pwd string) error {
	stmtIns:= `INSERT INTO users(name,password) values (?,?)`

	_,err := DB.Exec(stmtIns,loginName,pwd)
	return err
}

//根据用户名查找用户,如果用户存在返回用户，error返回具体的错误信息
func GetUser(loginName string) (user *models.User, err error) {
	//判断用户名是否存在
	exist := UserExist(loginName)
	if !exist {
		user = nil
		err = errors.New("用户不存在")
		return
	}
	query :=`SELECT id,name,password from users where name=? limit 1`

	user = &models.User{}
	err = DB.Get(user,query,loginName)
	if err != nil {
		user = nil
	}
	return
}

//判断用户名是否已存在
func UserExist(loginName string)  bool {
	query :=`SELECT count(*) from users where name=? `
	var num int
	err := DB.Get(&num,query,loginName)
	if err != nil {
		return false
	}
	if num == 0 {
		return false
	}
	return true
}

// 更新用户密码
func UpdateUser(username, oldpass, newpass string) error {
	str:= `UPDATE users SET password=? WHERE name=? AND password=?`
	_, err := DB.Exec(str,username,oldpass)

	return err
}

//删除用户
func DeleteUser(loginName string,pwd string) error {
	stmtDel := `DELETE FROM users where name=? and password=?`
	_,err :=DB.Exec(stmtDel,loginName,pwd)
	return err
}

