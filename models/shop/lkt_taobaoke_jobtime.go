package shop

import (
	"time"
)

type TaobaokeJobtime struct {
	Id                    int       `xorm:"not null pk autoincr INT(11)"`
	StoreId               int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	AllYesterdayStartTime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('采集昨天所有订单开始时间') TIMESTAMP"`
	OkYesterdayStartTime  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('采集昨天结算订单开始时间') TIMESTAMP"`
}
