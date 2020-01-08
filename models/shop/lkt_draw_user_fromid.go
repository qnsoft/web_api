package shop

import (
	"time"
)

type DrawUserFromid struct {
	Id       int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId  int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	OpenId   string    `xorm:"not null default '' comment('用户openid') CHAR(40)"`
	Fromid   string    `xorm:"not null default '' comment('用户fromid') CHAR(50)"`
	Lifetime time.Time `xorm:"comment('生命周期') DATETIME"`
}
