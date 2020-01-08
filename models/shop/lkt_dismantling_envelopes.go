package shop

import (
	"time"
)

type DismantlingEnvelopes struct {
	Id        int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId   int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name      string    `xorm:"not null comment('活动名称') VARCHAR(50)"`
	Image     string    `xorm:"not null comment('活动图片') VARCHAR(255)"`
	ZMoney    string    `xorm:"not null default 0.00 comment('红包金额') DECIMAL(12,2)"`
	Type      int       `xorm:"not null default 0 comment('类型 0:首页拆红包 1:确认收货拆红包') TINYINT(4)"`
	AddTime   time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	StartTime time.Time `xorm:"comment('开始时间') TIMESTAMP"`
	EndTime   time.Time `xorm:"comment('结束时间') TIMESTAMP"`
	Content   string    `xorm:"comment('活动介绍') TEXT"`
	Status    int       `xorm:"not null default 0 comment('状态 0：未启用 1：启用 2:禁用 3：已结束') TINYINT(4)"`
}
