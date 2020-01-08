package shop

import (
	"time"
)

type MchBrowse struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	MchId   string    `xorm:"comment('店铺id') CHAR(15)"`
	Token   string    `xorm:"comment('token') VARCHAR(255)"`
	UserId  string    `xorm:"comment('用户id') CHAR(15)"`
	Event   string    `xorm:"comment('事件') VARCHAR(255)"`
	AddTime time.Time `xorm:"comment('时间') TIMESTAMP"`
}
