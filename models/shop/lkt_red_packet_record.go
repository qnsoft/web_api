package shop

import (
	"time"
)

type RedPacketRecord struct {
	Id      int       `xorm:"not null pk autoincr comment('ID') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	RId     int       `xorm:"comment('红包ID') INT(11)"`
	UserId  string    `xorm:"not null comment('user_id') VARCHAR(255)"`
	Money   string    `xorm:"not null comment('金额') DECIMAL(10,2)"`
	Time    time.Time `xorm:"comment('时间') DATETIME"`
	Type    int       `xorm:"comment('1.领红包  2.正常使用红包 3.提现的红包') INT(11)"`
}
