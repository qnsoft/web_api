package shop

import (
	"time"
)

type UserDistribution struct {
	Id         int       `xorm:"not null pk autoincr comment('id') INT(11)"`
	StoreId    int       `xorm:"not null default 0 comment('商城id') INT(11)"`
	UserId     string    `xorm:"not null comment('用户ID') VARCHAR(32)"`
	Pid        string    `xorm:"not null comment('上级id') VARCHAR(32)"`
	Level      int       `xorm:"not null comment('等级') INT(11)"`
	Lt         int       `xorm:"not null comment('lt') INT(11)"`
	Rt         int       `xorm:"not null comment('rt') INT(11)"`
	Uplevel    int       `xorm:"not null comment('第几代') INT(11)"`
	AddDate    time.Time `xorm:"comment('添加时间') TIMESTAMP"`
	Usets      string    `xorm:"VARCHAR(255)"`
	Commission string    `xorm:"not null default 0.00 comment('累计佣金') DECIMAL(10,2)"`
	Onlyamount string    `xorm:"not null default 0.00 comment('累计消费') DECIMAL(15,2)"`
	Allamount  string    `xorm:"not null default 0.00 comment('销售业绩') DECIMAL(15,2)"`
	Type       int       `xorm:"not null default 2 comment('类型（推人1、信用卡2、商城消费3、商家销售返佣4、实体店积分消费5') INT(11)"`
	OrderId    string    `xorm:"not null comment('分润相关订单id') VARCHAR(32)"`
	UserId1    string    `xorm:"VARCHAR(32)"`
}
