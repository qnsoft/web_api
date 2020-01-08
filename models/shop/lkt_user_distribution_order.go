package shop

import (
	"time"
)

type UserDistributionOrder struct {
	Id        int       `xorm:"not null autoincr index INT(11)"`
	Type      int       `xorm:"default 2 comment('订单类型') INT(11)"`
	OrderId   string    `xorm:"not null comment('订单id') VARCHAR(32)"`
	OrderJine string    `xorm:"not null comment('订单金额') DECIMAL(10)"`
	AddTime   time.Time `xorm:"not null comment('订单时间') DATETIME"`
	Remarks   string    `xorm:"not null comment('订单备注') TEXT"`
}
