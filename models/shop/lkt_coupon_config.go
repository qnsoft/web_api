package shop

import (
	"time"
)

type CouponConfig struct {
	Id              int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId         int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	ActivityOverdue int       `xorm:"not null default 0 comment('活动过期删除时间') INT(11)"`
	CouponValidity  int       `xorm:"not null default 0 comment('优惠券有效期') INT(11)"`
	CouponOverdue   int       `xorm:"not null default 0 comment('优惠券过期删除时间') INT(11)"`
	ModifyDate      time.Time `xorm:"comment('修改时间') TIMESTAMP"`
}
