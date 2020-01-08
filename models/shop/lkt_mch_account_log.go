package shop

import (
	"time"
)

type MchAccountLog struct {
	Id           int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId      int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	MchId        string    `xorm:"comment('店铺id') CHAR(15)"`
	Price        string    `xorm:"not null default 0.00 comment('金额') DECIMAL(12,2)"`
	AccountMoney string    `xorm:"not null default 0.00 comment('商户余额') DECIMAL(12,2)"`
	Status       int       `xorm:"not null default 1 comment('状态：1.入账 2.出账') TINYINT(3)"`
	Type         int       `xorm:"not null default 1 comment('类型：1.订单 2.退款 3.提现') TINYINT(3)"`
	Addtime      time.Time `xorm:"comment('时间') TIMESTAMP"`
}
