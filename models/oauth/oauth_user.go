package oauth

import (
	"github.com/qnsoft/utils/DbHelper"
	"github.com/qnsoft/utils/ErrorHelper"
	"time"
)

/*
*@token用户信息实体
 */
type OauthUser struct {
	Id            int       `xorm:"not null pk autoincr comment('主键') INT(11)"`
	Name          string    `xorm:"not null comment('用户名') VARCHAR(32)"`
	Passwd        string    `xorm:"not null comment('密码') VARCHAR(32)"`
	Email         string    `xorm:"not null comment('emial') VARCHAR(200)"`
	Status        int       `xorm:"not null comment('') INT(11)"`
	CreateTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
	LastLoginTime time.Time `xorm:"not null comment('最后时间') DATETIME"`
	RoleId        string    `xorm:"not null comment('') INT(11)"`
	Appid         string    `xorm:"not null comment('') VARCHAR(32)"`
	Appsecret     string    `xorm:"not null comment('') VARCHAR(32)"`
	Salt          string    `xorm:"not null comment('') VARCHAR(255)"`
}

// func init() {
// 	orm.RegisterModelWithPrefix(beego.AppConfig.String("database::db_prefix"), new(User))
// }

/*
*@检查登录
*@_Appid 用户id
*@_Appsecret 用户密码
 */
func (s *OauthUser) Check_Oauth(_Appid, _Appsecret string) (bool, error) {
	_model := &OauthUser{Appid: _Appid, Appsecret: _Appsecret}
	results, err := DbHelper.MySqlDb().Get(_model)
	ErrorHelper.CheckErr(err)
	return results, err
}
