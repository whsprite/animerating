package mysql

import (
	"Goweb_cli/modules"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

const secret string = "f2mvvdenjcll"

var (
	ErrorUserExist    = errors.New("用户名已存在")
	ErrorInvalidParam = errors.New("参数不正确")
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(username) from user where username = ?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {

		return
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}
func InsertUser(u *modules.User) (err error) {
	passowrd := encryptPassword(u.Password)
	sqlStr := `insert into user (user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, u.UserID, u.Username, passowrd)
	return
}

func CheckUser(u *modules.User) (err error) {
	password := encryptPassword(u.Password)
	sqlStr := `select user_id,username,password from user where username = ?`
	if err = db.Get(u, sqlStr, u.Username); err != nil {
		return
	}
	if password != u.Password {
		return ErrorInvalidParam
	}
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
