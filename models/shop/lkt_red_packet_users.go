package shop

import (
	"time"
)

type RedPacketUsers struct {
	Id      int       `xorm:"not null pk autoincr comment('ID') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId  string    `xorm:"not null default '' comment('活动名称') VARCHAR(50)"`
	Status  string    `xorm:"default '' comment('0:关闭 1:打开') CHAR(6)"`
	Addtime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}
