package shop

import (
	"time"
)

type OrderConfig struct {
	Id           int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId      int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Days         int       `xorm:"comment('承若天数') INT(11)"`
	Content      string    `xorm:"comment('承若内容') VARCHAR(30)"`
	Back         int       `xorm:"default 2 comment('退货时间') INT(11)"`
	OrderFailure int       `xorm:"comment('订单失效') INT(11)"`
	Company      string    `xorm:"default '天' comment('单位') VARCHAR(20)"`
	OrderOverdue int       `xorm:"default 2 comment('订单过期删除时间') INT(11)"`
	Unit         string    `xorm:"default '天' comment('单位') VARCHAR(50)"`
	ModifyDate   time.Time `xorm:"comment('修改时间') TIMESTAMP"`
	AutoTheGoods int       `xorm:"default 7 comment('自动收货时间') INT(11)"`
}
