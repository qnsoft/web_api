package shop

import (
	"time"
)

type TwelveDrawUserAddress struct {
	Id      int       `xorm:"not null comment('地址id') INT(11)"`
	StoreId int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	Name    string    `xorm:"not null comment('收货人') VARCHAR(10)"`
	Tel     string    `xorm:"not null comment('联系方式') CHAR(15)"`
	Address string    `xorm:"not null comment('收货地址') VARCHAR(255)"`
	Uid     string    `xorm:"not null default '0' comment('用户ID') VARCHAR(32)"`
	Oid     int       `xorm:"not null comment('记录id') INT(11)"`
	Addtime time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Kdname  string    `xorm:"comment('快递名称') VARCHAR(50)"`
	Kdid    string    `xorm:"comment('快递名称id') VARCHAR(50)"`
	KdNum   string    `xorm:"comment('快递单号') VARCHAR(100)"`
}
