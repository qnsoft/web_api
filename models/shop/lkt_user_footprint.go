package shop

import (
	"time"
)

type UserFootprint struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId  string    `xorm:"not null comment('用户ID') VARCHAR(15)"`
	PId     int       `xorm:"comment('产品id') INT(11)"`
	AppType string    `xorm:"default 'mini_program' comment('平台类型') VARCHAR(50)"`
	AddTime time.Time `xorm:"comment('添加时间') TIMESTAMP"`
}
