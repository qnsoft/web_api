package shop

import (
	"time"
)

type OrderData struct {
	Id      int       `xorm:"not null INT(11)"`
	TradeNo string    `xorm:"comment('微信订单号') CHAR(32)"`
	Data    string    `xorm:"TEXT"`
	Addtime time.Time `xorm:"not null comment('创建时间') DATETIME"`
}
