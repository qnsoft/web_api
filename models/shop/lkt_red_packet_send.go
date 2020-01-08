package shop

import (
	"time"
)

type RedPacketSend struct {
	RId        int       `xorm:"not null pk autoincr comment('红包ID') INT(11)"`
	UserId     string    `xorm:"not null comment('user_id') VARCHAR(255)"`
	Money      string    `xorm:"comment('金额') DECIMAL(18,2)"`
	Num        int       `xorm:"not null comment('数量') INT(11)"`
	Remarks    string    `xorm:"default '恭喜发财,大吉大利' comment('备注') VARCHAR(255)"`
	Time       time.Time `xorm:"comment('发送时间') TIMESTAMP"`
	Detailed   string    `xorm:"not null comment('红包金额') TEXT"`
	Version    int       `xorm:"default 1 comment('版本控制') INT(11)"`
	Type       int       `xorm:"not null default 1 comment('红包类型（1.随机  2.普通）') INT(2)"`
	TradeNo    string    `xorm:"comment('微信订单号') CHAR(32)"`
	StoreId    int       `xorm:"not null comment('商城id') INT(11)"`
	ExpireDate time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('过期时间') TIMESTAMP"`
}
