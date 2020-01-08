package shop

import (
	"time"
)

/*
用户实名认证信息
*/
type UserAuth struct {
	Id            int       `xorm:"not null pk autoincr comment('主键') INT(10)"`
	StoreId       int       `xorm:"not null comment('商城id') INT(11)"`
	UserId        string    `xorm:"not null comment('用户id') CHAR(15)"`
	Truename      string    `xorm:"not null comment('真实姓名') VARCHAR(50)"`
	Idcard        string    `xorm:"not null comment('身份证号') VARCHAR(50)"`
	Personimg     string    `xorm:"comment('身份证正面') VARCHAR(255)"`
	Personimgback string    `xorm:"comment('身份证背面') VARCHAR(255)"`
	Authflag      int       `xorm:"not null comment('实名标识') BIT(1)"`
	Value         string    `xorm:"comment('预留字段') VARCHAR(255)"`
	Ysbflag       int       `xorm:"not null comment('认证标识') BIT(1)"`
	Ordertop      int       `xorm:"not null comment('排序') INT(11)"`
	Addtime       time.Time `xorm:"not null comment('认证时间') DATETIME"`
	EditUser      string    `xorm:"comment('') VARCHAR(20)"`
	Cardno        string    `xorm:"comment('') VARCHAR(50)"`
	Bankno        string    `xorm:"comment('') VARCHAR(20)"`
	Submerchid    string    `xorm:"comment('工易付商户编号') VARCHAR(50)"`
}
