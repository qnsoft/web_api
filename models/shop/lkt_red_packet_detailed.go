package shop

import (
	"time"
)

type RedPacketDetailed struct {
	Id      int       `xorm:"not null pk autoincr comment('ID') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	RId     int       `xorm:"not null comment('红包ID') INT(11)"`
	UserId  string    `xorm:"not null comment('user_id') VARCHAR(255)"`
	Money   string    `xorm:"comment('领取的金额') DECIMAL(18,2)"`
	Time    time.Time `xorm:"comment('领取时间') DATETIME"`
}
