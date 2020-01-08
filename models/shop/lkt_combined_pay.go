package shop

import (
	"time"
)

type CombinedPay struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	WeixinPay  string    `xorm:"not null default 0.00 comment('微信支付金额') DECIMAL(11,2)"`
	BalancePay string    `xorm:"not null default 0.00 comment('余额支付') DECIMAL(11,2)"`
	Total      string    `xorm:"not null comment('总金额') DECIMAL(11,2)"`
	OrderId    string    `xorm:"not null comment('订单号') VARCHAR(255)"`
	AddTime    time.Time `xorm:"not null comment('创建时间') DATETIME"`
	UserId     string    `xorm:"comment('用户id') VARCHAR(100)"`
}
