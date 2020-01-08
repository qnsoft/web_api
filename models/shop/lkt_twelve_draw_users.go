package shop

import (
	"time"
)

type TwelveDrawUsers struct {
	Id      int       `xorm:"not null comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId  string    `xorm:"not null default '' comment('活动名称') VARCHAR(50)"`
	Rate    string    `xorm:"default '' comment('单独几率') VARCHAR(11)"`
	Fid     int       `xorm:"comment('活动场次id') INT(10)"`
	Pid     int       `xorm:"default 0 comment('指定中奖商品') INT(10)"`
	Status  string    `xorm:"not null default '' comment('黑名单状态 0:关闭 1:打开') CHAR(6)"`
	Addtime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}
