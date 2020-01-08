package shop

import (
	"time"
)

type ReturnGoods struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name       string    `xorm:"not null comment('收货人') VARCHAR(50)"`
	Tel        string    `xorm:"not null comment('联系方式') CHAR(15)"`
	Express    string    `xorm:"not null comment('快递名称') VARCHAR(255)"`
	ExpressNum string    `xorm:"not null comment('快递单号') VARCHAR(255)"`
	Uid        string    `xorm:"not null default '0' comment('用户ID') VARCHAR(32)"`
	Oid        string    `xorm:"not null default '0' comment('订单id') VARCHAR(32)"`
	AddData    time.Time `xorm:"default 'CURRENT_TIMESTAMP' comment('填写时间') TIMESTAMP"`
	UserId     string    `xorm:"not null default '' comment('用户id') CHAR(20)"`
}
