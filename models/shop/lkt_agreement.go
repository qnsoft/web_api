package shop

import (
	"time"
)

type Agreement struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Type       int       `xorm:"not null default 0 comment('类型 0:注册 1:店铺') TINYINT(4)"`
	Content    string    `xorm:"not null comment('内容') TEXT"`
	ModifyDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('修改时间') TIMESTAMP"`
}
