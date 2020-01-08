package shop

import (
	"time"
)

type Role struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	Name       string    `xorm:"not null default '' comment('角色名称') VARCHAR(30)"`
	Permission string    `xorm:"comment('许可') TEXT"`
	Status     int       `xorm:"not null default 0 comment('状态 0:角色 1:客户端') TINYINT(4)"`
	AdminId    int       `xorm:"not null default 0 comment('管理员id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	AddDate    time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
