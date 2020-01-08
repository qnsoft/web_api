package shop

import (
	"time"
)

type UserCollection struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId  string    `xorm:"not null default '0' comment('用户ID') VARCHAR(32)"`
	PId     int       `xorm:"comment('产品id') INT(11)"`
	AddTime time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	MchId   int       `xorm:"default 0 comment('店铺ID') INT(11)"`
}
