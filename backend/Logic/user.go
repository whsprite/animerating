package Logic

import (
	"Goweb_cli/dao/mysql"
	"Goweb_cli/modules"
	"Goweb_cli/pkg/jwt"
	"Goweb_cli/pkg/snowflake"
)

func SignUp(p *modules.SignUpParam) (err error) {
	//验证用户名是否重复
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return
	}
	//生成ID
	userid := snowflake.GenID()
	u := &modules.User{
		UserID:   userid,
		Username: p.Username,
		Password: p.Password,
	}
	//插入用户
	if err = mysql.InsertUser(u); err != nil {
		return
	}
	return
}
func Login(p *modules.LoginParam) (token string, err error) {
	u := &modules.User{
		UserID:   0,
		Username: p.Username,
		Password: p.Password,
	}
	if err = mysql.CheckUser(u); err != nil {

		return "", err
	}
	token, err = jwt.GenToken(u.UserID, u.Username)
	if err != nil {
		return "", err
	}
	return token, err
}
