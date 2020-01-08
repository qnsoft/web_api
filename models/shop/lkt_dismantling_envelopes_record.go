package shop

import (
	"time"
)

type DismantlingEnvelopesRecord struct {
	Id      int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Hid     int       `xorm:"not null default 0 comment('活动id') INT(11)"`
	UserId  string    `xorm:"not null comment('用户ID') VARCHAR(15)"`
	Money   string    `xorm:"not null default 0.00 comment('金额') DECIMAL(12,2)"`
	AddTime time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Event   string    `xorm:"comment('事件') VARCHAR(200)"`
	Type    int       `xorm:"not null default 0 comment('类型 0:发起人 1:被邀请人') TINYINT(4)"`
	Status  int       `xorm:"not null default 0 comment('状态 0:发起拆红包  2:拆红包成功 3:逾期失效') TINYINT(4)"`
}
