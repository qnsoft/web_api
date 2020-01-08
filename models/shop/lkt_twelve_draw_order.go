package shop

import (
	"time"
)

type TwelveDrawOrder struct {
	Id          int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId     int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId      string    `xorm:"not null comment('用户id') CHAR(15)"`
	Dpid        string    `xorm:"not null comment('奖品id') VARCHAR(20)"`
	Num         int       `xorm:"comment('数量') INT(11)"`
	Fid         int       `xorm:"default 0 comment('活动场次id') INT(11)"`
	Price       int       `xorm:"not null default 0 comment('总价') INT(12)"`
	Pay         string    `xorm:"not null comment('支付方式') CHAR(15)"`
	Status      int       `xorm:"not null default 0 comment('状态 1：已领取 2：待领取 3：发货中 4：已失效 5:待发货') TINYINT(4)"`
	Addtime     time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	ReceiveTime time.Time `xorm:"comment('领取时间') TIMESTAMP"`
}
