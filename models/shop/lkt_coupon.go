package shop

import (
	"time"
)

type Coupon struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId     string    `xorm:"not null comment('用户ID') VARCHAR(15)"`
	Money      int       `xorm:"not null default 0 comment('优惠券金额') INT(11)"`
	AddTime    time.Time `xorm:"comment('领取时间') TIMESTAMP"`
	ExpiryTime time.Time `xorm:"comment('到期时间') TIMESTAMP"`
	Hid        int       `xorm:"not null default 0 comment('活动id') INT(11)"`
	Type       int       `xorm:"not null default 0 comment('状态 0:未使用 1:使用中 2:已使用 3:已过期') TINYINT(4)"`
}
